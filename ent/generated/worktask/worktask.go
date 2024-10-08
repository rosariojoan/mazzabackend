// Code generated by ent, DO NOT EDIT.

package worktask

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the worktask type in the database.
	Label = "worktask"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSubtasks holds the string denoting the subtasks field in the database.
	FieldSubtasks = "subtasks"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeCreatedBy holds the string denoting the createdby edge name in mutations.
	EdgeCreatedBy = "createdBy"
	// EdgeAssignedTo holds the string denoting the assignedto edge name in mutations.
	EdgeAssignedTo = "assignedTo"
	// EdgeWorkShifts holds the string denoting the workshifts edge name in mutations.
	EdgeWorkShifts = "workShifts"
	// EdgeWorkTags holds the string denoting the worktags edge name in mutations.
	EdgeWorkTags = "workTags"
	// Table holds the table name of the worktask in the database.
	Table = "worktasks"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "worktasks"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_work_tasks"
	// CreatedByTable is the table that holds the createdBy relation/edge.
	CreatedByTable = "worktasks"
	// CreatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedByInverseTable = "users"
	// CreatedByColumn is the table column denoting the createdBy relation/edge.
	CreatedByColumn = "user_created_tasks"
	// AssignedToTable is the table that holds the assignedTo relation/edge. The primary key declared below.
	AssignedToTable = "employee_assignedTasks"
	// AssignedToInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	AssignedToInverseTable = "employees"
	// WorkShiftsTable is the table that holds the workShifts relation/edge.
	WorkShiftsTable = "workshifts"
	// WorkShiftsInverseTable is the table name for the Workshift entity.
	// It exists in this package in order to avoid circular dependency with the "workshift" package.
	WorkShiftsInverseTable = "workshifts"
	// WorkShiftsColumn is the table column denoting the workShifts relation/edge.
	WorkShiftsColumn = "worktask_work_shifts"
	// WorkTagsTable is the table that holds the workTags relation/edge. The primary key declared below.
	WorkTagsTable = "worktask_workTags"
	// WorkTagsInverseTable is the table name for the Worktag entity.
	// It exists in this package in order to avoid circular dependency with the "worktag" package.
	WorkTagsInverseTable = "worktags"
)

// Columns holds all SQL columns for worktask fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldDescription,
	FieldStatus,
	FieldSubtasks,
	FieldTitle,
	FieldStartTime,
	FieldEndTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "worktasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_work_tasks",
	"user_created_tasks",
}

var (
	// AssignedToPrimaryKey and AssignedToColumn2 are the table columns denoting the
	// primary key for the assignedTo relation (M2M).
	AssignedToPrimaryKey = []string{"employee_id", "worktask_id"}
	// WorkTagsPrimaryKey and WorkTagsColumn2 are the table columns denoting the
	// primary key for the workTags relation (M2M).
	WorkTagsPrimaryKey = []string{"worktask_id", "worktag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "mazza/ent/generated/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusDRAFT    Status = "DRAFT"
	StatusASSIGNED Status = "ASSIGNED"
	StatusDONE     Status = "DONE"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDRAFT, StatusASSIGNED, StatusDONE:
		return nil
	default:
		return fmt.Errorf("worktask: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Worktask queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deletedAt field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByStartTime orders the results by the startTime field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the endTime field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreatedByField orders the results by createdBy field.
func ByCreatedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedByStep(), sql.OrderByField(field, opts...))
	}
}

// ByAssignedToCount orders the results by assignedTo count.
func ByAssignedToCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssignedToStep(), opts...)
	}
}

// ByAssignedTo orders the results by assignedTo terms.
func ByAssignedTo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssignedToStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkShiftsCount orders the results by workShifts count.
func ByWorkShiftsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkShiftsStep(), opts...)
	}
}

// ByWorkShifts orders the results by workShifts terms.
func ByWorkShifts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkShiftsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkTagsCount orders the results by workTags count.
func ByWorkTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkTagsStep(), opts...)
	}
}

// ByWorkTags orders the results by workTags terms.
func ByWorkTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
func newCreatedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatedByTable, CreatedByColumn),
	)
}
func newAssignedToStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssignedToInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AssignedToTable, AssignedToPrimaryKey...),
	)
}
func newWorkShiftsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkShiftsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkShiftsTable, WorkShiftsColumn),
	)
}
func newWorkTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkTagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, WorkTagsTable, WorkTagsPrimaryKey...),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
