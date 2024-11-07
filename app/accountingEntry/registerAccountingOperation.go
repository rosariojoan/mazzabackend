package accountingentry

import (
	"context"
	"fmt"
	ent "mazza/ent/generated"
	"mazza/ent/generated/payable"
	"mazza/ent/generated/receivable"
	"mazza/ent/generated/treasury"
	u "mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func RegisterAccountingOperations(ctx context.Context, client *ent.Client, input model.BaseEntryRegistrationInput) (*string, error) {
	currentUser, currentCompany, _ := u.GetSession(&ctx)
	companyQ := u.CurrentCompanyQuery(&ctx)
	// country, lang := "mz", "pt"
	// accountNames, err := utils.LoadAccountNames(country, lang)
	// if err != nil {
	// 	return nil, err
	// }
	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	if input.CashInput != nil {
		_ = tx.Treasury.Update().Where(treasury.HasCompanyWith(companyQ)).AddBalance(*input.CashInput).SaveX(ctx)
	}

	if input.ProductInput != nil {
		_ = tx.Product.Update().AddStock(*input.ProductInput).SaveX(ctx)
	}

	var accountingEntries []*ent.AccountingEntryCreate
	var entryCounter = GetEntryCounter(ctx, client, currentCompany.ID)

	if input.ReceivableInput != nil {
		_ = tx.Receivable.Create().SetInput(ent.CreateReceivableInput{
			CompanyID:          &currentCompany.ID,
			EntryGroup:         entryCounter.Group,
			Date:               input.Date,
			OutstandingBalance: input.ReceivableInput.Amount,
			TotalTransaction:   input.ReceivableInput.Amount,
			DueDate:            input.ReceivableInput.DueDate,
			Status:             receivable.StatusPending,
		}).SaveX(ctx)
	}

	if input.PayableInput != nil {
		_ = tx.Payable.Create().SetInput(ent.CreatePayableInput{
			CompanyID:          &currentCompany.ID,
			EntryGroup:         entryCounter.Group,
			Date:               input.Date,
			OutstandingBalance: input.PayableInput.Amount,
			TotalTransaction:   input.PayableInput.Amount,
			DueDate:            input.PayableInput.DueDate,
			Status:             payable.StatusPending,
		}).SaveX(ctx)
	}

	entries := append(input.Main, input.Counterpart...)
	for _, entry := range entries {
		if entry.Amount == 0 {
			continue
		}
		entryCounter.Number += 1

		var newEntry = tx.AccountingEntry.Create().
			SetInput(ent.CreateAccountingEntryInput{
				CompanyID:   &currentCompany.ID,
				UserID:      &currentUser.ID,
				Number:      entryCounter.Number,
				Group:       entryCounter.Group,
				Date:        &input.Date,
				Category:    &entry.Category,
				Account:     entry.Account,
				Label:       entry.Label,
				Amount:      entry.Amount,
				Description: *input.Description,
				AccountType: entry.AccountType,
				IsDebit:     entry.IsDebit,
				// Files:       files,
			})
		accountingEntries = append(accountingEntries, newEntry)
	}

	// Save changes:
	// 1. Accounting entries
	_, err = tx.AccountingEntry.CreateBulk(accountingEntries...).Save(ctx)
	if err != nil {
		fmt.Println("##", err)
		return nil, err
	}

	// 2. Update company: last entry date
	_, err = currentCompany.Update().SetLastEntryDate(input.Date).Save(ctx)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	// This the operation is part of the initial setup, remove the incomplete setup flag from the company
	if input.OperationType == model.BaseOperationTypeInitialSetup {
		_, err = tx.Company.UpdateOneID(currentCompany.ID).SetIncompleteSetup(false).Save(ctx)
		fmt.Println("* err:", err, input.OperationType, model.BaseOperationTypeInitialSetup)
		if err != nil {
			return nil, fmt.Errorf("an error occurred while processing the entry")
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	output := "new accounting entry was recorded successfully"
	return &output, nil
}