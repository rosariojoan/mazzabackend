package incomestatement

import (
	"context"
	"fmt"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/mazza/generated/model"
	"time"
)

type data struct {
	Account     string
	AccountType string
	Balance     float64
}

func BuildReport(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, endDate time.Time) (output *model.FileDetailsOutput, err error) {
	result, err := GetIncomeStatement(client, ctx, user, currentCompany, endDate)
	if err != nil {
		return nil, err
	}

	startDate := utils.StartOfYear(endDate)
	file, _, err := generatePDFDoc(result, &currentCompany, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	output = &model.FileDetailsOutput{
		Message: "Demonstração de resultados",
		File: &model.FileOutput{
			Encoding: "base64",
			Kind:     "application/pdf",
			Name:     "dmr.pdf",
			Data:     *file,
		},
	}

	return output, err
}

func GetIncomeStatement(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, endDate time.Time) (result *model.IncomeStatementOuput, err error) {
	country, lang := "MZ", "pt"

	result = &model.IncomeStatementOuput{
		IsProvisional: true,
		Revenues:      []*model.ReportRowItem{},
		Expenses:      []*model.ReportRowItem{},
		Period: &model.Period{
			Start: utils.StartOfYear(endDate),
			End:   endDate,
		},
	}

	accountNames, err := utils.LoadAccountNames(country, lang)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve income statement now")
	}

	sqlStr := fmt.Sprintf(`
		SELECT account_type, LEFT (account, 3) AS account, sum(amount) AS balance
		FROM accounting_entries
		WHERE company_accounting_entries = %d AND date <= '%s' AND account_type IN ('%s', '%s', '%s', '%s', '%s', '%s', '%s')
		GROUP BY account_type, account
		ORDER BY account ASC
		`,
		currentCompany.ID,
		endDate.Format(time.RFC3339), // Convert time to RFC3339 format which PostgreSQL accepts
		accountingentry.AccountTypeREVENUE,
		accountingentry.AccountTypeCONTRA_REVENUE,
		accountingentry.AccountTypeEXPENSE,
		accountingentry.AccountTypeCONTRA_EXPENSE,
		accountingentry.AccountTypeTAX_EXPENSE,
		accountingentry.AccountTypeINCOME,
		accountingentry.AccountTypeDIVIDEND_EXPENSE,
	)

	rows, err := client.QueryContext(ctx, sqlStr)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	var scannedRows []data
	defer rows.Close()
	for rows.Next() {
		var item data
		if err := rows.Scan(&item.Account, &item.AccountType, &item.Balance); err != nil {
			// Check for a scan error. Query rows will be closed with defer.
			fmt.Println("err:", err)
			return nil, err
		}
		scannedRows = append(scannedRows, item)
	}

	for _, entry := range scannedRows {
		if entry.AccountType == accountingentry.AccountTypeREVENUE.String() {
			result.Revenues = append(result.Revenues, &model.ReportRowItem{
				Account: entry.Account,
				Label:   accountNames[entry.Account],
				Value:   entry.Balance,
			})
			result.NetRevenue += entry.Balance
		} else if entry.AccountType == accountingentry.AccountTypeCONTRA_REVENUE.String() {
			result.Revenues = append(result.Revenues, &model.ReportRowItem{
				Account: entry.Account,
				Label:   accountNames[entry.Account],
				Value:   entry.Balance,
			})
			result.NetRevenue -= entry.Balance

		} else if entry.AccountType == accountingentry.AccountTypeEXPENSE.String() {
			result.Expenses = append(result.Expenses, &model.ReportRowItem{
				Account: entry.Account,
				Label:   accountNames[entry.Account],
				Value:   entry.Balance,
			})
			result.TotalExpenses += entry.Balance
		} else if entry.AccountType == accountingentry.AccountTypeCONTRA_EXPENSE.String() {
			result.TotalExpenses -= entry.Balance

		} else if entry.AccountType == accountingentry.AccountTypeTAX_EXPENSE.String() {
			result.TaxExpense += entry.Balance
		}
	}

	if result.Revenues == nil {
		result.Revenues = []*model.ReportRowItem{} // empty slice
	}
	if result.Expenses == nil {
		result.Expenses = []*model.ReportRowItem{} // empty slice
	}

	result.EarningsBeforeTax = result.NetRevenue - result.TotalExpenses
	result.NetIncome = result.EarningsBeforeTax - result.TaxExpense

	return result, err
}
