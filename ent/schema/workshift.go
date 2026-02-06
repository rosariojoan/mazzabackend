package schema

import (
	"context"
	gen "mazza/ent/generated"
	"mazza/ent/generated/hook"
	"mazza/ent/generated/user"
	"mazza/ent/generated/workshift"
	"mazza/ent/utils"
	"mazza/firebase"

	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Workshift holds the schema definition for the Workshift entity.
type Workshift struct {
	ent.Schema
}

// type Location struct {
// 	Latitude    float64 `json:"latitude"`
// 	Longitude   float64 `json:"longitude"`
// 	Description string  `json:"description"`
// }

func (Workshift) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Workshift.
func (Workshift) Fields() []ent.Field {
	return []ent.Field{
		field.Time("approved_at").Nillable().Optional().Annotations(entgql.OrderField("approvedAt")).Comment("time that this shift was approved by the supervisor"),
		field.Time("clock_in").Default(time.Now).Annotations(entgql.OrderField("clockIn")),
		field.Time("clock_out").Nillable().Optional().Annotations(entgql.OrderField("clockOut")),
		field.String("clock_in_location").Comment("it expects a serialized json like: {latitude: float, longitude: float, description: string}"),
		field.String("clock_out_location").Optional().Comment("it expects a serialized json like: {latitude: float, longitude: float, description: string}"),
		field.String("description").Optional(),
		field.String("note").Optional().Comment("this is only used when the current item is a shift edit request"),
		field.Enum("status").Values("approved", "pending").Default("pending").Annotations(entgql.OrderField("status")),
	}
}

// Edges of the Workshift.
func (Workshift) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("work_shifts").Unique(),
		edge.From("user", User.Type).Ref("work_shifts").Unique(),
		edge.From("approved_by", User.Type).Ref("approved_work_shifts").Unique(),
		edge.From("task", ProjectTask.Type).Ref("work_shifts").Unique(),
		edge.To("edit_request", Workshift.Type).Unique(),
		edge.From("work_shift", Workshift.Type).Ref("edit_request").Unique(),
	}
}

// Enable query and mutation for the CompanyClient schema
func (Workshift) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		// entgql.Type("WorkshiftGroupBy"),
	}
}

// Hooks of the card. Send notification:
// to the employee's leader on shift creation
// to the employee on shift approval
func (Workshift) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.WorkshiftFunc(func(ctx context.Context, m *gen.WorkshiftMutation) (ent.Value, error) {
					// On clock-out, if the employee has a leader, set status to "PENDING", otherwise set status "APPROVED"
					userID, exists := m.UserID()
					if !exists {
						currentUser, _ := utils.GetSession(&ctx)
						userID = currentUser.ID
						// return next.Mutate(ctx, m)
					}

					leader, err := m.Client().User.Query().Where(user.IDEQ(userID)).
						QueryLeader().Select(user.FieldID, user.FieldFcmToken).First(ctx)

					if err != nil {
						// This employee does not have a leader, so his workshift status is automatically approved.
						// Proceed with the mutation operation and return.
						m.SetStatus(workshift.StatusApproved)
						m.SetApprovedAt(time.Now())
						m.SetApprovedByID(userID)
						return next.Mutate(ctx, m)
					}
					v, err := next.Mutate(ctx, m)

					// Post mutation: this user has a leader, send a notification to the leader after creating the workshift
					if leader.FcmToken != nil {
						data := struct {
							AlertType string    `json:"alertType"`
							UserID    int       `json:"userID"`
							TimeSent  time.Time `json:"timeSent"`
						}{
							AlertType: firebase.AlertType.WorkShiftApprovalRequest,
							UserID:    leader.ID,
							TimeSent:  time.Now(),
						}

						go func() {
							dataMap, err := firebase.GetMap(data)
							if err != nil {
								return
							}
							// Send notification to the closest, available driver
							title := "Novo timesheet submetido"
							firebase.SendDataNotification([]string{*leader.FcmToken}, &title, dataMap)
						}()
					}

					return v, err
				})
			},
			// Limit the hook for only update operations, when the user clocks out
			ent.OpUpdateOne,
		),

		hook.On(
			// Workshift update request hook: send a notification to the leader
			func(next ent.Mutator) ent.Mutator {
				return hook.WorkshiftFunc(func(ctx context.Context, m *gen.WorkshiftMutation) (ent.Value, error) {
					workshiftID, exists := m.WorkShiftID()

					// Do nothing if this is not a workshift update request
					if !exists {
						return next.Mutate(ctx, m)
					}

					leader, err := m.Client().Workshift.Query().Where(workshift.IDEQ(workshiftID)).
						QueryUser().QueryLeader().
						Select(user.FieldID, user.FieldFcmToken).First(ctx)

					if err != nil {
						updater := m.Client().Workshift.UpdateOneID(workshiftID).SetApprovedAt(time.Now())
						if clockIn, exists := m.ClockIn(); exists {
							updater.SetClockIn(clockIn)
						}
						if clockOut, exists := m.ClockOut(); exists {
							updater.SetClockOut(clockOut)
						}
						if description, exists := m.Description(); exists {
							updater.SetDescription(description)
						}
						if note, exists := m.Note(); exists {
							updater.SetNote(note)
						}

						return updater.Save(ctx)
						// return next.Mutate(ctx, m)
					}

					data := struct {
						AlertType string    `json:"alertType"`
						UserID    int       `json:"userID"`
						TimeSent  time.Time `json:"timeSent"`
					}{
						AlertType: firebase.AlertType.WorkShifUpdateRequest,
						UserID:    leader.ID,
						TimeSent:  time.Now(),
					}

					go func() {
						dataMap, err := firebase.GetMap(data)
						if err != nil {
							return
						}
						// Send notification to the closest, available driver
						title := "Timesheet"
						firebase.SendDataNotification([]string{*leader.FcmToken}, &title, dataMap)
					}()
					_ = leader
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}
