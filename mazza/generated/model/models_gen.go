// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"strconv"
	"time"
)

type Assets struct {
	CurrentAssets      []*ReportRowItem `json:"currentAssets"`
	TotalCurrentAssets float64          `json:"totalCurrentAssets"`
	FixedAssets        []*ReportRowItem `json:"fixedAssets"`
	TotalFixedAssets   float64          `json:"totalFixedAssets"`
	TotalAssets        float64          `json:"totalAssets"`
}

type BalanceSheetOuput struct {
	Assets                  *Assets      `json:"assets"`
	Liabilities             *Liabilities `json:"liabilities"`
	Equity                  *Equity      `json:"equity"`
	TotalLiabilityAndEquity float64      `json:"totalLiabilityAndEquity"`
	Period                  *Period      `json:"period"`
	IsProvisional           bool         `json:"isProvisional"`
}

type CustomerAggregationOutput struct {
	Company *int `json:"company,omitempty"`
	Count   *int `json:"count,omitempty"`
}

type EntryCustomerInput struct {
	ID      int     `json:"id"`
	Amount  float64 `json:"amount"`
	DaysDue int     `json:"daysDue"`
}

type EntryItem struct {
	Account     string                      `json:"account"`
	AccountType accountingentry.AccountType `json:"accountType"`
	Amount      float64                     `json:"amount"`
	IsDebit     bool                        `json:"isDebit"`
	Label       string                      `json:"label"`
	Quantity    *int                        `json:"quantity,omitempty"`
	ProductID   *int                        `json:"productID,omitempty"`
	TreasuryID  *int                        `json:"treasuryID,omitempty"`
}

type EntryProductInput struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Quantity int     `json:"quantity"`
}

type EntrySupplierInput struct {
	ID      int     `json:"id"`
	Amount  float64 `json:"amount"`
	DaysDue int     `json:"daysDue"`
}

type Equity struct {
	Equity      []*ReportRowItem `json:"equity"`
	TotalEquity float64          `json:"totalEquity"`
}

type FileDetailsOutput struct {
	Message string      `json:"message"`
	File    *FileOutput `json:"file,omitempty"`
}

