package accountingentry

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/invoice"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func IssueSalesQuotation(ctx context.Context, client *generated.Client, input model.SalesQuotationInput) (*generated.Invoice, error) {
	activeUser, activeCompany := utils.GetSession(&ctx)
	docNumber := *input.InvoiceData.Number

	// FIRST, CHECK THAT THE DOCUMENT NUMBER DOES NOT EXIST FOR THE COMPANY
	exists, _ := client.Invoice.Query().Where(
		invoice.HasCompanyWith(company.ID(activeCompany.ID)),
		invoice.IsInvoice(false),
		invoice.Number(*input.InvoiceData.Number),
	).Exist(ctx)
	if exists {
		docNumber = fmt.Sprintf("%s-%04d", activeCompany.QuotationPrefix, activeCompany.LastSalesQuotationNumber+1)
	}

	// Start db transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	currency := input.InvoiceData.Currency
	if len(currency) == 0 {
		currency = activeCompany.BaseCurrency
	}

	// Create new entry
	entry, err := tx.Invoice.Create().SetInput(*input.InvoiceData).
		SetCurrency(currency).
		SetIsInvoice(false).
		SetNumber(docNumber).
		SetCompanyID(activeCompany.ID).
		SetIssuedByID(activeUser.ID).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	// Update the number of issued quotations in the company model
	_, err = tx.Company.UpdateOneID(activeCompany.ID).AddLastSalesQuotationNumber(1).Save(ctx)
	if err != nil {
		fmt.Println("err:", err)
		_ = tx.Rollback()
		return nil, fmt.Errorf("an error occurred")
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("err:", err)
		_ = tx.Rollback()
		return nil, fmt.Errorf("an error occurred")
	}

	return entry, nil
}
