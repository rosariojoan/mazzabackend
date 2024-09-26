package ledger

import "time"

type ledgerStruct struct {
	ID uint `json:"id"`
	// CompanyID   uint      `json:"companyId"`
	Number      int       `json:"number"`
	AccountType string    `json:"account_type"`
	Date        time.Time `json:"date"`
	Account     string    `json:"account"`
	Group       int       `json:"group"`
	Label       string    `json:"label"`
	Debit       float64   `json:"debit"`
	Credit      float64   `json:"credit"`
	Description string    `json:"description"`
	IsReversal  bool      `json:"is_reversal"`
	// Files       []string  `json:"files"`
	Reversed bool `json:"reversed"`
}
