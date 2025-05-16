package finance

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func LoanProviderList(ctx context.Context, client *generated.Client, top *int) ([]*model.LoanProviderList, error) {
	_, activeCompany := utils.GetSession(&ctx)
	if top == nil {
		limit := 500
		top = &limit
	}
	_ = activeCompany

	sqlStr := `
	SELECT 
		LOWER(provider) AS name,
		SUM(outstanding_balance) AS outstandingBalance,
		SUM(amount) AS totalBorrowed,
		COUNT(*) AS loansCount,
		AVG(interest_rate) AS averageInterestRate
	FROM
		loans
	WHERE
		company_loans = $1
		AND outstanding_balance > 0
	GROUP BY LOWER(provider)
	ORDER BY outstandingBalance DESC
	LIMIT $2;
	`

	rows, err := client.QueryContext(ctx, sqlStr, activeCompany.ID, *top)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	var scannedRows []*model.LoanProviderList
	defer rows.Close()
	for rows.Next() {
		var item model.LoanProviderList
		if err := rows.Scan(&item.Name, &item.OutstandingBalance, &item.TotalBorrowed, &item.LoansCount, &item.AverageInterestRate); err != nil {
			fmt.Println("err:", err)
			return nil, err
		}
		scannedRows = append(scannedRows, &item)
	}

	return scannedRows, nil
}
