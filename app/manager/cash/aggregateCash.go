package cash

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
	"time"
)

func AggregateCash(ctx context.Context, client *generated.Client, input *model.AggregateCashInput) (*model.CashAggregationOutput, error) {
	_, currentCompany := utils.GetSession(&ctx)
	var zeroValue = 0.0
	var result = model.CashAggregationOutput{
		Inflow:  &zeroValue,
		Outflow: &zeroValue,
		Balance: &zeroValue,
	}
	var startDate time.Time
	var endDate = input.EndDate

	if input.StartDate != nil {
		startDate = *input.StartDate
	} else {
		startDate = time.Now().AddDate(-20, 0, 0)
	}

	sqlStr := fmt.Sprintf(`
		SELECT
			sum(debit) AS inflow,
			sum(credit) AS outflow,
			sum(debit) - sum(credit) AS balance
		FROM (
			SELECT 
				CASE WHEN is_debit THEN amount ELSE 0 END AS debit,
				CASE WHEN is_debit THEN 0 ELSE amount END AS credit
			FROM accounting_entries
			WHERE
				company_accounting_entries = %d AND
				date > '%s' AND
				date < '%s' AND
				category = 'cash'
			GROUP BY is_debit, amount
		) AS result
	`, currentCompany.ID, startDate.Format(time.RFC3339), endDate.Format(time.RFC3339))

	rows, err := client.QueryContext(ctx, sqlStr)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var item model.CashAggregationOutput
		if err := rows.Scan(&item.Inflow, &item.Outflow, &item.Balance); err != nil {
			// Check for a scan error. Query rows will be closed with defer.
			return nil, err
		}

		if true {
			value := *result.Inflow + *item.Inflow
			result.Inflow = &value
		}
		if true {
			value := *result.Outflow + *item.Outflow
			result.Outflow = &value
		}
		if true {
			value := *result.Balance + *item.Balance
			result.Balance = &value
		}
	}

	return &result, err
}
