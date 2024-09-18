// Code generated by ent, DO NOT EDIT.

package payable

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the payable type in the database.
	Label = "payable"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntryGroup holds the string denoting the entrygroup field in the database.
	FieldEntryGroup = "entry_group"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldOutstandingBalance holds the string denoting the outstandingbalance field in the database.
	FieldOutstandingBalance = "outstanding_balance"
	// FieldTotalTransaction holds the string denoting the totaltransaction field in the database.
	FieldTotalTransaction = "total_transaction"
	// FieldDaysDue holds the string denoting the daysdue field in the database.
	FieldDaysDue = "days_due"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeSupplier holds the string denoting the supplier edge name in mutations.
	EdgeSupplier = "supplier"
	// Table holds the table name of the payable in the database.
	Table = "payables"
	// SupplierTable is the table that holds the supplier relation/edge.
	SupplierTable = "payables"
	// SupplierInverseTable is the table name for the Supplier entity.
	// It exists in this package in order to avoid circular dependency with the "supplier" package.
	SupplierInverseTable = "suppliers"
	// SupplierColumn is the table column denoting the supplier relation/edge.
	SupplierColumn = "supplier_payables"
)

// Columns holds all SQL columns for payable fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntryGroup,
	FieldDate,
	FieldOutstandingBalance,
	FieldTotalTransaction,
	FieldDaysDue,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payables"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"supplier_payables",
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
	// EntryGroupValidator is a validator for the "entryGroup" field. It is called by the builders before save.
	EntryGroupValidator func(int) error
	// DaysDueValidator is a validator for the "daysDue" field. It is called by the builders before save.
	DaysDueValidator func(int) error
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusPaid     Status = "paid"
	StatusUnpaid   Status = "unpaid"
	StatusDoubtful Status = "doubtful"
	StatusDefault  Status = "default"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPaid, StatusUnpaid, StatusDoubtful, StatusDefault:
		return nil
	default:
		return fmt.Errorf("payable: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Payable queries.
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

// ByEntryGroup orders the results by the entryGroup field.
func ByEntryGroup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntryGroup, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByOutstandingBalance orders the results by the outstandingBalance field.
func ByOutstandingBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutstandingBalance, opts...).ToFunc()
}

// ByTotalTransaction orders the results by the totalTransaction field.
func ByTotalTransaction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalTransaction, opts...).ToFunc()
}

// ByDaysDue orders the results by the daysDue field.
func ByDaysDue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDaysDue, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySupplierField orders the results by supplier field.
func BySupplierField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSupplierStep(), sql.OrderByField(field, opts...))
	}
}
func newSupplierStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SupplierInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SupplierTable, SupplierColumn),
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
