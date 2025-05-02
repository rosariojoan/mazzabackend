package clients

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
	"time"
)

func AccountsReceivableAging(ctx context.Context, client *generated.Client, name *string) ([]*model.AgingBucket, error) {
	_, currentCompany := utils.GetSession(&ctx)

	query := `
	SELECT 
		age_interval,
		SUM(outstanding_balance) AS total_amount,
		COUNT(*) AS count
	FROM (
		SELECT 
			CASE
				WHEN DATE_PART('day', due_date - $1::TIMESTAMP) < -30 THEN '> 30 days overdue'
				WHEN DATE_PART('day', due_date - $1::TIMESTAMP) < 0 THEN '1-30 days overdue'
				WHEN DATE_PART('day', due_date - $1::TIMESTAMP) <= 7 THEN 'due in 0-7 days'
				WHEN DATE_PART('day', due_date - $1::TIMESTAMP) <= 30 THEN 'due in 8-30 days'
				ELSE 'due in 30+ days'
			END AS age_interval,
			outstanding_balance
		FROM receivables
		WHERE 
			company_receivables = $2
			AND outstanding_balance > 0  -- Only include unpaid invoices
			%s
	) AS subquery
	GROUP BY age_interval
	ORDER BY 
		CASE age_interval
			WHEN '> 30 days overdue' THEN 1
			WHEN '1-30 days overdue' THEN 2
			WHEN 'due in 0-7 days' THEN 3
			WHEN 'due in 8-30 days' THEN 4
			WHEN 'due in 30+ days' THEN 5
		END;
	`

	// Conditionally add the name filter
	var (
		args       []interface{}
		nameClause string
	)

	// Current date for age calculation
	now := time.Now().Format(time.RFC3339) // Convert time to RFC3339 format which PostgreSQL accepts

	args = append(args, now, currentCompany.ID)

	if name != nil && *name != "" {
		nameClause = "AND LOWER(name) = LOWER($3)"
		args = append(args, *name)
	} else {
		nameClause = ""
	}

	// Final query with conditional clause
	finalQuery := fmt.Sprintf(query, nameClause)
	
	rows, err := client.QueryContext(ctx, finalQuery, args...)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	var scannedRows []*model.AgingBucket
	defer rows.Close()
	for rows.Next() {
		var item model.AgingBucket
		if err := rows.Scan(&item.Range, &item.TotalAmount, &item.Count); err != nil {
			// Check for a scan error. Query rows will be closed with defer.
			fmt.Println("err:", err)
			return nil, err
		}
		scannedRows = append(scannedRows, &item)
	}

	// appUtils.PP(scannedRows)
	return scannedRows, nil
}
