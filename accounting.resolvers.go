package mazza

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.52

import (
	"context"
	"encoding/json"
	"fmt"
	accountingentry "mazza/app/accountingEntry"
	balancesheet "mazza/app/financial-reports/balance-sheet"
	incomestatement "mazza/app/financial-reports/income-statement"
	"mazza/app/financial-reports/ledger"
	trialbalance "mazza/app/financial-reports/trial-balance"
	appUtils "mazza/app/utils"
	"mazza/ent/generated"
	"mazza/ent/generated/file"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
	"time"
)

// IssueSalesQuotation is the resolver for the issueSalesQuotation field.
func (r *mutationResolver) IssueSalesQuotation(ctx context.Context, input model.SalesQuotationInput) (*model.InvoiceIssuanceOutput, error) {
	title := "Cotação"
	input.InvoiceData.Title = &title
	fileByte, size, err := accountingentry.GenerateSalesQuotationPDF(input.InvoiceData)
	if err != nil {
		fmt.Println(" ** IssueSalesQuotation GenerateSalesQuotationPDF error:", err)
		return nil, fmt.Errorf("error processing document")
	}

	date, err := time.Now().MarshalText()
	if err != nil {
		date = []byte(time.Now().String())
	}
	dateString := string(date)

	_, company := utils.GetSession(&ctx)
	filename := "cotacao_" + dateString + ".pdf"
	fileProps := appUtils.FileProps{
		Bucket:      "quotation",
		CompanyID:   int(company.ID),
		ContentType: appUtils.ContentType.PDF,
		Content:     fileByte,
		Name:        filename,
	}

	uri, fileUrl, err := appUtils.UploadFile(fileProps)
	if err != nil {
		return nil, fmt.Errorf("entry error. quotation could not be issued")
	}

	// Generate file description
	descrip := map[string]any{
		"quotationNumber": dateString,
		"clientName":      input.InvoiceData.CustomerDetails.Name,
		"date":            dateString, // this is deliberatelly used instead of input.Invoice.Date
		"total":           input.InvoiceData.Totals.Total,
	}
	_descrip, _ := json.Marshal(descrip)
	description := string(_descrip)
	fileEntry, err := r.client.File.Create().SetInput(generated.CreateFileInput{
		Category:    file.CategorySALESQUOTATION,
		Extension:   "pdf",
		Size:        appUtils.FloatToString(size),
		URI:         *uri,
		URL:         *fileUrl,
		Description: description,
		CompanyID:   &company.ID,
	}).Save(ctx)
	if err != nil {
		fmt.Println("InvoiceIssuance err:", err)
		return nil, fmt.Errorf("an error occurred")
	}
	_ = fileEntry

	output := model.InvoiceIssuanceOutput{
		Message: "Cotação emitida com sucesso",
		FileURL: *fileUrl,
	}

	return &output, nil
}