type FileOutput struct {
	Encoding string `json:"encoding"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Data     string `json:"data"`
}

type IncomeStatementOuput struct {
	Revenues          []*ReportRowItem `json:"revenues"`
	NetRevenue        float64          `json:"netRevenue"`
	Expenses          []*ReportRowItem `json:"expenses"`
	TotalExpenses     float64          `json:"totalExpenses"`
	EarningsBeforeTax float64          `json:"earningsBeforeTax"`
	TaxExpense        float64          `json:"taxExpense"`
	NetIncome         float64          `json:"netIncome"`
	Period            *Period          `json:"period"`
	IsProvisional     bool             `json:"isProvisional"`
}

type Invoice struct {
	Date            string           `json:"date"`
	Filename        string           `json:"filename"`
	Keywords        string           `json:"keywords"`
	Number          string           `json:"number"`
	PaymentDetails  *PaymentDetails  `json:"paymentDetails"`
	Title           *string          `json:"title,omitempty"`
	IssuerDetails   *InvoiceIssuer   `json:"issuerDetails"`
	CustomerDetails *InvoiceCustomer `json:"customerDetails"`
	Body            [][]string       `json:"body"`
	Totals          *InvoiceTotals   `json:"totals"`
}

type InvoiceCustomer struct {
	Name    string `json:"name"`
	TaxID   string `json:"taxID"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type InvoiceIssuer struct {
	Name    string `json:"name"`
	TaxID   string `json:"taxID"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type InvoiceTotals struct {
	Subtotal string `json:"subtotal"`
	VatRate  string `json:"vatRate"`
	Vat      string `json:"vat"`
	Total    string `json:"total"`
}

type LedgerDownloadInput struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type Liabilities struct {
	CurrentLiabilities         []*ReportRowItem `json:"currentLiabilities"`
	TotalCurrentLiabilities    float64          `json:"totalCurrentLiabilities"`
	NonCurrentLiabilities      []*ReportRowItem `json:"nonCurrentLiabilities"`
	TotalNonCurrentLiabilities float64          `json:"totalNonCurrentLiabilities"`
	TotalLiabilities           float64          `json:"totalLiabilities"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FcmToken string `json:"fcmToken"`
}

type LoginOutput struct {
	User         *generated.User      `json:"user"`
	CompanyID    int                  `json:"companyId"`
	Companies    []*generated.Company `json:"companies"`
	AccessToken  string               `json:"accessToken"`
	RefreshToken string               `json:"refreshToken"`
	TTL          int                  `json:"ttl"`
}

type PaymentDetails struct {
	BankName      string `json:"bankName"`
	AccountNumber string `json:"accountNumber"`
	Iban          string `json:"iban"`
	DueDate       string `json:"dueDate"`
}

type Period struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type PurchaseRegistrationInput struct {
	Main                  []*EntryItem          `json:"main"`
	Counterpart           []*EntryItem          `json:"counterpart"`
	Supplier              *EntrySupplierInput   `json:"supplier"`
	Date                  time.Time             `json:"date"`
	Description           *string               `json:"description,omitempty"`
	OperationType         PurchaseOperationType `json:"operationType"`
	TotalTransactionValue float64               `json:"totalTransactionValue"`
}

type ReceivableAggregationOutput struct {
	Company *int     `json:"company,omitempty"`
	Count   *int     `json:"count,omitempty"`
	Sum     *float64 `json:"sum,omitempty"`
}

type ReportInput struct {
	CompanyID *int      `json:"companyID,omitempty"`
	Date      time.Time `json:"date"`
}

type ReportRowItem struct {
	Account string  `json:"account"`
	Label   string  `json:"label"`
	Value   float64 `json:"value"`
}

type ResetPasswordInput struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type SalesQuotationInput struct {
	Data *Invoice `json:"data"`
}

type SalesRegistrationInput struct {
	Main                  []*EntryItem        `json:"main"`
	Counterpart           []*EntryItem        `json:"counterpart"`
	Customer              *EntryCustomerInput `json:"customer"`
	Date                  time.Time           `json:"date"`
	Description           *string             `json:"description,omitempty"`
	OperationType         SalesOperationType  `json:"operationType"`
	TotalTransactionValue float64             `json:"totalTransactionValue"`
	IssueInvoice          bool                `json:"issueInvoice"`
	Invoice               *Invoice            `json:"invoice,omitempty"`
}

type SignupInput struct {
	CompanyInput *generated.CreateCompanyInput `json:"companyInput"`
	UserInput    *generated.CreateUserInput    `json:"userInput"`
}

type TrialBalanceRowItem struct {
	Account string  `json:"account"`
	Label   string  `json:"label"`
	Debit   float64 `json:"debit"`
	Credit  float64 `json:"credit"`
	Balance float64 `json:"balance"`
}

type WorkShiftAggregationPayload struct {
	Date              string `json:"date"`
	Count             int    `json:"count"`
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"`
	PendingCount      *int   `json:"pendingCount,omitempty"`
}

type InvitedUserSignupInput struct {
	UserInput       *generated.CreateUserInput `json:"userInput"`
	InvitationToken string                     `json:"invitationToken"`
}

type CustomersGroupBy string

const (
	CustomersGroupByCompany CustomersGroupBy = "COMPANY"
	CustomersGroupByCity    CustomersGroupBy = "CITY"
)

var AllCustomersGroupBy = []CustomersGroupBy{
	CustomersGroupByCompany,
	CustomersGroupByCity,
}

func (e CustomersGroupBy) IsValid() bool {
	switch e {
	case CustomersGroupByCompany, CustomersGroupByCity:
		return true
	}
	return false
}

func (e CustomersGroupBy) String() string {
	return string(e)
}

func (e *CustomersGroupBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CustomersGroupBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CustomersGroupBy", str)
	}
	return nil
}

