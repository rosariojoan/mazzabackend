package balancesheet

// type item struct {
// 	Account string  `json:"account"`
// 	Label   string  `json:"label"`
// 	Value   float64 `json:"value"`
// }

// type balanceSheet struct {
// 	Assets                  assets      `json:"assets"`
// 	Liabilities             liabilities `json:"liabilities"`
// 	Equity                  equity      `json:"equity"`
// 	TotalLiabilityAndEquity float64     `json:"totalLiabilityAndEquity"`
// 	Period                  period      `json:"period"`
// 	IsProvisional           bool        `json:"isProvisional"`
// }

// type period struct {
// 	Start time.Time `json:"start"`
// 	End   time.Time `json:"end"`
// }

// type assets struct {
// 	CurrentAssets      []item  `json:"currentAssets"`
// 	TotalCurrentAssets float64 `json:"totalCurrentAssets"`
// 	FixedAssets        []item  `json:"fixedAssets"`
// 	TotalFixedAssets   float64 `json:"totalFixedAssets"`
// 	TotalAssets        float64 `json:"totalAssets"`
// }

// type liabilities struct {
// 	CurrentLiabilities         []item  `json:"currentLiabilities"`
// 	TotalCurrentLiabilities    float64 `json:"totalCurrentLiabilities"`
// 	NonCurrentLiabilities      []item  `json:"nonCurrentLiabilities"`
// 	TotalNonCurrentLiabilities float64 `json:"totalNonCurrentLiabilities"`
// 	TotalLiabilities           float64 `json:"totalLiabilities"`
// }

// type equity struct {
// 	Equity      []item  `json:"equity"`
// 	TotalEquity float64 `json:"totalEquity"`
// }

type initials struct {
	CurrentAssets         []string
	FixedAssets           []string
	CurrentLiabilities    []string
	NonCurrentLiabilities []string
	Equity                []string
}

// Get the initials of balance sheet accounts. Initials are used to filter or aggregate accounts
func balanceSheetAccountInitials(country string) initials {
	if country != "MZ" {
		country = "MZ"
	}

	return initials{
		CurrentAssets:         []string{"1", "2", "4.1", "4.5", "4.9.2", "4.9.4"},
		FixedAssets:           []string{"3"},
		CurrentLiabilities:    []string{"4.2", "4.3.1.1", "4.4", "4.6", "4.7", "4.8", "4.9.1", "4.9.3"},
		NonCurrentLiabilities: []string{"4.3.1", "4.3.1.2", "4.3.2", "4.3.3", "4.3.9"},
		Equity:                []string{"5"},
	}
}
