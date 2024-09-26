package inventory

import (
	"context"
	"fmt"

	entryPkg "mazza/app/accountingEntry"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/payable"
	"mazza/ent/generated/product"
	"mazza/ent/generated/treasury"
	u "mazza/ent/utils"
	"mazza/mazza/generated/model"
)

// This function handles the registration of purchase and return of inventory
func PurchaseRegistration(ctx context.Context, client *ent.Client, input model.PurchaseRegistrationInput) (string, error) {
	if input.OperationType != model.PurchaseOperationTypePurchase {
		return "", fmt.Errorf("operation not supported")
	}

	country, lang := "mz", "pt"
	currentUser, currentCompany, _ := u.GetSession(&ctx)
	accountNames, err := utils.LoadAccountNames(country, lang)
	if err != nil {
		return "", err
	}

	// Check if the purchase was paid partially or fully paid via cash (or bank transfer)
	var retrievedTreasury = []*ent.Treasury{}

	cashIds := []int{}
	for _, p := range input.Counterpart {
		if p.TreasuryID != nil {
			cashIds = append(cashIds, *p.TreasuryID)
		}
	}

	if len(cashIds) > 0 {
		retrievedTreasury, err = client.Treasury.Query().
			Where(
				treasury.HasCompanyWith(company.IDEQ(currentCompany.ID)),
				treasury.IDIn(cashIds...),
			).All(ctx)
		if err != nil {
			return "", err
		}
		if len(cashIds) != len(retrievedTreasury) {
			return "", fmt.Errorf("invalid treasury account input")
		}
		// utils.PP(retrievedTreasury)
		// Check if the cash input amount is less than the balance.
		// Update cash balance
		for _, account := range retrievedTreasury {
			for _, p := range input.Counterpart {
				if p.TreasuryID != nil && *p.TreasuryID == account.ID {
					if p.Amount > account.Balance {
						return "", fmt.Errorf("not enoughc cash balance")
					}
					account.Balance += -p.Amount
					continue
				}
			}
		}
	}

	productIDs := []int{}
	for _, p := range input.Main {
		if p.ProductID != nil {
			productIDs = append(productIDs, *p.ProductID)
		}
	}

	// Check if purchased products are registered in the inventory
	// The purchase product must be a tangible or generic product, not a service
	retrievedProducts, err := client.Product.Query().
		Where(product.HasCompanyWith(company.IDEQ(currentCompany.ID)), product.IDIn(productIDs...)).
		All(ctx)
	if err != nil {
		return "", err
	}
	if len(productIDs) != len(retrievedProducts) {
		return "", fmt.Errorf("invalid product input")
	}

	// Update stock
	for _, prod := range retrievedProducts {
		if prod.Category != product.CategoryMERCHANDISE {
			continue
		}
		for _, p := range input.Main {
			inputQuantity := 0
			if p.Quantity != nil {
				inputQuantity = *p.Quantity
			}
			if p.ProductID == &prod.ID {
				prod.UnitCost += (prod.Stock*prod.UnitCost + p.Amount) / (prod.Stock + float64(inputQuantity))
				prod.Stock += float64(inputQuantity)
				continue
			}
		}
	}

	var accountingEntries []*ent.AccountingEntryCreate
	var entryCounter = entryPkg.GetEntryCounter(ctx, client, currentCompany.ID)

	// Prepare accounts payable
	accountsPayable := client.Payable.Create()
	if input.OperationType == model.PurchaseOperationTypePurchase && input.Supplier.Amount > 0 {
		accountsPayable.SetInput(ent.CreatePayableInput{
			SupplierID:         &input.Supplier.ID,
			EntryGroup:         entryCounter.Group,
			Date:               input.Date,
			OutstandingBalance: input.Supplier.Amount,
			TotalTransaction:   input.TotalTransactionValue,
			DaysDue:            input.Supplier.DaysDue,
			Status:             payable.StatusUNPAID,
		})
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
				Label:       accountNames[entry.Account],
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
		return "", err
	}

	// 2. Accounts payable
	if input.Supplier.Amount != 0 {
		_, err := accountsPayable.Save(ctx)
		if err != nil {
			return "", err
		}
	}

	// 3. Save product updates
	for _, product := range retrievedProducts {
		_, err := product.Update().SetStock(product.Stock).SetUnitCost(product.UnitCost).Save(ctx)
		if err != nil {
			return "", err
		}
	}

	if len(retrievedTreasury) > 0 {
		// 4. Cash and cash equivalent (treasury) balances
		for _, treasury := range retrievedTreasury {
			_, err := treasury.Update().SetBalance(treasury.Balance).Save(ctx)
			if err != nil {
				return "", err
			}
		}
	}

	// 5. Company: last entry date, last invoice number
	_, err = currentCompany.Update().SetLastEntryDate(input.Date).Save(ctx)
	if err != nil {
		return "", err
	}

	return "new accounting entry was recorded successfully", nil
}
