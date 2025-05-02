package schema

import (
	"context"
	"fmt"
	"math"
	gen "mazza/ent/generated"
	"mazza/ent/generated/hook"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projecttask"

	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ProjectTask holds the schema definition for the ProjectTask entity.
type ProjectTask struct {
	ent.Schema
}

// Fields of the ProjectTask.
func (ProjectTask) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now).Immutable().Optional(),
		field.String("name").NotEmpty(),
		field.String("assigneeName").NotEmpty(),
		field.String("location").Optional().Comment("Where is task will be executed"),
		field.Time("dueDate").Annotations(entgql.OrderField("DUE_DATE")), // creates an order index on this fiels to avoid a full table scan
		field.Time("startDate").Annotations(entgql.OrderField("START_DATE")),
		field.Time("endDate").Nillable().Optional().Annotations(entgql.OrderField("END_DATE")),
		field.String("description").Optional(),
		field.Enum("status").Values("notStarted", "inProgress", "completed").Annotations(entgql.OrderField("STATUS")),
	}
}

// Edges of the ProjectTask.
func (ProjectTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("tasks").Unique().Required(),   // a task can belong to only one project
		edge.From("assignee", User.Type).Ref("assignedProjectTasks").Unique(), // only one user responds for a task
		edge.From("participants", User.Type).Ref("participatedProjectTasks"),  // more than one user can participate in a task. E.g. meeting
		edge.From("createdBy", User.Type).Ref("createdTasks").Unique().Immutable(),
		edge.To("workShifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

// Create indexes
func (ProjectTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("project").Unique(), // each milestone in a project is unique
	}
}

// Enable query and mutation for the Project schema
func (ProjectTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Hooks of the workttask. Update project status on workttask mutation.
func (ProjectTask) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.ProjectTaskFunc(func(ctx context.Context, m *gen.ProjectTaskMutation) (ent.Value, error) {
					// Get projectID before executing the mutation
					var projectID int
					var err error
					if projID, exists := m.ProjectID(); exists {
						projectID = projID
					} else if taskID, exists := m.ID(); exists {
						projectID, err = m.Client().Project.Query().Where(project.HasTasksWith(projecttask.IDEQ(taskID))).FirstID(ctx)
						if err != nil {
							fmt.Println("hook err:", err)
							return nil, fmt.Errorf("invalid project or task")
						}
					} else {
						return nil, fmt.Errorf("invalid project or task")
					}

					// Execute mutation
					v, err := next.Mutate(ctx, m)

					// Post mutation: update project status
					if err == nil {
						tasks, err := m.Client().Project.Query().Where(project.IDEQ(projectID)).QueryTasks().Select(projecttask.FieldStatus).All(ctx)
						if err != nil {
							fmt.Println("err:", err)
							return nil, fmt.Errorf("an error occurred")
						}

						total := len(tasks)
						inProgress := 0
						completed := 0
						hasStartedTask := false
						for _, task := range tasks {
							if task.Status == projecttask.StatusCompleted {
								completed += 1
							} else if task.Status == projecttask.StatusInProgress {
								inProgress += 1
							}

							if task.Status != projecttask.StatusNotStarted {
								hasStartedTask = true
							}
						}

						var progress float64 = 0
						if total > 0 {
							progress = math.Min(float64(completed)/float64(total), 1)
						}

						projectUpdate := m.Client().Project.UpdateOneID(projectID).SetProgress(progress)
						if progress == 1 {
							// A project is marked as completed when all the tasks are completed
							projectUpdate.SetStatus(project.StatusCompleted)
						} else if (inProgress > 0 && progress > 0) || hasStartedTask {
							// If there is any task in progress, the project is in progress
							projectUpdate.SetStatus(project.StatusInProgress)
						} else {
							// If there is no completed task and not in-progress task, the project has not started
							projectUpdate.SetStatus(project.StatusNotStarted)
						}
						_, err = projectUpdate.Save(ctx)
						if err != nil {
							fmt.Println("err:", err)
							return nil, fmt.Errorf("an error occurred")
						}
					}

					return v, err
				})
			},
			// Limit the hook for only update operations, when the user clocks out
			ent.OpCreate|ent.OpUpdateOne|ent.OpDeleteOne,
		),
	}
}
