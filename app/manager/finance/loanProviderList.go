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
		LOWER(provider) AS provider,
		SUM(outstanding_balance) AS outstandingBalance,
		COUNT(*) AS invoiceCount
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
		// var item model.LoanProviderList
		// if err := rows.Scan(&item.Name, &item.Amount); err != nil {
		// 	fmt.Println("err:", err)
		// 	return nil, err
		// }
		// scannedRows = append(scannedRows, &item)
		fmt.Println(rows)
	}

	_ = scannedRows
	return nil, nil
}
