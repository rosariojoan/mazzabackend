package suppliers

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func GetSupplierList(ctx context.Context, client *generated.Client, top *int) ([]*model.SupplierList, error) {
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
		payables
	WHERE
		company_payables = %d
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

	var supplierList []*model.SupplierList
	defer rows.Close()
	for rows.Next() {
		var item model.SupplierList
		if err := rows.Scan(&item.Name, &item.OutstandingBalance, &item.InvoiceCount); err != nil {
			// Check for a scan error. Query rows will be closed with defer.
			fmt.Println("err:", err)
			return nil, err
		}
		supplierList = append(supplierList, &item)
	}

	return supplierList, err
}
