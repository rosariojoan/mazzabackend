package balancesheet

import (
	"context"
	"fmt"
	incomestatement "mazza/app/financial-reports/income-statement"
	trialbalance "mazza/app/financial-reports/trial-balance"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/mazza/generated/model"
	"strings"
	"time"
)

func BuildReport(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, endDate time.Time) (output *model.FileDetailsOutput, err error) {
	result, err := GetBalanceSheet(client, ctx, user, currentCompany, endDate)
	if err != nil {
		return nil, err
	}

	startDate := utils.StartOfYear(endDate)
	file, _, err := generatePDFDoc(result, &currentCompany, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	output = &model.FileDetailsOutput{
		Message: "Balanço",
		File: &model.FileOutput{
			Encoding: "base64",
			Kind:     "application/pdf",
			Name:     "balanco.pdf",
			Data:     *file,
		},
	}

	return output, err
}

func GetBalanceSheet(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, endDate time.Time) (result *model.BalanceSheetOuput, err error) {
	country, _ := "MZ", "pt"

	excludedTypes := []string{
		string(accountingentry.AccountTypeREVENUE),
		string(accountingentry.AccountTypeCONTRA_REVENUE),
		string(accountingentry.AccountTypeEXPENSE),
		string(accountingentry.AccountTypeCONTRA_EXPENSE),
		string(accountingentry.AccountTypeTAX_EXPENSE),
		string(accountingentry.AccountTypeINCOME),
		string(accountingentry.AccountTypeDIVIDEND_EXPENSE),
	}
	balanceSheetAccounts, err := trialbalance.GetTrialBalance(client, ctx, user, currentCompany, endDate, excludedTypes)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve balance sheet now")
	}

	incomeStatement, err := incomestatement.GetIncomeStatement(client, ctx, user, currentCompany, endDate)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve balance sheet now")
	}

	initials := balanceSheetAccountInitials(country)
	result = &model.BalanceSheetOuput{
		IsProvisional: true,
		Period: &model.Period{
			Start: utils.StartOfYear(endDate),
			End: endDate,
		},
		Assets: &model.Assets{
			CurrentAssets: []*model.ReportRowItem{},
			FixedAssets: []*model.ReportRowItem{},
		},
		Liabilities: &model.Liabilities{
			CurrentLiabilities: []*model.ReportRowItem{},
			NonCurrentLiabilities: []*model.ReportRowItem{},
		},
		Equity: &model.Equity{
			Equity: []*model.ReportRowItem{},
		},
	}

	for _, entry := range balanceSheetAccounts {
		if utils.StartsWith(entry.Account, initials.CurrentAssets) {
			result.Assets.CurrentAssets = append(result.Assets.CurrentAssets, &model.ReportRowItem{Account: entry.Account, Label: entry.Label, Value: entry.Balance})
			result.Assets.TotalCurrentAssets += entry.Balance
			result.Assets.TotalAssets += entry.Balance

		} else if utils.StartsWith(entry.Account, initials.FixedAssets) {
			result.Assets.FixedAssets = append(result.Assets.FixedAssets, &model.ReportRowItem{Account: entry.Account, Label: entry.Label, Value: entry.Balance})
			result.Assets.TotalFixedAssets += entry.Balance
			result.Assets.TotalAssets += entry.Balance

		} else if utils.StartsWith(entry.Account, initials.CurrentLiabilities) {
			result.Liabilities.CurrentLiabilities = append(result.Liabilities.CurrentLiabilities, &model.ReportRowItem{Account: entry.Account, Label: entry.Label, Value: entry.Balance})
			result.Liabilities.TotalCurrentLiabilities += entry.Balance
			result.Liabilities.TotalLiabilities += entry.Balance
			result.TotalLiabilityAndEquity += entry.Balance

		} else if utils.StartsWith(entry.Account, initials.NonCurrentLiabilities) {
			result.Liabilities.NonCurrentLiabilities = append(result.Liabilities.NonCurrentLiabilities, &model.ReportRowItem{Account: entry.Account, Label: entry.Label, Value: entry.Balance})
			result.Liabilities.TotalNonCurrentLiabilities += entry.Balance
			result.Liabilities.TotalLiabilities += entry.Balance
			result.TotalLiabilityAndEquity += entry.Balance

		} else if utils.StartsWith(entry.Account, initials.Equity) {
			result.Equity.Equity = append(result.Equity.Equity, &model.ReportRowItem{Account: entry.Account, Label: entry.Label, Value: entry.Balance})
			result.Equity.TotalEquity += entry.Balance
			result.TotalLiabilityAndEquity += entry.Balance
		}
	}

	// Adjust retained earnings on the balance sheet to include the provisional net income
	if incomeStatement.NetIncome != 0 {
		retainedEarningsAccount := "5.9"
		idx := -1
		for i, entry := range result.Equity.Equity {
			if strings.HasPrefix(entry.Account, retainedEarningsAccount) {
				idx = i
				break
			}
		}

		if idx >= 0 {
			result.Equity.Equity[idx].Value += incomeStatement.NetIncome
		} else {
			result.Equity.Equity = append(result.Equity.Equity, &model.ReportRowItem{
				Account: retainedEarningsAccount,
				Label:   "Resultados retidos",
				Value:   incomeStatement.NetIncome,
			})
		}

		result.Equity.TotalEquity += incomeStatement.NetIncome
		result.TotalLiabilityAndEquity += incomeStatement.NetIncome
	}

	return result, nil
}
