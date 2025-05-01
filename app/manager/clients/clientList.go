package clients

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func GetClientList(ctx context.Context, client *generated.Client, top *int) ([]*model.ClientList, error) {
	_, activeCompany := utils.GetSession(&ctx)
	if top == nil {
		limit := 500
		top = &limit
	}

	sqlStr := fmt.Sprintf(`
	SELECT 
		LOWER(name) AS name,
		SUM(outstanding_balance) AS outstandingBalance,
		COUNT(*) AS invoiceCount
	FROM
		receivables
	WHERE
		company_receivables = %d
		AND outstanding_balance > 0  -- Only include unpaid invoices
	GROUP BY LOWER(name)
	ORDER BY outstandingBalance DESC
	LIMIT %d;
	`, activeCompany.ID, *top)

	rows, err := client.QueryContext(ctx, sqlStr)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	var clientList []*model.ClientList
	defer rows.Close()
	for rows.Next() {
		var item model.ClientList
		if err := rows.Scan(&item.Name, &item.OutstandingBalance, &item.InvoiceCount); err != nil {
			// Check for a scan error. Query rows will be closed with defer.
			fmt.Println("err:", err)
			return nil, err
		}
		clientList = append(clientList, &item)
	}

	// selector := client.Receivable.Query().Where(
	// 	receivable.HasCompanyWith(company.ID(activeCompany.ID)),
	// 	receivable.StatusNotIn(receivable.StatusPaid, receivable.StatusDefault),
	// ).GroupBy(receivable.FieldName).
	// 	Aggregate(
	// 		generated.As(generated.Count(), "invoiceCount"),
	// 		generated.As(generated.Sum(receivable.FieldOutstandingBalance), "outstandingBalance"),
	// 	).Scan(ctx, &clientList)

	return clientList, err
}
