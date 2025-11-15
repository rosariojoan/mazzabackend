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

// type data struct {
// 	Account     string
// 	AccountType string
// 	Balance     float64
// }

func BuildReport(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, endDate time.Time) (output *model.FileDetailsOutput, err error) {
	startDate := utils.StartOfYear(endDate)
	result, err := GetIncomeStatement(client, ctx, user, currentCompany, startDate, endDate)
	if err != nil {
		return nil, err
	}

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

func GetIncomeStatement(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, startDate time.Time, endDate time.Time) (result *model.IncomeStatementOuput, err error) {
	country, lang := "MZ", "PT"

	result = &model.IncomeStatementOuput{
		IsProvisional: true,
		Revenues:      []*model.IncomeStatementRowItem{},
		Expenses:      []*model.IncomeStatementRowItem{},
		Period: &model.Period{
			Start: startDate,
			End:   endDate,
		},
	}

	accountNames, err := utils.LoadAccountNames(country, lang)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve income statement now")
	}

	sqlStr := fmt.Sprintf(`
		SELECT 
			account_type, 
			account, 
			category, 
			sum(balance) AS balance
		FROM (
			SELECT 
				account_type, 
				LEFT (account, 3) AS account, 
				category, 
				sum(amount) AS balance
			FROM 
				accounting_entries
			WHERE
				company_accounting_entries = %d 
				AND date >= '%s' 
				AND date <= '%s' 
				AND account_type IN ('%s', '%s', '%s', '%s', '%s', '%s', '%s')
			GROUP BY 
				account_type, 
				account, 
				category
			ORDER BY 
				account ASC
		) AS summary
		GROUP BY 
			account_type, 
			account, 
			category
		`,
		currentCompany.ID,
		startDate.Format(time.RFC3339), // Convert time to RFC3339 format which PostgreSQL accepts
		endDate.Format(time.RFC3339),
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

	// var scannedRows []data
	var scannedRows []model.IncomeStatementRowItem
	defer rows.Close()
	for rows.Next() {
		var item model.IncomeStatementRowItem
		if err := rows.Scan(&item.AccountType, &item.Account, &item.Category, &item.Value); err != nil {
			// Check for a scan error. Query rows will be closed with defer.
			fmt.Println("err:", err)
			return nil, err
		}
		scannedRows = append(scannedRows, item)
	}

	for _, entry := range scannedRows {
		if entry.AccountType == accountingentry.AccountTypeREVENUE.String() {
			result.Revenues = append(result.Revenues, &model.IncomeStatementRowItem{
				Account:     entry.Account,
				AccountType: entry.AccountType,
				Category:    entry.Category,
				Label:       accountNames[entry.Account],
				Value:       entry.Value,
			})
			result.NetRevenue += entry.Value
		} else if entry.AccountType == accountingentry.AccountTypeCONTRA_REVENUE.String() {
			result.Revenues = append(result.Revenues, &model.IncomeStatementRowItem{
				Account:     entry.Account,
				AccountType: entry.AccountType,
				Category:    entry.Category,
				Label:       accountNames[entry.Account],
				Value:       -entry.Value,
			})
			result.NetRevenue -= entry.Value

		} else if entry.AccountType == accountingentry.AccountTypeEXPENSE.String() {
			result.Expenses = append(result.Expenses, &model.IncomeStatementRowItem{
				Account:     entry.Account,
				AccountType: entry.AccountType,
				Category:    entry.Category,
				Label:       accountNames[entry.Account],
				Value:       entry.Value,
			})
			result.TotalExpenses += entry.Value
		} else if entry.AccountType == accountingentry.AccountTypeCONTRA_EXPENSE.String() {
			result.TotalExpenses -= entry.Value

		} else if entry.AccountType == accountingentry.AccountTypeTAX_EXPENSE.String() {
			result.TaxExpense += entry.Value
		}
	}

	if result.Revenues == nil {
		result.Revenues = []*model.IncomeStatementRowItem{} // empty slice
	}
	if result.Expenses == nil {
		result.Expenses = []*model.IncomeStatementRowItem{} // empty slice
	}

	result.EarningsBeforeTax = result.NetRevenue - result.TotalExpenses
	result.NetIncome = result.EarningsBeforeTax - result.TaxExpense
	// utils.PP(result)
	// fmt.Println("date:", endDate.Format(time.RFC3339))
	return result, err
}
