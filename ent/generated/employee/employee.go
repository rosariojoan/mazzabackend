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
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldDepartment holds the string denoting the department field in the database.
	FieldDepartment = "department"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldHireDate holds the string denoting the hiredate field in the database.
	FieldHireDate = "hire_date"
	// FieldMonthlySalary holds the string denoting the monthlysalary field in the database.
	FieldMonthlySalary = "monthly_salary"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPerformaceScore holds the string denoting the performacescore field in the database.
	FieldPerformaceScore = "performace_score"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSubordinates holds the string denoting the subordinates edge name in mutations.
	EdgeSubordinates = "subordinates"
	// EdgeLeader holds the string denoting the leader edge name in mutations.
	EdgeLeader = "leader"
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
	// SubordinatesTable is the table that holds the subordinates relation/edge.
	SubordinatesTable = "employees"
	// SubordinatesColumn is the table column denoting the subordinates relation/edge.
	SubordinatesColumn = "employee_subordinates"
	// LeaderTable is the table that holds the leader relation/edge.
	LeaderTable = "employees"
	// LeaderColumn is the table column denoting the leader relation/edge.
	LeaderColumn = "employee_subordinates"
)

// Columns holds all SQL columns for employee fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldBirthdate,
	FieldGender,
	FieldPosition,
	FieldDepartment,
	FieldEmail,
	FieldPhone,
	FieldAvatar,
	FieldHireDate,
	FieldMonthlySalary,
	FieldStatus,
	FieldPerformaceScore,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "employees"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_employees",
	"employee_subordinates",
	"user_employee",
}

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
	// DefaultDepartment holds the default value on creation for the "department" field.
	DefaultDepartment string
	// DefaultMonthlySalary holds the default value on creation for the "monthlySalary" field.
	DefaultMonthlySalary int
	// MonthlySalaryValidator is a validator for the "monthlySalary" field. It is called by the builders before save.
	MonthlySalaryValidator func(int) error
	// DefaultPerformaceScore holds the default value on creation for the "performaceScore" field.
	DefaultPerformaceScore float64
	// PerformaceScoreValidator is a validator for the "performaceScore" field. It is called by the builders before save.
	PerformaceScoreValidator func(float64) error
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

// Status defines the type for the "status" enum field.
type Status string

// StatusACTIVE is the default value of the Status enum.
const DefaultStatus = StatusACTIVE

// Status values.
const (
	StatusACTIVE   Status = "ACTIVE"
	StatusON_LEAVE Status = "ON_LEAVE"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusACTIVE, StatusON_LEAVE:
		return nil
	default:
		return fmt.Errorf("employee: invalid enum value for status field: %q", s)
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

// ByBirthdate orders the results by the birthdate field.
func ByBirthdate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthdate, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByPosition orders the results by the position field.
func ByPosition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPosition, opts...).ToFunc()
}

// ByDepartment orders the results by the department field.
func ByDepartment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartment, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByHireDate orders the results by the hireDate field.
func ByHireDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHireDate, opts...).ToFunc()
}

// ByMonthlySalary orders the results by the monthlySalary field.
func ByMonthlySalary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMonthlySalary, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPerformaceScore orders the results by the performaceScore field.
func ByPerformaceScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPerformaceScore, opts...).ToFunc()
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

// BySubordinatesCount orders the results by subordinates count.
func BySubordinatesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubordinatesStep(), opts...)
	}
}

// BySubordinates orders the results by subordinates terms.
func BySubordinates(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubordinatesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLeaderField orders the results by leader field.
func ByLeaderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLeaderStep(), sql.OrderByField(field, opts...))
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
func newSubordinatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubordinatesTable, SubordinatesColumn),
	)
}
func newLeaderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LeaderTable, LeaderColumn),
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
