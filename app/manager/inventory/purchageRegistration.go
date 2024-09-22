package inventory

import (
	"context"
	"fmt"

	entryPkg "mazza/app/accountingEntry"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/generated/payable"
	"mazza/ent/generated/product"
	"mazza/ent/generated/treasury"
	u "mazza/ent/utils"
	"mazza/mazza/generated/model"
	"strconv"
)

// This function handles the registration of purchase and return of inventory
func PurchaseRegistration(ctx context.Context, client *ent.Client, input model.PurchaseRegistrationInput) (string, error) {
	if input.OperationType != model.PurchaseOperationTypePurchase {
		return "", fmt.Errorf("operation not supported")
	}

	var assetAccounts = []string{"1", "2", "3"}
	// var liabilityAccounts = []string{"4"}
	var contraLiabilityAccounts = []string{}

	getAccountType := func(account *string) (accountType accountingentry.AccountType) {
		if utils.StartsWith(*account, assetAccounts) {
			return accountingentry.AccountTypeASSET
		} else if utils.StartsWith(*account, contraLiabilityAccounts) {
			return accountingentry.AccountTypeCONTRA_LIABILITY
		}
		return accountingentry.AccountTypeLIABILITY
	}

	country, lang := "mz", "pt"
	currentUser, currentCompany, _ := u.GetSession(&ctx)
	accountNames, err := utils.LoadAccountNames(&country, &lang)
	if err != nil {
		return "", err
	}

	var entryCounter = entryPkg.GetEntryCounter(ctx, client, currentCompany.ID)

	// Check if the purchase was paid partially or full via cash (or bank transfer)
	var retrievedTreasury = []*ent.Treasury{}
	var treasuryMovements []*ent.CashMovementCreate

	cashIds := []int{}
	for _, cashInput := range input.Cash {
		cashIds = append(cashIds, cashInput.ID)
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

		// Prepare treasury balances update
		for _, cashInput := range input.Cash {
			for i := range retrievedTreasury {
				if retrievedTreasury[i].ID == cashInput.ID {
					if retrievedTreasury[i].Balance < cashInput.Amount {
						name := retrievedTreasury[i].Name
						balance := strconv.Itoa(int(retrievedTreasury[i].Balance))
						purchaseAmount := strconv.Itoa(int(cashInput.Amount))
						return "", fmt.Errorf("not enough cash balance to register purchase by cash. Account: " + name + "  available balance: " + balance + "  purchase: " + purchaseAmount)
					}

					retrievedTreasury[i].Balance += cashInput.Amount
					var createObj ent.CashMovementCreate
					createObj.SetInput(ent.CreateCashMovementInput{
						TreasuryID: &cashInput.ID,
						Amount:     cashInput.Amount,
						Date:       input.Date,
						EntryGroup: entryCounter.Group,
					})

					treasuryMovements = append(treasuryMovements, &createObj)
					break
				}
			}
		}
	}

	productIDs := []int{}
	for _, p := range input.Products {
		productIDs = append(productIDs, p.ID)
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

	// Prepare product movement update
	var productMovements []*ent.ProductMovementCreate
	for i, product := range retrievedProducts {
		for _, inputProduct := range input.Products {
			if product.ID != inputProduct.ID {
				continue
			}

			newProdMove := client.ProductMovement.Create().SetProductID(product.ID).SetEntryGroup(entryCounter.Group)

			if product.IsDefault {
				retrievedProducts[i].UnitCost += inputProduct.Amount // Only update unit cost if it is a generic product
				newProdMove.SetAverageCost(inputProduct.Amount)
				newProdMove.SetUnitCost(inputProduct.Amount)
				newProdMove.SetPrice(0)
				newProdMove.SetQuantity(1)
			} else {
				currentTotalCost := retrievedProducts[i].UnitCost * float64(retrievedProducts[i].Stock)
				inputTotalCost := inputProduct.Amount
				newTotalStock := retrievedProducts[i].Stock + float64(inputProduct.Quantity)
				newUnitCost := (currentTotalCost + inputTotalCost) / float64(newTotalStock)

				retrievedProducts[i].Stock += float64(inputProduct.Quantity)
				retrievedProducts[i].UnitCost = newUnitCost
				newProdMove.SetAverageCost(newUnitCost)
				newProdMove.SetUnitCost(inputProduct.Amount / float64(inputProduct.Quantity))
				newProdMove.SetPrice(0)
				newProdMove.SetQuantity(inputProduct.Quantity)
			}

			productMovements = append(productMovements, newProdMove)
		}
	}

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

	// Classify entry accounts
	// Debit and credit entries are separated to be sorted: first debit then credit
	type entry struct {
		Account    string
		Amount     float64
		TreasuryID *uint
		IsDebit    bool
		Type       accountingentry.AccountType
	}
	entries := []entry{}
	debitEntries := []entry{}
	creditEntries := []entry{}

	for _, item := range input.Main {
		accountType := accountingentry.AccountTypeASSET
		isDebit := true
		if input.OperationType == model.PurchaseOperationTypePurchaseReturn {
			item.Amount = -item.Amount
			isDebit = false
		}

		newEntry := entry{Account: item.Account, Amount: item.Amount, IsDebit: isDebit, Type: accountType}
		if isDebit {
			debitEntries = append(debitEntries, newEntry)
		} else {
			creditEntries = append(creditEntries, newEntry)
		}
	}

	for _, item := range input.Counterpart {
		accountType := getAccountType(&item.Account)
		var isDebit bool
		if input.OperationType == model.PurchaseOperationTypePurchase {
			if accountType == accountingentry.AccountTypeASSET {
				item.Amount *= -1 // Reduce the asset used to pay for the purchase
			}
			isDebit = false
		} else {
			if accountType == accountingentry.AccountTypeLIABILITY || accountType == accountingentry.AccountTypeCONTRA_LIABILITY {
				item.Amount *= -1
			}
			isDebit = true
		}

		newEntry := entry{Account: item.Account, Amount: item.Amount, IsDebit: isDebit, Type: accountType}
		if isDebit {
			debitEntries = append(debitEntries, newEntry)
		} else {
			creditEntries = append(creditEntries, newEntry)
		}
	}

	// Sort entries: first debit then credit
	entries = append(debitEntries, creditEntries...)
	var accountingEntries []*ent.AccountingEntryCreate
	fmt.Println(" - ## ok")
	for _, entry := range entries {
		if entry.Amount == 0 {
			continue
		}
		entryCounter.Number += 1

		var newEntry = client.AccountingEntry.Create()
		fmt.Println("!! ok g:", entryCounter.Group, " n:", entryCounter.Number)
		newEntry.SetInput(ent.CreateAccountingEntryInput{
			CompanyID:   &currentCompany.ID,
			UserID:      &currentUser.ID,
			Number:      entryCounter.Number,
			Group:       entryCounter.Group,
			Date:        &input.Date,
			Account:     entry.Account,
			Label:       accountNames[entry.Account],
			Amount:      entry.Amount,
			Description: *input.Description,
			AccountType: accountingentry.AccountType(entry.Type),
			IsDebit:     entry.IsDebit,
			// Files:       files,
		})

		accountingEntries = append(accountingEntries, newEntry)
	}
	fmt.Println("*** ok")
	// Save changes:
	// 1. Accounting entries
	_, err = client.AccountingEntry.CreateBulk(accountingEntries...).Save(ctx)
	if err != nil {
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
		_, err := product.Update().Save(ctx)
		if err != nil {
			return "", err
		}
	}

	// 4. Product movements
	_, err = client.ProductMovement.CreateBulk(productMovements...).Save(ctx)
	if err != nil {
		return "", err
	}

	// 5. Inventory stock
	for _, product := range retrievedProducts {
		_, err := product.Update().Save(ctx)
		if err != nil {
			return "", err
		}
	}

	if len(input.Cash) > 0 {
		// 6. Cash and cash equivalent movements
		_, err := client.CashMovement.CreateBulk(treasuryMovements...).Save(ctx)
		if err != nil {
			return "", err
		}

		// 7. Cash and cash equivalent (treasury) balances
		for _, treasury := range retrievedTreasury {
			_, err := treasury.Update().Save(ctx)
			if err != nil {
				return "", err
			}
		}
	}

	// 8. Company: last entry date, last invoice number
	_, err = currentCompany.Update().SetLastEntryDate(input.Date).Save(ctx)
	if err != nil {
		return "", err
	}

	return "new accounting entry was recorded successfully", nil
}