// IssueInvoice is the resolver for the issueInvoice field.
func (r *mutationResolver) IssueInvoice(ctx context.Context, input model.InvoiceInput) (*model.InvoiceIssuanceOutput, error) {
	fileByte, size, err := accountingentry.GenerateInvoicePDF(input.InvoiceData)
	if err != nil {
		fmt.Println("InvoiceIssuance GenerateInvoicePDF err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	_, company := utils.GetSession(&ctx)
	filename := "fatura_" + input.InvoiceData.Number + ".pdf"
	fileProps := appUtils.FileProps{
		Bucket:      "invoice",
		CompanyID:   int(company.ID),
		ContentType: appUtils.ContentType.PDF,
		Content:     fileByte,
		Name:        filename,
	}

	uri, fileUrl, err := appUtils.UploadFile(fileProps)
	if err != nil {
		return nil, fmt.Errorf("entry error. invoice could not be issued")
	}

	// Generate file description
	date, err := input.AccountingEntryData.Date.MarshalText()
	if err != nil {
		date = []byte(input.AccountingEntryData.Date.String())
	}

	descrip := map[string]any{
		"invoiceNumber": input.InvoiceData.Number,
		"clientName":    input.InvoiceData.CustomerDetails.Name,
		"date":          string(date), // this is deliberatelly used instead of input.Invoice.Date
		"total":         input.InvoiceData.Totals.Total,
	}
	_descrip, _ := json.Marshal(descrip)
	description := string(_descrip)
	fileEntry, err := r.client.File.Create().SetInput(generated.CreateFileInput{
		Category:    file.CategoryINVOICE,
		Extension:   "pdf",
		Size:        appUtils.FloatToString(size),
		URI:         *uri,
		URL:         *fileUrl,
		Description: description,
		CompanyID:   &company.ID,
	}).Save(ctx)
	if err != nil {
		fmt.Println("InvoiceIssuance err:", err)
		return nil, fmt.Errorf("an error occurred")
	}
	_ = fileEntry

	result, err := r.RegisterAccountingEntries(ctx, *input.AccountingEntryData)
	if err != nil {
		fmt.Println("InvoiceIssuance err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	// Update the number of issued invoices in the company model
	_, err = r.client.Company.UpdateOneID(company.ID).AddLastInvoiceNumber(1).Save(ctx)
	if err != nil {
		fmt.Println("err:", err)
	}

	output := model.InvoiceIssuanceOutput{
		Message: *result,
		FileURL: *fileUrl,
	}

	return &output, nil
}

// RegisterAccountingEntries is the resolver for the registerAccountingEntries field.
func (r *mutationResolver) RegisterAccountingEntries(ctx context.Context, input model.BaseEntryRegistrationInput) (*string, error) {
	return accountingentry.RegisterAccountingOperations(ctx, r.client, input)
}

// TrialBalance is the resolver for the trialBalance field.
func (r *queryResolver) TrialBalance(ctx context.Context, date time.Time) ([]*model.TrialBalanceRowItem, error) {
	user, company := utils.GetSession(&ctx)
	output, err := trialbalance.GetTrialBalance(r.client, ctx, *user, *company, date, []string{})
	return output, err
}

// IncomeStatement is the resolver for the incomeStatement field.
func (r *queryResolver) IncomeStatement(ctx context.Context, date time.Time) (*model.IncomeStatementOuput, error) {
	user, company := utils.GetSession(&ctx)
	output, err := incomestatement.GetIncomeStatement(r.client, ctx, *user, *company, date)
	return output, err
}

// BalanceSheet is the resolver for the balanceSheet field.
func (r *queryResolver) BalanceSheet(ctx context.Context, date time.Time) (*model.BalanceSheetOuput, error) {
	user, company := utils.GetSession(&ctx)
	output, err := balancesheet.GetBalanceSheet(r.client, ctx, *user, *company, date)
	return output, err
}

// DownloadLedger is the resolver for the downloadLedger field.
func (r *queryResolver) DownloadLedger(ctx context.Context, where model.LedgerDownloadInput) (*model.FileDetailsOutput, error) {
	user, company := utils.GetSession(&ctx)
	output, err := ledger.BuildReport(r.client, ctx, *user, *company, where.StartDate, where.EndDate)
	return output, err
}

// DownloadTrialBalance is the resolver for the downloadTrialBalance field.
func (r *queryResolver) DownloadTrialBalance(ctx context.Context, where model.ReportInput) (*model.FileDetailsOutput, error) {
	user, company := utils.GetSession(&ctx)
	output, err := trialbalance.BuildReport(r.client, ctx, *user, *company, where.Date)
	return output, err
}

// DownloadIncomeStatement is the resolver for the downloadIncomeStatement field.
func (r *queryResolver) DownloadIncomeStatement(ctx context.Context, where model.ReportInput) (*model.FileDetailsOutput, error) {
	user, company := utils.GetSession(&ctx)
	output, err := incomestatement.BuildReport(r.client, ctx, *user, *company, where.Date)
	return output, err
}

// DownloadBalanceSheet is the resolver for the downloadBalanceSheet field.
func (r *queryResolver) DownloadBalanceSheet(ctx context.Context, where model.ReportInput) (*model.FileDetailsOutput, error) {
	user, company := utils.GetSession(&ctx)
	output, err := balancesheet.BuildReport(r.client, ctx, *user, *company, where.Date)
	return output, err
}
