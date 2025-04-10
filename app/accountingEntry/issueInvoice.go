package accountingentry

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/invoice"
	"mazza/ent/utils"
	"mazza/firebase"
	"mazza/mazza/generated/model"
)

func IssueInvoice(ctx context.Context, client *generated.Client, input model.InvoiceInput) (*model.InvoiceIssuanceOutput, error) {
	activeUser, activeCompany := utils.GetSession(&ctx)
	// FIRST, CHECK THAT THE INVOICE NUMBER DOES NOT EXIST FOR THE COMPANY
	exists, err := client.Invoice.Query().Where(
		invoice.HasCompanyWith(company.ID(activeCompany.ID)),
		invoice.Number(*input.InvoiceData.Number),
	).Exist(ctx)

	if exists || err != nil {
		fmt.Println("invoice number exists or err:", &input.InvoiceData.Number, *input.InvoiceData.Number, err)
		// delete the invoice file from the storage as it is not valid
		go firebase.DeleteItem(*input.InvoiceData.StorageURI)
		return nil, fmt.Errorf("invalid invoice number. refresh company details")
	}

	// Start db transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println("err:", err)
		// delete the invoice file from the storage as it is not valid
		go firebase.DeleteItem(*input.InvoiceData.StorageURI)
		return nil, fmt.Errorf("an error occurred")
	}

	// Create invoice entry
	invoiceEntry, err := tx.Invoice.Create().SetInput(*input.InvoiceData).
		SetCompanyID(activeCompany.ID).
		SetIssuedByID(activeUser.ID).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println("err:", err)
		// delete the invoice file from the storage as it is not valid
		go firebase.DeleteItem(*input.InvoiceData.StorageURI)
		return nil, fmt.Errorf("an error occurred")
	}

	// Register accounting entries
	result, err := RegisterAccountingOperations(ctx, tx, *input.AccountingEntryData, &invoiceEntry.ID)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println("InvoiceIssuance err:", err)
		// delete the invoice file from the storage as it is not valid
		go firebase.DeleteItem(*input.InvoiceData.StorageURI)
		return nil, fmt.Errorf("an error occurred")
	}

	if input.InventoryMovements != nil {
		// Create inventory movement
		for _, movement := range input.InventoryMovements {
			_, err := CreateInventoryMovement(ctx, tx, *movement, nil)
			if err != nil {
				fmt.Println("InvoiceIssuance err:", err)
				return nil, err
			}
		}
	}

	// Update the number of issued invoices in the company model
	_, err = tx.Company.UpdateOneID(activeCompany.ID).AddLastInvoiceNumber(1).Save(ctx)
	if err != nil {
		fmt.Println("err:", err)
		_ = tx.Rollback()
		// delete the invoice file from the storage as it is not valid
		go firebase.DeleteItem(*input.InvoiceData.StorageURI)
		return nil, fmt.Errorf("an error occurred")
	}

	// Create invoice entry
	// status := invoice.StatusPAID
	// status := *input.InvoiceData.Status
	// if input.InvoiceData.Terms != nil && *input.InvoiceData.Terms > 0 {
	// 	status = invoice.StatusPENDING
	// }

	if err = tx.Commit(); err != nil {
		fmt.Println("err:", err)
		_ = tx.Rollback()
		// delete the invoice file from the storage as it is not valid
		go firebase.DeleteItem(*input.InvoiceData.StorageURI)
		return nil, fmt.Errorf("an error occurred")
	}

	output := model.InvoiceIssuanceOutput{
		Message: *result,
		FileURL: *input.InvoiceData.URL,
	}

	return &output, nil
}
