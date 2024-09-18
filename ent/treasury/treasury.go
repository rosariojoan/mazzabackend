// Code generated by ent, DO NOT EDIT.

package treasury

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the treasury type in the database.
	Label = "treasury"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAccountNumber holds the string denoting the accountnumber field in the database.
	FieldAccountNumber = "account_number"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldBankName holds the string denoting the bankname field in the database.
	FieldBankName = "bank_name"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIban holds the string denoting the iban field in the database.
	FieldIban = "iban"
	// FieldIsDefault holds the string denoting the isdefault field in the database.
	FieldIsDefault = "is_default"
	// FieldIsMainAccount holds the string denoting the ismainaccount field in the database.
	FieldIsMainAccount = "is_main_account"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldSwiftCode holds the string denoting the swiftcode field in the database.
	FieldSwiftCode = "swift_code"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeCashMovements holds the string denoting the cashmovements edge name in mutations.
	EdgeCashMovements = "cashMovements"
	// Table holds the table name of the treasury in the database.
	Table = "treasuries"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "treasuries"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_treasuries"
	// CashMovementsTable is the table that holds the cashMovements relation/edge.
	CashMovementsTable = "cash_movements"
	// CashMovementsInverseTable is the table name for the CashMovement entity.
	// It exists in this package in order to avoid circular dependency with the "cashmovement" package.
	CashMovementsInverseTable = "cash_movements"
	// CashMovementsColumn is the table column denoting the cashMovements relation/edge.
	CashMovementsColumn = "treasury_cash_movements"
)

// Columns holds all SQL columns for treasury fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAccountNumber,
	FieldBalance,
	FieldBankName,
	FieldCurrency,
	FieldDescription,
	FieldIban,
	FieldIsDefault,
	FieldIsMainAccount,
	FieldName,
	FieldCategory,
	FieldSwiftCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "treasuries"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_treasuries",
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
	// DefaultIsDefault holds the default value on creation for the "isDefault" field.
	DefaultIsDefault bool
	// DefaultIsMainAccount holds the default value on creation for the "isMainAccount" field.
	DefaultIsMainAccount bool
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// Currency defines the type for the "currency" enum field.
type Currency string

// Currency values.
const (
	CurrencyMzn Currency = "mzn"
)

func (c Currency) String() string {
	return string(c)
}

// CurrencyValidator is a validator for the "currency" field enum values. It is called by the builders before save.
func CurrencyValidator(c Currency) error {
	switch c {
	case CurrencyMzn:
		return nil
	default:
		return fmt.Errorf("treasury: invalid enum value for currency field: %q", c)
	}
}

// Category defines the type for the "category" enum field.
type Category string

// Category values.
const (
	CategoryDeposit Category = "deposit"
	CategoryCash    Category = "cash"
)

func (c Category) String() string {
	return string(c)
}

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c Category) error {
	switch c {
	case CategoryDeposit, CategoryCash:
		return nil
	default:
		return fmt.Errorf("treasury: invalid enum value for category field: %q", c)
	}
}

// OrderOption defines the ordering options for the Treasury queries.
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

// ByAccountNumber orders the results by the accountNumber field.
func ByAccountNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountNumber, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByBankName orders the results by the bankName field.
func ByBankName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankName, opts...).ToFunc()
}

// ByCurrency orders the results by the currency field.
func ByCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrency, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIban orders the results by the iban field.
func ByIban(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIban, opts...).ToFunc()
}

// ByIsDefault orders the results by the isDefault field.
func ByIsDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDefault, opts...).ToFunc()
}

// ByIsMainAccount orders the results by the isMainAccount field.
func ByIsMainAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsMainAccount, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// BySwiftCode orders the results by the swiftCode field.
func BySwiftCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSwiftCode, opts...).ToFunc()
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}

// ByCashMovementsCount orders the results by cashMovements count.
func ByCashMovementsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCashMovementsStep(), opts...)
	}
}

// ByCashMovements orders the results by cashMovements terms.
func ByCashMovements(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCashMovementsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
func newCashMovementsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CashMovementsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CashMovementsTable, CashMovementsColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Currency) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Currency) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Currency(str)
	if err := CurrencyValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Currency", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Category) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Category) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Category(str)
	if err := CategoryValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}
