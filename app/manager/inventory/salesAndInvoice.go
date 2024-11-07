package inventory

import (
	"context"
	"fmt"

	entryPkg "mazza/app/accountingEntry"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/treasury"
	u "mazza/ent/utils"
	"mazza/mazza/generated/model"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

/*
Accounting entry: sales and sales return

# Sold quantities should be negative

Returned quantities should be positive
*/
func SalesRegistration(ctx context.Context, client *ent.Client, input model.SalesRegistrationInput) (string, error) {
	utils.PP(input)
	companyQ := u.CurrentCompanyQuery(&ctx)
	if input.OperationType != model.SalesOperationTypeSales {
		return "", gqlerror.Errorf("operation not supported")
	}

	country, lang := "mz", "pt"
	currentUser, currentCompany, _ := u.GetSession(&ctx)
	accountNames, err := utils.LoadAccountNames(country, lang)
	if err != nil {
		return "", err
	}

	var retrievedTreasury = []*ent.Treasury{}
	// var treasuryMovements []*ent.CashMovementCreate
	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return "", fmt.Errorf("an error occurred")
	}

	if input.CashInput != nil {
		_ = tx.Treasury.Update().Where(treasury.HasCompanyWith(companyQ)).AddBalance(*input.CashInput).SaveX(ctx)
	}

	if input.ProductInput != nil {
		_ = tx.Product.Update().AddStock(*input.ProductInput).SaveX(ctx)
	}

	// lastEntry := models.AccountingEntry{}
	// inits.DB.Where("company_id = ?", company.ID).Last(&lastEntry)

	// Classify entry accounts
	type entry struct {
		Account    string
		Amount     float64
		TreasuryID *uint
		IsDebit    bool
		Type       accountingentry.AccountType
	}
	// var entries = []entry{}

	// // Debit and credit entries are separated to be sorted: first debit then credit
	// debitEntries := []entry{}
	// creditEntries := []entry{}

	// for _, item := range input.Main {
	// 	accountType := accountingentry.AccountTypeREVENUE
	// 	isDebit := false
	// 	if input.OperationType == model.SalesOperationTypeSales {
	// 		item.Amount = -item.Amount
	// 		isDebit = true
	// 	}

	// 	newEntry := entry{Account: item.Account, Amount: item.Amount, IsDebit: isDebit, Type: accountType}
	// 	if isDebit {
	// 		debitEntries = append(debitEntries, newEntry)
	// 	} else {
	// 		creditEntries = append(creditEntries, newEntry)
	// 	}
	// }

	// for _, item := range input.Counterpart {
	// 	accountType := accountingentry.AccountTypeASSET
	// 	isDebit := true
	// 	if input.OperationType == model.SalesOperationTypeSalesReturn {
	// 		isDebit = false
	// 		item.Amount = -item.Amount
	// 	}

	// 	newEntry := entry{Account: item.Account, Amount: item.Amount, IsDebit: isDebit, Type: accountType}
	// 	if isDebit {
	// 		debitEntries = append(debitEntries, newEntry)
	// 	} else {
	// 		creditEntries = append(creditEntries, newEntry)
	// 	}
	// }

	// // Sort entries: first debit then credit
	// entries = append(debitEntries, creditEntries...)

	extraEntries := []*model.EntryItem{}

	totalTransactionValue := input.TotalTransactionValue
	_ = totalTransactionValue

	// Prepare accounting entries
	var accountingEntries []*ent.AccountingEntryCreate
	var entryCounter = entryPkg.GetEntryCounter(ctx, client, currentCompany.ID)

	entries := append(input.Main, input.Counterpart...)
	entries = append(entries, extraEntries...)

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
				Label:       accountNames[entry.Account],
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
	_, err = client.AccountingEntry.CreateBulk(accountingEntries...).Save(ctx)
	if err != nil {
		return "", err
	}

	// // 2. Accounts receivable
	// if input.ReceivablesInput.Amount > 0 {
	// 	accountsReceivable := client.Receivable.Create().
	// 		SetInput(ent.CreateReceivableInput{
	// 			CustomerID:         &input.Customer.ID,
	// 			EntryGroup:         entryCounter.Group,
	// 			Date:               input.Date,
	// 			OutstandingBalance: input.Customer.Amount,
	// 			TotalTransaction:   totalTransactionValue,
	// 			DaysDue:            input.Customer.DaysDue,
	// 			Status:             receivable.StatusUnpaid,
	// 		})

	// 	if _, err := accountsReceivable.Save(ctx); err != nil {
	// 		return nil, err
	// 	}
	// }

	if len(retrievedTreasury) > 0 {
		// 4. Cash and cash equivalent (treasury) balances
		for _, treasury := range retrievedTreasury {
			_, err := treasury.Update().Save(ctx)
			if err != nil {
				return "", err
			}
		}
	}

	// 8. Update company info
	companyUpdate := currentCompany.Update().SetLastEntryDate(input.Date)
	_, _ = companyUpdate.Save(ctx)

	err = tx.Commit()
	if err != nil {
		fmt.Println("err:", err)
		return "", fmt.Errorf("an error occurred")
	}

	return "ok", nil
}
