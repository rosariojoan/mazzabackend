package inventory

type invoice struct {
	Date            string
	Filename        string
	Keywords        string
	Number          string
	PaymentDetails  paymentDetails
	Title           string
	IssuerDetails   invoiceIssuer
	CustomerDetails invoiceCustomer
	Body            [][]string
	Totals          invoiceTotals
}

type invoiceIssuer struct {
	Name    string
	TaxID   string
	Address string
	City    string
	Country string
	Phone   string
	Email   string
}

type invoiceCustomer struct {
	Name    string
	TaxID   string
	Address string
	City    string
	Country string
	Phone   string
	Email   string
}

type invoiceTotals struct {
	Subtotal string
	VATRate  string
	VAT      string
	Total    string
}

type paymentDetails struct {
	BankName      string
	AccountNumber string
	Iban          string
	DueDate       string
}
