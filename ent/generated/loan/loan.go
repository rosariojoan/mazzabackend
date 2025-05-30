// Code generated by ent, DO NOT EDIT.

package loan

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the loan type in the database.
	Label = "loan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldCollateral holds the string denoting the collateral field in the database.
	FieldCollateral = "collateral"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldInterestRate holds the string denoting the interestrate field in the database.
	FieldInterestRate = "interest_rate"
	// FieldInstallments holds the string denoting the installments field in the database.
	FieldInstallments = "installments"
	// FieldMaturityDate holds the string denoting the maturitydate field in the database.
	FieldMaturityDate = "maturity_date"
	// FieldNextPayment holds the string denoting the nextpayment field in the database.
	FieldNextPayment = "next_payment"
	// FieldNextPaymentAmount holds the string denoting the nextpaymentamount field in the database.
	FieldNextPaymentAmount = "next_payment_amount"
	// FieldOutstandingBalance holds the string denoting the outstandingbalance field in the database.
	FieldOutstandingBalance = "outstanding_balance"
	// FieldPaymentFrequency holds the string denoting the paymentfrequency field in the database.
	FieldPaymentFrequency = "payment_frequency"
	// FieldPaidInstallments holds the string denoting the paidinstallments field in the database.
	FieldPaidInstallments = "paid_installments"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldStartDate holds the string denoting the startdate field in the database.
	FieldStartDate = "start_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeTransactionHistory holds the string denoting the transactionhistory edge name in mutations.
	EdgeTransactionHistory = "transactionHistory"
	// Table holds the table name of the loan in the database.
	Table = "loans"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "loans"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_loans"
	// TransactionHistoryTable is the table that holds the transactionHistory relation/edge.
	TransactionHistoryTable = "accounting_entries"
	// TransactionHistoryInverseTable is the table name for the AccountingEntry entity.
	// It exists in this package in order to avoid circular dependency with the "accountingentry" package.
	TransactionHistoryInverseTable = "accounting_entries"
	// TransactionHistoryColumn is the table column denoting the transactionHistory relation/edge.
	TransactionHistoryColumn = "loan_transaction_history"
)

// Columns holds all SQL columns for loan fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAmount,
	FieldCategory,
	FieldCollateral,
	FieldDescription,
	FieldInterestRate,
	FieldInstallments,
	FieldMaturityDate,
	FieldNextPayment,
	FieldNextPaymentAmount,
	FieldOutstandingBalance,
	FieldPaymentFrequency,
	FieldPaidInstallments,
	FieldProvider,
	FieldStartDate,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "loans"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_loans",
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
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(float64) error
	// InterestRateValidator is a validator for the "interestRate" field. It is called by the builders before save.
	InterestRateValidator func(float64) error
	// InstallmentsValidator is a validator for the "installments" field. It is called by the builders before save.
	InstallmentsValidator func(int) error
	// DefaultNextPaymentAmount holds the default value on creation for the "nextPaymentAmount" field.
	DefaultNextPaymentAmount float64
	// NextPaymentAmountValidator is a validator for the "nextPaymentAmount" field. It is called by the builders before save.
	NextPaymentAmountValidator func(float64) error
	// OutstandingBalanceValidator is a validator for the "outstandingBalance" field. It is called by the builders before save.
	OutstandingBalanceValidator func(float64) error
	// DefaultPaidInstallments holds the default value on creation for the "paidInstallments" field.
	DefaultPaidInstallments int
	// ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	ProviderValidator func(string) error
	// DefaultStartDate holds the default value on creation for the "startDate" field.
	DefaultStartDate func() time.Time
)

// Category defines the type for the "category" enum field.
type Category string

// Category values.
const (
	CategoryAuto           Category = "auto"
	CategoryEquipment      Category = "equipment"
	CategoryExpansion      Category = "expansion"
	CategoryLineOfCredit   Category = "line_of_credit"
	CategoryMortgage       Category = "mortgage"
	CategoryPersonal       Category = "personal"
	CategoryWorkingCapital Category = "working_capital"
)

