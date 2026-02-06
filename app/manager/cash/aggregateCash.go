package cash

import (
	"context"
	"mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
	"time"
)

type DebitCreditSum struct {
	IsDebit bool    `json:"is_debit"`
	Sum     float64 `json:"sum"`
}

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

	var data []DebitCreditSum

	err := client.AccountingEntry.Query().
		Where(
			accountingentry.HasCompanyWith(company.ID(currentCompany.ID)),
			accountingentry.Category("cash"),
			accountingentry.DateGTE(startDate),
			accountingentry.DateLTE(endDate),
		).
		GroupBy(accountingentry.FieldIsDebit).
		Aggregate(generated.Sum(accountingentry.FieldAmount)).
		Scan(ctx, &data)
	if err != nil {
		return &result, nil
	}

	for _, v := range data {
		if v.IsDebit {
			result.Inflow = &v.Sum
		} else {
			result.Outflow = &v.Sum
		}
	}
	balance := *result.Inflow - *result.Outflow
	result.Balance = &balance

	// sqlStr := fmt.Sprintf(`
	// 	SELECT
	// 		sum(debit) AS inflow,
	// 		sum(credit) AS outflow,
	// 		sum(debit) - sum(credit) AS balance
	// 	FROM (
	// 		SELECT
	// 			CASE WHEN is_debit THEN amount ELSE 0 END AS debit,
	// 			CASE WHEN is_debit THEN 0 ELSE amount END AS credit
	// 		FROM accounting_entries
	// 		WHERE
	// 			company_accounting_entries = %d AND
	// 			date > '%s' AND
	// 			date < '%s' AND
	// 			category = 'cash'
	// 		GROUP BY is_debit, amount
	// 	) AS result
	// `, currentCompany.ID, startDate.Format(time.RFC3339), endDate.Format(time.RFC3339))

	// rows, err := client.QueryContext(ctx, sqlStr)
	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()
	// for rows.Next() {
	// 	var item model.CashAggregationOutput
	// 	if err := rows.Scan(&item.Inflow, &item.Outflow, &item.Balance); err != nil {
	// 		// Check for a scan error. Query rows will be closed with defer.
	// 		return nil, err
	// 	}

	// 	var itemInflow float64 = 0
	// 	var itemOutflow float64 = 0
	// 	var itemBalance float64 = 0

	// 	if item.Inflow != nil {
	// 		itemInflow = *item.Inflow
	// 	}
	// 	if item.Outflow != nil {
	// 		itemOutflow = *item.Outflow
	// 	}
	// 	if item.Balance != nil {
	// 		itemBalance = *item.Balance
	// 	}
	// 	if true {
	// 		value := *result.Inflow + itemInflow
	// 		result.Inflow = &value
	// 	}
	// 	if true {
	// 		value := *result.Outflow + itemOutflow
	// 		result.Outflow = &value
	// 	}
	// 	if true {
	// 		value := *result.Balance + itemBalance
	// 		result.Balance = &value
	// 	}
	// }

	return &result, err
}
