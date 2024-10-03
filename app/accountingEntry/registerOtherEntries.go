package accountingentry

import (
	"context"
	"fmt"
	ent "mazza/ent/generated"
	"mazza/ent/generated/payable"
	"mazza/ent/generated/receivable"
	u "mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func RegisterOtherEntries(ctx context.Context, client *ent.Client, input model.BaseEntryRegistrationInput) (*string, error) {
	currentUser, currentCompany, _ := u.GetSession(&ctx)
	// country, lang := "mz", "pt"

	var retrievedTreasury = []*ent.Treasury{}
	err := CheckAndUpdateCashInput(ctx, client, currentCompany.ID, append(input.Main, input.Counterpart...), retrievedTreasury, false)
	if err != nil {
		return nil, err
	} else if len(retrievedTreasury) > 0 {
		// Update cash and cash equivalent (treasury) balances
		for _, treasury := range retrievedTreasury {
			_, err := treasury.Update().SetBalance(treasury.Balance).Save(ctx)
			if err != nil {
				return nil, err
			}
		}
	}

	var retrievedProducts = []*ent.Product{}
	err = CheckAndUpdateProductInput(ctx, client, currentCompany.ID, append(input.Main, input.Counterpart...), retrievedProducts, false)
	if err != nil {
		return nil, err
	} else {
		// Save product updates
		for _, product := range retrievedProducts {
			_, err := product.Update().SetStock(product.Stock).SetUnitCost(product.UnitCost).Save(ctx)
			if err != nil {
				return nil, err
			}
		}
	}

	var accountingEntries []*ent.AccountingEntryCreate
	var entryCounter = GetEntryCounter(ctx, client, currentCompany.ID)

	if input.Customer != nil && input.Customer.Amount > 0 {
		// Increase in account receivable: create accounts receivable
		accountsReceivable := client.Receivable.Create()
		accountsReceivable.SetInput(ent.CreateReceivableInput{
			CustomerID:         &input.Customer.ID,
			EntryGroup:         entryCounter.Group,
			Date:               input.Date,
			OutstandingBalance: input.Customer.Amount,
			TotalTransaction:   input.TotalTransactionValue,
			DaysDue:            input.Customer.DaysDue,
			Status:             receivable.StatusUnpaid,
		})
		if _, err = accountsReceivable.Save(ctx); err != nil {
			return nil, fmt.Errorf("invalid customer input")
		}
	} else if input.Customer != nil && input.Customer.Amount < 0 {
		// TODO: Decrease in account receivable: check if there is an outstanding balance of this customer and update it
	}

	if input.Supplier != nil && input.Supplier.Amount > 0 {
		// Increase in account payable: create accounts payable
		accountsPayable := client.Payable.Create()
		accountsPayable.SetInput(ent.CreatePayableInput{
			SupplierID:         &input.Supplier.ID,
			EntryGroup:         entryCounter.Group,
			Date:               input.Date,
			OutstandingBalance: input.Supplier.Amount,
			TotalTransaction:   input.TotalTransactionValue,
			DaysDue:            input.Supplier.DaysDue,
			Status:             payable.StatusUNPAID,
		})
		if _, err = accountsPayable.Save(ctx); err != nil {
			return nil, fmt.Errorf("invalid supplier input")
		}
	} else if input.Supplier != nil && input.Supplier.Amount < 0 {
		// TODO: Decrease in account payable: check if there is an outstanding balance of this supplier and update it
	}

	entries := append(input.Main, input.Counterpart...)
	for _, entry := range entries {
		if entry.Amount == 0 {
			continue
		}
		entryCounter.Number += 1

		var newEntry = client.AccountingEntry.Create().
			SetInput(ent.CreateAccountingEntryInput{
				CompanyID:   &currentCompany.ID,
				UserID:      &currentUser.ID,
				Number:      entryCounter.Number,
				Group:       entryCounter.Group,
				Date:        &input.Date,
				Account:     entry.Account,
				Label:       entry.Label,
				Amount:      entry.Amount,
				Description: *input.Description,
				AccountType: entry.AccountType,
				IsDebit:     entry.IsDebit,
				Quantity:    entry.Quantity,
				ProductID:   entry.ProductID,
				TreasuryID:  entry.TreasuryID,
				// Files:       files,
			})

		accountingEntries = append(accountingEntries, newEntry)
	}

	// Save changes:
	// 1. Accounting entries
	_, err = client.AccountingEntry.CreateBulk(accountingEntries...).Save(ctx)
	if err != nil {
		fmt.Println("##", err)
		return nil, err
	}

	// Update company: last entry date
	_, err = currentCompany.Update().SetLastEntryDate(input.Date).Save(ctx)
	if err != nil {
		return nil, err
	}

	// This the operation is part of the initial setup, remove the incomplete setup flag from the company
	if input.OperationType != model.BaseOperationTypeInitialSetup {
		_, err = client.Company.UpdateOneID(currentCompany.ID).SetIncompleteSetup(false).Save(ctx)
		fmt.Println("err:", err)
		return nil, fmt.Errorf("an error occurred while processing the entry")
	}

	output := "registration successful"
	return &output, nil
}
