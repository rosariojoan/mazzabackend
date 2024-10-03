// Code generated by ent, DO NOT EDIT.

package projecttask

import (
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldName, v))
}

// AssigneeName applies equality check predicate on the "assigneeName" field. It's identical to AssigneeNameEQ.
func AssigneeName(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldAssigneeName, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldLocation, v))
}

// DueDate applies equality check predicate on the "dueDate" field. It's identical to DueDateEQ.
func DueDate(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldDueDate, v))
}

// StartDate applies equality check predicate on the "startDate" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "endDate" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldEndDate, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldDescription, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContainsFold(FieldName, v))
}

// AssigneeNameEQ applies the EQ predicate on the "assigneeName" field.
func AssigneeNameEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldAssigneeName, v))
}

// AssigneeNameNEQ applies the NEQ predicate on the "assigneeName" field.
func AssigneeNameNEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldAssigneeName, v))
}

// AssigneeNameIn applies the In predicate on the "assigneeName" field.
func AssigneeNameIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldAssigneeName, vs...))
}

// AssigneeNameNotIn applies the NotIn predicate on the "assigneeName" field.
func AssigneeNameNotIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldAssigneeName, vs...))
}

// AssigneeNameGT applies the GT predicate on the "assigneeName" field.
func AssigneeNameGT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldAssigneeName, v))
}

// AssigneeNameGTE applies the GTE predicate on the "assigneeName" field.
func AssigneeNameGTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldAssigneeName, v))
}

// AssigneeNameLT applies the LT predicate on the "assigneeName" field.
func AssigneeNameLT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldAssigneeName, v))
}

// AssigneeNameLTE applies the LTE predicate on the "assigneeName" field.
func AssigneeNameLTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldAssigneeName, v))
}

// AssigneeNameContains applies the Contains predicate on the "assigneeName" field.
func AssigneeNameContains(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContains(FieldAssigneeName, v))
}

// AssigneeNameHasPrefix applies the HasPrefix predicate on the "assigneeName" field.
func AssigneeNameHasPrefix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasPrefix(FieldAssigneeName, v))
}

// AssigneeNameHasSuffix applies the HasSuffix predicate on the "assigneeName" field.
func AssigneeNameHasSuffix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasSuffix(FieldAssigneeName, v))
}

// AssigneeNameEqualFold applies the EqualFold predicate on the "assigneeName" field.
func AssigneeNameEqualFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEqualFold(FieldAssigneeName, v))
}

// AssigneeNameContainsFold applies the ContainsFold predicate on the "assigneeName" field.
func AssigneeNameContainsFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContainsFold(FieldAssigneeName, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationIsNil applies the IsNil predicate on the "location" field.
func LocationIsNil() predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIsNull(FieldLocation))
}

// LocationNotNil applies the NotNil predicate on the "location" field.
func LocationNotNil() predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotNull(FieldLocation))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContainsFold(FieldLocation, v))
}

// DueDateEQ applies the EQ predicate on the "dueDate" field.
func DueDateEQ(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldDueDate, v))
}

// DueDateNEQ applies the NEQ predicate on the "dueDate" field.
func DueDateNEQ(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldDueDate, v))
}

// DueDateIn applies the In predicate on the "dueDate" field.
func DueDateIn(vs ...time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldDueDate, vs...))
}

// DueDateNotIn applies the NotIn predicate on the "dueDate" field.
func DueDateNotIn(vs ...time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldDueDate, vs...))
}

// DueDateGT applies the GT predicate on the "dueDate" field.
func DueDateGT(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldDueDate, v))
}

// DueDateGTE applies the GTE predicate on the "dueDate" field.
func DueDateGTE(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldDueDate, v))
}

// DueDateLT applies the LT predicate on the "dueDate" field.
func DueDateLT(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldDueDate, v))
}

// DueDateLTE applies the LTE predicate on the "dueDate" field.
func DueDateLTE(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldDueDate, v))
}

// StartDateEQ applies the EQ predicate on the "startDate" field.
func StartDateEQ(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "startDate" field.
func StartDateNEQ(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "startDate" field.
func StartDateIn(vs ...time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "startDate" field.
func StartDateNotIn(vs ...time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "startDate" field.
func StartDateGT(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "startDate" field.
func StartDateGTE(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "startDate" field.
func StartDateLT(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "startDate" field.
func StartDateLTE(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldStartDate, v))
}

// EndDateEQ applies the EQ predicate on the "endDate" field.
func EndDateEQ(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "endDate" field.
func EndDateNEQ(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "endDate" field.
func EndDateIn(vs ...time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "endDate" field.
func EndDateNotIn(vs ...time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "endDate" field.
func EndDateGT(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "endDate" field.
func EndDateGTE(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "endDate" field.
func EndDateLT(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "endDate" field.
func EndDateLTE(v time.Time) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldEndDate, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldContainsFold(FieldDescription, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.ProjectTask {
	return predicate.ProjectTask(sql.FieldNotIn(FieldStatus, vs...))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.ProjectTask {
	return predicate.ProjectTask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.ProjectTask {
	return predicate.ProjectTask(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssignee applies the HasEdge predicate on the "assignee" edge.
func HasAssignee() predicate.ProjectTask {
	return predicate.ProjectTask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssigneeTable, AssigneeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssigneeWith applies the HasEdge predicate on the "assignee" edge with a given conditions (other predicates).
func HasAssigneeWith(preds ...predicate.User) predicate.ProjectTask {
	return predicate.ProjectTask(func(s *sql.Selector) {
		step := newAssigneeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParticipants applies the HasEdge predicate on the "participants" edge.
func HasParticipants() predicate.ProjectTask {
	return predicate.ProjectTask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ParticipantsTable, ParticipantsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipantsWith applies the HasEdge predicate on the "participants" edge with a given conditions (other predicates).
func HasParticipantsWith(preds ...predicate.User) predicate.ProjectTask {
	return predicate.ProjectTask(func(s *sql.Selector) {
		step := newParticipantsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProjectTask) predicate.ProjectTask {
	return predicate.ProjectTask(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProjectTask) predicate.ProjectTask {
	return predicate.ProjectTask(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProjectTask) predicate.ProjectTask {
	return predicate.ProjectTask(sql.NotPredicates(p))
}