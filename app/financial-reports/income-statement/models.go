package incomestatement

import "time"

type item struct {
	Account string  `json:"account"`
	Label   string  `json:"label"`
	Value   float64 `json:"value"`
}

type incomeStatement struct {
	Revenues          []item  `json:"revenues"`
	NetRevenue        float64 `json:"netRevenue"`
	Expenses          []item  `json:"expenses"`
	TotalExpenses      float64 `json:"totalExpenses"`
	EarningsBeforeTax float64 `json:"incomeBeforeTax"`
	TaxExpense        float64 `json:"taxExpense"`
	NetIncome         float64 `json:"netIncome"`
	Period            period  `json:"period"`
	IsProvisional     bool    `json:"isProvisional"`
	// DividendsPaid float64
}

type period struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
