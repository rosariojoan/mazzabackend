// Code generated by ent, DO NOT EDIT.

package worktag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the worktag type in the database.
	Label = "worktag"
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
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeWorkTasks holds the string denoting the worktasks edge name in mutations.
	EdgeWorkTasks = "workTasks"
	// Table holds the table name of the worktag in the database.
	Table = "worktags"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "worktags"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_work_tags"
	// WorkTasksTable is the table that holds the workTasks relation/edge. The primary key declared below.
	WorkTasksTable = "worktask_workTags"
	// WorkTasksInverseTable is the table name for the Worktask entity.
	// It exists in this package in order to avoid circular dependency with the "worktask" package.
	WorkTasksInverseTable = "worktasks"
)

// Columns holds all SQL columns for worktag fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldColor,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "worktags"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_work_tags",
}

var (
	// WorkTasksPrimaryKey and WorkTasksColumn2 are the table columns denoting the
	// primary key for the workTasks relation (M2M).
	WorkTasksPrimaryKey = []string{"worktask_id", "worktag_id"}
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

// OrderOption defines the ordering options for the Worktag queries.
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

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkTasksCount orders the results by workTasks count.
func ByWorkTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkTasksStep(), opts...)
	}
}

// ByWorkTasks orders the results by workTasks terms.
func ByWorkTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
func newWorkTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkTasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WorkTasksTable, WorkTasksPrimaryKey...),
	)
}