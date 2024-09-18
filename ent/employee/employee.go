// Code generated by ent, DO NOT EDIT.

package employee

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the employee type in the database.
	Label = "employee"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeEmployees holds the string denoting the employees edge name in mutations.
	EdgeEmployees = "employees"
	// EdgeSupervisor holds the string denoting the supervisor edge name in mutations.
	EdgeSupervisor = "supervisor"
	// EdgeWorkShifts holds the string denoting the workshifts edge name in mutations.
	EdgeWorkShifts = "workShifts"
	// EdgeApprovedWorkShifts holds the string denoting the approvedworkshifts edge name in mutations.
	EdgeApprovedWorkShifts = "approvedWorkShifts"
	// EdgeAssignedTasks holds the string denoting the assignedtasks edge name in mutations.
	EdgeAssignedTasks = "assignedTasks"
	// Table holds the table name of the employee in the database.
	Table = "employees"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "employees"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_employees"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "employees"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_employee"
	// EmployeesTable is the table that holds the employees relation/edge.
	EmployeesTable = "employees"
	// EmployeesColumn is the table column denoting the employees relation/edge.
	EmployeesColumn = "employee_employees"
	// SupervisorTable is the table that holds the supervisor relation/edge.
	SupervisorTable = "employees"
	// SupervisorColumn is the table column denoting the supervisor relation/edge.
	SupervisorColumn = "employee_employees"
	// WorkShiftsTable is the table that holds the workShifts relation/edge.
	WorkShiftsTable = "workshifts"
	// WorkShiftsInverseTable is the table name for the Workshift entity.
	// It exists in this package in order to avoid circular dependency with the "workshift" package.
	WorkShiftsInverseTable = "workshifts"
	// WorkShiftsColumn is the table column denoting the workShifts relation/edge.
	WorkShiftsColumn = "employee_work_shifts"
	// ApprovedWorkShiftsTable is the table that holds the approvedWorkShifts relation/edge.
	ApprovedWorkShiftsTable = "workshifts"
	// ApprovedWorkShiftsInverseTable is the table name for the Workshift entity.
	// It exists in this package in order to avoid circular dependency with the "workshift" package.
	ApprovedWorkShiftsInverseTable = "workshifts"
	// ApprovedWorkShiftsColumn is the table column denoting the approvedWorkShifts relation/edge.
	ApprovedWorkShiftsColumn = "employee_approved_work_shifts"
	// AssignedTasksTable is the table that holds the assignedTasks relation/edge. The primary key declared below.
	AssignedTasksTable = "employee_assignedTasks"
	// AssignedTasksInverseTable is the table name for the Worktask entity.
	// It exists in this package in order to avoid circular dependency with the "worktask" package.
	AssignedTasksInverseTable = "worktasks"
)

// Columns holds all SQL columns for employee fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldGender,
	FieldPosition,
	FieldEmail,
	FieldPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "employees"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_employees",
	"employee_employees",
	"user_employee",
}

var (
	// AssignedTasksPrimaryKey and AssignedTasksColumn2 are the table columns denoting the
	// primary key for the assignedTasks relation (M2M).
	AssignedTasksPrimaryKey = []string{"employee_id", "worktask_id"}
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

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMale, GenderFemale:
		return nil
	default:
		return fmt.Errorf("employee: invalid enum value for gender field: %q", ge)
	}
}

// OrderOption defines the ordering options for the Employee queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByPosition orders the results by the position field.
func ByPosition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPosition, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByEmployeesCount orders the results by employees count.
func ByEmployeesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEmployeesStep(), opts...)
	}
}

// ByEmployees orders the results by employees terms.
func ByEmployees(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmployeesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySupervisorField orders the results by supervisor field.
func BySupervisorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSupervisorStep(), sql.OrderByField(field, opts...))
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

// ByApprovedWorkShiftsCount orders the results by approvedWorkShifts count.
func ByApprovedWorkShiftsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newApprovedWorkShiftsStep(), opts...)
	}
}

// ByApprovedWorkShifts orders the results by approvedWorkShifts terms.
func ByApprovedWorkShifts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApprovedWorkShiftsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAssignedTasksCount orders the results by assignedTasks count.
func ByAssignedTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssignedTasksStep(), opts...)
	}
}

// ByAssignedTasks orders the results by assignedTasks terms.
func ByAssignedTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssignedTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
func newEmployeesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EmployeesTable, EmployeesColumn),
	)
}
func newSupervisorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SupervisorTable, SupervisorColumn),
	)
}
func newWorkShiftsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkShiftsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkShiftsTable, WorkShiftsColumn),
	)
}
func newApprovedWorkShiftsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApprovedWorkShiftsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ApprovedWorkShiftsTable, ApprovedWorkShiftsColumn),
	)
}
func newAssignedTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssignedTasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AssignedTasksTable, AssignedTasksPrimaryKey...),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Gender) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Gender) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Gender(str)
	if err := GenderValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}
