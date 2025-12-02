package finance

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func LoanCounterpartySummary(ctx context.Context, client *generated.Client, lending bool, top *int) ([]*model.LoanCounterpartySummary, error) {
	_, activeCompany := utils.GetSession(&ctx)
	if top == nil {
		limit := 500
		top = &limit
	}
	_ = activeCompany

	sqlStr := `
	SELECT 
		LOWER(counterparty_name) AS name,
		SUM(outstanding_balance) AS outstandingBalance,
		SUM(amount) AS totalBorrowed,
		COUNT(*) AS loansCount,
		AVG(interest_rate) AS averageInterestRate
	FROM
		loans
	WHERE
		company_loans = $1
		AND outstanding_balance > 0
		AND is_lending = $2
	GROUP BY LOWER(counterparty_name)
	ORDER BY outstandingBalance DESC
	LIMIT $3;
	`

	rows, err := client.QueryContext(ctx, sqlStr, activeCompany.ID, lending, *top)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	var scannedRows []*model.LoanCounterpartySummary
	defer rows.Close()
	for rows.Next() {
		var item model.LoanCounterpartySummary
		if err := rows.Scan(&item.Name, &item.OutstandingBalance, &item.TotalBorrowed, &item.LoansCount, &item.AverageInterestRate); err != nil {
			fmt.Println("err:", err)
			return nil, err
		}
		scannedRows = append(scannedRows, &item)
	}

	return scannedRows, nil
}