func (c Category) String() string {
	return string(c)
}

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c Category) error {
	switch c {
	case CategoryAuto, CategoryEquipment, CategoryExpansion, CategoryLineOfCredit, CategoryMortgage, CategoryPersonal, CategoryWorkingCapital:
		return nil
	default:
		return fmt.Errorf("loan: invalid enum value for category field: %q", c)
	}
}

// PaymentFrequency defines the type for the "paymentFrequency" enum field.
type PaymentFrequency string

// PaymentFrequencyMonthly is the default value of the PaymentFrequency enum.
const DefaultPaymentFrequency = PaymentFrequencyMonthly

// PaymentFrequency values.
const (
	PaymentFrequencyWeekly     PaymentFrequency = "weekly"
	PaymentFrequencyBiweekly   PaymentFrequency = "biweekly"
	PaymentFrequencyMonthly    PaymentFrequency = "monthly"
	PaymentFrequencyQuartely   PaymentFrequency = "quartely"
	PaymentFrequencySemiannual PaymentFrequency = "semiannual"
	PaymentFrequencyAnnual     PaymentFrequency = "annual"
)

func (pf PaymentFrequency) String() string {
	return string(pf)
}

// PaymentFrequencyValidator is a validator for the "paymentFrequency" field enum values. It is called by the builders before save.
func PaymentFrequencyValidator(pf PaymentFrequency) error {
	switch pf {
	case PaymentFrequencyWeekly, PaymentFrequencyBiweekly, PaymentFrequencyMonthly, PaymentFrequencyQuartely, PaymentFrequencySemiannual, PaymentFrequencyAnnual:
		return nil
	default:
		return fmt.Errorf("loan: invalid enum value for paymentFrequency field: %q", pf)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusActive Status = "active"
	StatusPaid   Status = "paid"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusPaid:
		return nil
	default:
		return fmt.Errorf("loan: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Loan queries.
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

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByCollateral orders the results by the collateral field.
func ByCollateral(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollateral, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByInterestRate orders the results by the interestRate field.
func ByInterestRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInterestRate, opts...).ToFunc()
}

// ByInstallments orders the results by the installments field.
func ByInstallments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstallments, opts...).ToFunc()
}

// ByMaturityDate orders the results by the maturityDate field.
func ByMaturityDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaturityDate, opts...).ToFunc()
}

// ByNextPayment orders the results by the nextPayment field.
func ByNextPayment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNextPayment, opts...).ToFunc()
}

// ByNextPaymentAmount orders the results by the nextPaymentAmount field.
func ByNextPaymentAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNextPaymentAmount, opts...).ToFunc()
}

// ByOutstandingBalance orders the results by the outstandingBalance field.
func ByOutstandingBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutstandingBalance, opts...).ToFunc()
}

// ByPaymentFrequency orders the results by the paymentFrequency field.
func ByPaymentFrequency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentFrequency, opts...).ToFunc()
}

// ByPaidInstallments orders the results by the paidInstallments field.
func ByPaidInstallments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaidInstallments, opts...).ToFunc()
}

// ByProvider orders the results by the provider field.
func ByProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProvider, opts...).ToFunc()
}

// ByStartDate orders the results by the startDate field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}

// ByTransactionHistoryCount orders the results by transactionHistory count.
func ByTransactionHistoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionHistoryStep(), opts...)
	}
}

// ByTransactionHistory orders the results by transactionHistory terms.
func ByTransactionHistory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionHistoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
func newTransactionHistoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionHistoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionHistoryTable, TransactionHistoryColumn),
	)
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

// MarshalGQL implements graphql.Marshaler interface.
func (e PaymentFrequency) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *PaymentFrequency) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = PaymentFrequency(str)
	if err := PaymentFrequencyValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid PaymentFrequency", str)
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
