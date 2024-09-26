package inventory

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"math"

	entryPkg "mazza/app/accountingEntry"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/generated/customer"
	"mazza/ent/generated/file"
	"mazza/ent/generated/product"
	"mazza/ent/generated/receivable"
	"mazza/ent/generated/treasury"
	u "mazza/ent/utils"
	"mazza/mazza/generated/model"
	"strconv"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

/*
Accounting entry: sales and sales return

# Sold quantities should be negative

Returned quantities should be positive
*/
func SalesRegistration(ctx context.Context, client *ent.Client, input model.SalesRegistrationInput) (*model.FileDetailsOutput, error) {
	utils.PP(input)
	if input.OperationType != model.SalesOperationTypeSales {
		return nil, gqlerror.Errorf("operation not supported")
	}

	country, lang := "mz", "pt"
	currentUser, currentCompany, _ := u.GetSession(&ctx)
	accountNames, err := utils.LoadAccountNames(country, lang)
	if err != nil {
		return nil, err
	}

	outputFile := struct {
		Name string
		File []byte
		URI  string
	}{}
	utils.PP(input)
	var retrievedTreasury = []*ent.Treasury{}
	// var treasuryMovements []*ent.CashMovementCreate

	cashIds := []int{}
	for _, p := range input.Counterpart {
		if p.TreasuryID != nil {
			cashIds = append(cashIds, *p.TreasuryID)
		}
	}

	// Check if the sale was paid partially or fully paid via cash (or bank transfer)
	if len(cashIds) > 0 {
		retrievedTreasury, err = client.Treasury.Query().
			Where(
				treasury.HasCompanyWith(company.IDEQ(currentCompany.ID)),
				treasury.IDIn(cashIds...),
			).All(ctx)
		if err != nil {
			return nil, err
		}
		if len(cashIds) != len(retrievedTreasury) {
			return nil, gqlerror.Errorf("invalid treasury account input")
		}

		// Update cash balance
		for _, account := range retrievedTreasury {
			for _, p := range input.Counterpart {
				if p.TreasuryID == &account.ID {
					account.Balance += p.Amount
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

	// Check if input products exist and have enough stock
	retrievedProducts, err := client.Product.Query().
		Where(product.HasCompanyWith(company.IDEQ(currentCompany.ID)), product.IDIn(productIDs...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if len(productIDs) != len(retrievedProducts) {
		return nil, gqlerror.Errorf("invalid product input")
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
	var COGS = []entry{}

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

	isCountable := func(cat product.Category) bool {
		return cat == product.CategoryMERCHANDISE
	}

	extraEntries := []*model.EntryItem{}
	// var productMovements []*ent.ProductMovementCreate
	for i, product := range retrievedProducts {
		if product.IsDefault || !isCountable(product.Category) {
			continue
		}
		for _, inp := range input.Main {
			if product.ID != *inp.ProductID {
				continue
			}

			// Update product movement
			// unitCost := product.UnitCost
			inputQuantity := 1
			if inp.Quantity != nil {
				inputQuantity = *inp.Quantity
			}
			price := math.Abs(math.Round(100*inp.Amount/float64(inputQuantity))) / 100

			// Update inventory: product price, stock, unit cost
			retrievedProducts[i].Price = int(price)
			retrievedProducts[i].Stock += -float64(inputQuantity)
			inventoryAccount := inventoryAccounts[product.Category.String()]
			costAccount := costOfSalesAccounts[product.Category.String()]
			cost := product.UnitCost * float64(inputQuantity)

			if input.OperationType == model.SalesOperationTypeSales {
				// Sale: reduce inventory and increase cost of sales
				for i := range COGS {
					account := COGS[i].Account
					if account == inventoryAccount {
						COGS[i].Amount -= cost
						inventoryAccount = ""
					} else if COGS[i].Account == costAccount {
						COGS[i].Amount += cost
						costAccount = ""
					}
				}

				if inventoryAccount != "" {
					extraEntries = append(extraEntries, &model.EntryItem{
						Account: inventoryAccount,
						Amount:  -cost, IsDebit: false,
						AccountType: accountingentry.AccountTypeASSET,
						Label:       "",
					})
				}
				if costAccount != "" {
					extraEntries = append(extraEntries, &model.EntryItem{
						Account: costAccount,
						Amount:  cost, IsDebit: true,
						AccountType: accountingentry.AccountTypeEXPENSE,
						Label:       "",
					})
				}
			} else {
				// Sales return: add back inventory and reduce cost of sales
				for i := range COGS {
					if COGS[i].Account == inventoryAccount {
						COGS[i].Amount += cost
						inventoryAccount = ""
					} else if COGS[i].Account == costAccount {
						COGS[i].Amount -= cost
						costAccount = ""
					}
				}

				if inventoryAccount != "" {
					quantity := 0
					extraEntries = append(extraEntries, &model.EntryItem{
						Account: inventoryAccount,
						Amount:  cost, IsDebit: false,
						AccountType: accountingentry.AccountTypeASSET,
						Label:       "",
						Quantity:    &quantity,
					})
				}
				if costAccount != "" {
					quantity := 0
					extraEntries = append(extraEntries, &model.EntryItem{
						Account: costAccount,
						Amount:  -cost, IsDebit: true,
						AccountType: accountingentry.AccountTypeEXPENSE,
						Label:       "",
						Quantity:    &quantity,
					})
				}
			}
		}
	}

	totalTransactionValue := input.TotalTransactionValue

	// Make sure that the customer exists
	_, err = client.Customer.Query().
		Where(customer.IDEQ(input.Customer.ID), customer.HasCompanyWith(company.IDEQ(currentCompany.ID))).
		First(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("invalid customer input")
	}

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
		return nil, err
	}

	// 2. Accounts receivable
	if input.Customer.Amount < totalTransactionValue {
		accountsReceivable := client.Receivable.Create().
			SetInput(ent.CreateReceivableInput{
				CustomerID:         &input.Customer.ID,
				EntryGroup:         entryCounter.Group,
				Date:               input.Date,
				OutstandingBalance: input.Customer.Amount,
				TotalTransaction:   totalTransactionValue,
				DaysDue:            input.Customer.DaysDue,
				Status:             receivable.StatusUnpaid,
			})

		if _, err := accountsReceivable.Save(ctx); err != nil {
			return nil, err
		}
	}

	// 3. Inventory stock
	for _, product := range retrievedProducts {
		_, err := product.Update().Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	if len(retrievedTreasury) > 0 {
		// 4. Cash and cash equivalent (treasury) balances
		for _, treasury := range retrievedTreasury {
			_, err := treasury.Update().Save(ctx)
			if err != nil {
				return nil, err
			}
		}
	}

	// 5. If invoice has been requested, issue and save it
	if input.IssueInvoice {
		fileByte, size, err := GenerateInvoicePDF(input.Invoice)
		if err != nil {
			log.Default().Printf("error: %s", err)
			return nil, gqlerror.Errorf("entry error. invoice could not be issued")
		}

		filename := "fatura_" + input.Invoice.Number + ".pdf"
		fileProps := utils.FileProps{
			Bucket:      "invoice",
			CompanyID:   currentCompany.ID,
			ContentType: utils.ContentType.PDF,
			Content:     fileByte,
			Name:        filename,
		}

		uri, url, err := utils.UploadFile(fileProps)
		if err != nil {
			return nil, gqlerror.Errorf("entry error. invoice could not be issued")
		}

		// Generate file description
		date, err := input.Date.MarshalText()
		if err != nil {
			date = []byte(input.Date.String())
		}

		description, err := json.Marshal(map[string]any{
			"invoiceNumber": input.Invoice.Number,
			"clientName":    input.Invoice.CustomerDetails.Name,
			"date":          string(date), // this is deliberatelly used instead of input.Invoice.Date
			"total":         input.Invoice.Totals.Total,
		})
		if err != nil {
			return nil, err
		}

		_, err = client.File.Create().SetInput(ent.CreateFileInput{
			Category:    file.CategoryINVOICE,
			Extension:   "pdf",
			Size:        strconv.Itoa(int(size)),
			URI:         *uri,
			URL:         *url,
			Description: string(description),
		}).Save(ctx)
		if err != nil {
			_ = utils.DeleteFile(*uri)
			return nil, err
		}

		outputFile.File = fileByte
		outputFile.Name = filename
		outputFile.URI = *uri
		currentCompany.LastInvoiceNumber += 1
	}

	// 8. Update company info
	companyUpdate := currentCompany.Update().SetLastEntryDate(input.Date)
	_, err = companyUpdate.Save(ctx)
	if err != nil {
		_ = utils.DeleteFile(outputFile.URI)
		return nil, err
	}

	output := model.FileDetailsOutput{
		Message: "new accounting entry was recorded successfully",
	}
	if input.IssueInvoice {
		output.File = &model.FileOutput{
			Encoding: "base64",
			Kind:     "application/pdf",
			Name:     outputFile.Name,
			Data:     base64.StdEncoding.EncodeToString(outputFile.File),
		}
	}

	return &output, nil
}
