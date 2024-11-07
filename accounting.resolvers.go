package mazza

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"encoding/base64"
	"fmt"
	accountingentry "mazza/app/accountingEntry"
	balancesheet "mazza/app/financial-reports/balance-sheet"
	incomestatement "mazza/app/financial-reports/income-statement"
	"mazza/app/financial-reports/ledger"
	trialbalance "mazza/app/financial-reports/trial-balance"
	"mazza/app/manager/inventory"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// RegisterSales is the resolver for the registerSales field.
func (r *mutationResolver) RegisterSales(ctx context.Context, input model.SalesRegistrationInput) (string, error) {
	fmt.Println("##")
	return inventory.SalesRegistration(ctx, r.client, input)
}

// SalesQuotation is the resolver for the salesQuotation field.
func (r *mutationResolver) SalesQuotation(ctx context.Context, input model.SalesQuotationInput) (*model.FileDetailsOutput, error) {
	title := "Cotação"
	input.Data.Title = &title
	fileByte, _, err := inventory.GenerateSalesQuotationPDF(input.Data)
	if err != nil {
		fmt.Println(" ** error:", err)
		return nil, gqlerror.Errorf("error processing document")
	}

	output := &model.FileDetailsOutput{
		Message: "sales quotation",
		File: &model.FileOutput{
			Encoding: "base64",
			Kind:     "application/pdf",
			Name:     "cotacao_de_vendas.pdf",
			Data:     base64.StdEncoding.EncodeToString(fileByte),
		},
	}

	return output, nil
}

// RegisterAccountingEntries is the resolver for the registerAccountingEntries field.
func (r *mutationResolver) RegisterAccountingEntries(ctx context.Context, input model.BaseEntryRegistrationInput) (*string, error) {
	return accountingentry.RegisterAccountingOperations(ctx, r.client, input)
}

// TrialBalance is the resolver for the trialBalance field.
func (r *queryResolver) TrialBalance(ctx context.Context, date time.Time) ([]*model.TrialBalanceRowItem, error) {
	user, company, _ := utils.GetSession(&ctx)
	output, err := trialbalance.GetTrialBalance(r.client, ctx, *user, *company, date, []string{})
	return output, err
}

// IncomeStatement is the resolver for the incomeStatement field.
func (r *queryResolver) IncomeStatement(ctx context.Context, date time.Time) (*model.IncomeStatementOuput, error) {
	user, company, _ := utils.GetSession(&ctx)
	output, err := incomestatement.GetIncomeStatement(r.client, ctx, *user, *company, date)
	return output, err
}

// BalanceSheet is the resolver for the balanceSheet field.
func (r *queryResolver) BalanceSheet(ctx context.Context, date time.Time) (*model.BalanceSheetOuput, error) {
	user, company, _ := utils.GetSession(&ctx)
	output, err := balancesheet.GetBalanceSheet(r.client, ctx, *user, *company, date)
	return output, err
}

// DownloadLedger is the resolver for the downloadLedger field.
func (r *queryResolver) DownloadLedger(ctx context.Context, where model.LedgerDownloadInput) (*model.FileDetailsOutput, error) {
	user, company, _ := utils.GetSession(&ctx)
	output, err := ledger.BuildReport(r.client, ctx, *user, *company, where.StartDate, where.EndDate)
	return output, err
}

// DownloadTrialBalance is the resolver for the downloadTrialBalance field.
func (r *queryResolver) DownloadTrialBalance(ctx context.Context, where model.ReportInput) (*model.FileDetailsOutput, error) {
	user, company, _ := utils.GetSession(&ctx)
	output, err := trialbalance.BuildReport(r.client, ctx, *user, *company, where.Date)
	return output, err
}

// DownloadIncomeStatement is the resolver for the downloadIncomeStatement field.
func (r *queryResolver) DownloadIncomeStatement(ctx context.Context, where model.ReportInput) (*model.FileDetailsOutput, error) {
	user, company, _ := utils.GetSession(&ctx)
	output, err := incomestatement.BuildReport(r.client, ctx, *user, *company, where.Date)
	return output, err
}

// DownloadBalanceSheet is the resolver for the downloadBalanceSheet field.
func (r *queryResolver) DownloadBalanceSheet(ctx context.Context, where model.ReportInput) (*model.FileDetailsOutput, error) {
	user, company, _ := utils.GetSession(&ctx)
	output, err := balancesheet.BuildReport(r.client, ctx, *user, *company, where.Date)
	return output, err
}