func (e CustomersGroupBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PurchaseOperationType string

const (
	PurchaseOperationTypePurchase       PurchaseOperationType = "PURCHASE"
	PurchaseOperationTypePurchaseReturn PurchaseOperationType = "PURCHASE_RETURN"
)

var AllPurchaseOperationType = []PurchaseOperationType{
	PurchaseOperationTypePurchase,
	PurchaseOperationTypePurchaseReturn,
}

func (e PurchaseOperationType) IsValid() bool {
	switch e {
	case PurchaseOperationTypePurchase, PurchaseOperationTypePurchaseReturn:
		return true
	}
	return false
}

func (e PurchaseOperationType) String() string {
	return string(e)
}

func (e *PurchaseOperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PurchaseOperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PurchaseOperationType", str)
	}
	return nil
}

func (e PurchaseOperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReceivablesGroupBy string

const (
	ReceivablesGroupByAge      ReceivablesGroupBy = "AGE"
	ReceivablesGroupByCompany  ReceivablesGroupBy = "COMPANY"
	ReceivablesGroupByCustomer ReceivablesGroupBy = "CUSTOMER"
	ReceivablesGroupByDaysdue  ReceivablesGroupBy = "DAYSDUE"
	ReceivablesGroupByStatus   ReceivablesGroupBy = "STATUS"
)

var AllReceivablesGroupBy = []ReceivablesGroupBy{
	ReceivablesGroupByAge,
	ReceivablesGroupByCompany,
	ReceivablesGroupByCustomer,
	ReceivablesGroupByDaysdue,
	ReceivablesGroupByStatus,
}

func (e ReceivablesGroupBy) IsValid() bool {
	switch e {
	case ReceivablesGroupByAge, ReceivablesGroupByCompany, ReceivablesGroupByCustomer, ReceivablesGroupByDaysdue, ReceivablesGroupByStatus:
		return true
	}
	return false
}

func (e ReceivablesGroupBy) String() string {
	return string(e)
}

func (e *ReceivablesGroupBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReceivablesGroupBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReceivablesGroupBy", str)
	}
	return nil
}

func (e ReceivablesGroupBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SalesOperationType string

const (
	SalesOperationTypeSales       SalesOperationType = "SALES"
	SalesOperationTypeSalesReturn SalesOperationType = "SALES_RETURN"
)

var AllSalesOperationType = []SalesOperationType{
	SalesOperationTypeSales,
	SalesOperationTypeSalesReturn,
}

func (e SalesOperationType) IsValid() bool {
	switch e {
	case SalesOperationTypeSales, SalesOperationTypeSalesReturn:
		return true
	}
	return false
}

func (e SalesOperationType) String() string {
	return string(e)
}

func (e *SalesOperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SalesOperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SalesOperationType", str)
	}
	return nil
}

func (e SalesOperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShiftGroupBy string

const (
	ShiftGroupByApprovedAt ShiftGroupBy = "APPROVED_AT"
	ShiftGroupByClockIn    ShiftGroupBy = "CLOCK_IN"
	ShiftGroupByClockOut   ShiftGroupBy = "CLOCK_OUT"
	ShiftGroupByStatus     ShiftGroupBy = "STATUS"
)

var AllShiftGroupBy = []ShiftGroupBy{
	ShiftGroupByApprovedAt,
	ShiftGroupByClockIn,
	ShiftGroupByClockOut,
	ShiftGroupByStatus,
}

func (e ShiftGroupBy) IsValid() bool {
	switch e {
	case ShiftGroupByApprovedAt, ShiftGroupByClockIn, ShiftGroupByClockOut, ShiftGroupByStatus:
		return true
	}
	return false
}

func (e ShiftGroupBy) String() string {
	return string(e)
}

func (e *ShiftGroupBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShiftGroupBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShiftGroupBy", str)
	}
	return nil
}

func (e ShiftGroupBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
