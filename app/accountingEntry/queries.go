package accountingentry

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/mazza/generated/model"
)

func AggregateExpenses(ctx context.Context, client generated.AccountingEntryClient, companyID int, timeRange model.TimeRange) ([]*model.ExpensesBreakdownOutput, error) {
	sqlStr := `
	SELECT
		category,
		SUM(
			CASE
				WHEN is_debit = true THEN amount
				ELSE -amount
			END
		) as total_amount
	FROM accounting_entries
	WHERE
		company_accounting_entries = $1  -- Only select the accounting entries from the given company
		AND account_type = $2
		AND date >= DATE_TRUNC($3, CURRENT_DATE)
		AND date <= CURRENT_DATE
	GROUP BY
		category
	ORDER BY
		total_amount DESC;
	`

	rows, err := client.QueryContext(ctx, sqlStr, companyID, accountingentry.AccountTypeEXPENSE, timeRange)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	var scannedRows []*model.ExpensesBreakdownOutput
	defer rows.Close()
	for rows.Next() {
		var item model.ExpensesBreakdownOutput
		if err := rows.Scan(&item.Name, &item.Amount); err != nil {
			fmt.Println("err:", err)
			return nil, err
		}
		scannedRows = append(scannedRows, &item)
	}

	return scannedRows, nil
	// sevenDaysAgo := time.Now().AddDate(0, 0, -7)
}

func AggregateRevenue(ctx context.Context, client generated.AccountingEntryClient, companyID int, timeRange model.TimeRange) ([]*model.RevenueTrendOutput, error) {
	var sqlStr string
	if timeRange == model.TimeRangeWeek || timeRange == model.TimeRangeMonth {
		sqlStr = `
			SELECT
				date::date as date,
				SUM(
					CASE
						WHEN is_debit = false THEN amount
						ELSE -amount
					END
				) as total_amount
			FROM accounting_entries
			WHERE
				company_accounting_entries = $1  -- Only select the accounting entries from the given company
				AND account_type = $2
				AND date >= DATE_TRUNC($3, CURRENT_DATE)
				AND date <= CURRENT_DATE
			GROUP BY
				date::date
			ORDER BY
				total_amount DESC;
		`
	} else {
		sqlStr = `
			SELECT
				DATE_TRUNC('month', date) as date,
				SUM(
					CASE
						WHEN is_debit = false THEN amount
						ELSE -amount
					END
				) as total_amount
			FROM accounting_entries
			WHERE
				company_accounting_entries = $1  -- Only select the accounting entries from the given company
				AND account_type = $2
				AND date >= DATE_TRUNC($3, CURRENT_DATE)
				AND date <= CURRENT_DATE
			GROUP BY
				DATE_TRUNC('month', date)
			ORDER BY
				total_amount DESC;
		`
	}

	rows, err := client.QueryContext(ctx, sqlStr, companyID, accountingentry.AccountTypeREVENUE, timeRange)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	var scannedRows []*model.RevenueTrendOutput
	defer rows.Close()
	for rows.Next() {
		var item model.RevenueTrendOutput
		if err := rows.Scan(&item.Date, &item.Amount); err != nil {
			fmt.Println("err:", err)
			return nil, err
		}
		scannedRows = append(scannedRows, &item)
	}

	return scannedRows, nil
}
