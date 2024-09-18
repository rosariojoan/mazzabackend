// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldFcmToken holds the string denoting the fcmtoken field in the database.
	FieldFcmToken = "fcm_token"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldNotVerified holds the string denoting the notverified field in the database.
	FieldNotVerified = "not_verified"
	// EdgeAccountingEntries holds the string denoting the accountingentries edge name in mutations.
	EdgeAccountingEntries = "accountingEntries"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeAssignedRoles holds the string denoting the assignedroles edge name in mutations.
	EdgeAssignedRoles = "assignedRoles"
	// EdgeCreatedTasks holds the string denoting the createdtasks edge name in mutations.
	EdgeCreatedTasks = "createdTasks"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeTokens holds the string denoting the tokens edge name in mutations.
	EdgeTokens = "tokens"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AccountingEntriesTable is the table that holds the accountingEntries relation/edge.
	AccountingEntriesTable = "accounting_entries"
	// AccountingEntriesInverseTable is the table name for the AccountingEntry entity.
	// It exists in this package in order to avoid circular dependency with the "accountingentry" package.
	AccountingEntriesInverseTable = "accounting_entries"
	// AccountingEntriesColumn is the table column denoting the accountingEntries relation/edge.
	AccountingEntriesColumn = "user_accounting_entries"
	// CompanyTable is the table that holds the company relation/edge. The primary key declared below.
	CompanyTable = "company_users"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// AssignedRolesTable is the table that holds the assignedRoles relation/edge. The primary key declared below.
	AssignedRolesTable = "user_assignedRoles"
	// AssignedRolesInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	AssignedRolesInverseTable = "user_roles"
	// CreatedTasksTable is the table that holds the createdTasks relation/edge.
	CreatedTasksTable = "worktasks"
	// CreatedTasksInverseTable is the table name for the Worktask entity.
	// It exists in this package in order to avoid circular dependency with the "worktask" package.
	CreatedTasksInverseTable = "worktasks"
	// CreatedTasksColumn is the table column denoting the createdTasks relation/edge.
	CreatedTasksColumn = "user_created_tasks"
	// EmployeeTable is the table that holds the employee relation/edge.
	EmployeeTable = "employees"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "user_employee"
	// TokensTable is the table that holds the tokens relation/edge.
	TokensTable = "tokens"
	// TokensInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokensInverseTable = "tokens"
	// TokensColumn is the table column denoting the tokens relation/edge.
	TokensColumn = "user_tokens"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldFcmToken,
	FieldEmail,
	FieldName,
	FieldPassword,
	FieldUsername,
	FieldDisabled,
	FieldNotVerified,
}

var (
	// CompanyPrimaryKey and CompanyColumn2 are the table columns denoting the
	// primary key for the company relation (M2M).
	CompanyPrimaryKey = []string{"company_id", "user_id"}
	// AssignedRolesPrimaryKey and AssignedRolesColumn2 are the table columns denoting the
	// primary key for the assignedRoles relation (M2M).
	AssignedRolesPrimaryKey = []string{"user_id", "user_role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DefaultDisabled holds the default value on creation for the "disabled" field.
	DefaultDisabled bool
	// DefaultNotVerified holds the default value on creation for the "notVerified" field.
	DefaultNotVerified bool
)

// OrderOption defines the ordering options for the User queries.
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

// ByFcmToken orders the results by the fcmToken field.
func ByFcmToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFcmToken, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByDisabled orders the results by the disabled field.
func ByDisabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisabled, opts...).ToFunc()
}

// ByNotVerified orders the results by the notVerified field.
func ByNotVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotVerified, opts...).ToFunc()
}

// ByAccountingEntriesCount orders the results by accountingEntries count.
func ByAccountingEntriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccountingEntriesStep(), opts...)
	}
}

// ByAccountingEntries orders the results by accountingEntries terms.
func ByAccountingEntries(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountingEntriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompanyCount orders the results by company count.
func ByCompanyCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompanyStep(), opts...)
	}
}

// ByCompany orders the results by company terms.
func ByCompany(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAssignedRolesCount orders the results by assignedRoles count.
func ByAssignedRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssignedRolesStep(), opts...)
	}
}

// ByAssignedRoles orders the results by assignedRoles terms.
func ByAssignedRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssignedRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedTasksCount orders the results by createdTasks count.
func ByCreatedTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedTasksStep(), opts...)
	}
}

// ByCreatedTasks orders the results by createdTasks terms.
func ByCreatedTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEmployeeField orders the results by employee field.
func ByEmployeeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmployeeStep(), sql.OrderByField(field, opts...))
	}
}

// ByTokensCount orders the results by tokens count.
func ByTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTokensStep(), opts...)
	}
}

// ByTokens orders the results by tokens terms.
func ByTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAccountingEntriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountingEntriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccountingEntriesTable, AccountingEntriesColumn),
	)
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CompanyTable, CompanyPrimaryKey...),
	)
}
func newAssignedRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssignedRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AssignedRolesTable, AssignedRolesPrimaryKey...),
	)
}
func newCreatedTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedTasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedTasksTable, CreatedTasksColumn),
	)
}
func newEmployeeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmployeeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, EmployeeTable, EmployeeColumn),
	)
}
func newTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TokensTable, TokensColumn),
	)
}
