package trialbalance

type tbStruct struct {
	Account string  `json:"account"`
	Label   string  `json:"label"`
	Debit   float64 `json:"debit"`
	Credit  float64 `json:"credit"`
	Balance float64 `json:"balance"`
}
