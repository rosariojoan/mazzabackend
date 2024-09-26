package inventory

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	if input.OperationType != model.SalesOperationTypeSales {
		return nil, gqlerror.Errorf("operation not supported")
	}

	country, lang := "mz", "pt"
	currentUser, currentCompany, _ := u.GetSession(&ctx)
	accountNames, err := utils.LoadAccountNames(country, lang)
	if err != nil {
		return nil, err
	}

	var entryCounter = entryPkg.GetEntryCounter(ctx, client, currentCompany.ID)

	outputFile := struct {
		Name string
		File []byte
		URI  string
	}{}

	var retrievedTreasury = []*ent.Treasury{}
	var treasuryMovements []*ent.CashMovementCreate

	cashIds := []int{}
	for _, cashInput := range input.Cash {
		cashIds = append(cashIds, cashInput.ID)
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

		// Prepare treasury balances update
		for _, cashInput := range input.Cash {
			for i := range retrievedTreasury {
				if retrievedTreasury[i].ID == cashInput.ID {
					retrievedTreasury[i].Balance += cashInput.Amount

					createObj := client.CashMovement.Create().
						SetInput(ent.CreateCashMovementInput{
							TreasuryID: &cashInput.ID,
							Amount:     cashInput.Amount,
							Date:       input.Date,
							EntryGroup: entryCounter.Group,
						})

					treasuryMovements = append(treasuryMovements, createObj)
					break
				}
			}
		}
	}

	productIDs := []int{}
	for _, p := range input.Products {
		productIDs = append(productIDs, p.ID)
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
	var entries = []entry{}
	var COGS = []entry{}

	// Debit and credit entries are separated to be sorted: first debit then credit
	debitEntries := []entry{}
	creditEntries := []entry{}

	for _, item := range input.Main {
		accountType := accountingentry.AccountTypeREVENUE
		isDebit := false
		if input.OperationType == model.SalesOperationTypeSales {
			item.Amount = -item.Amount
			isDebit = true
		}

		newEntry := entry{Account: item.Account, Amount: item.Amount, IsDebit: isDebit, Type: accountType}
		if isDebit {
			debitEntries = append(debitEntries, newEntry)
		} else {
			creditEntries = append(creditEntries, newEntry)
		}
	}

	for _, item := range input.Counterpart {
		accountType := accountingentry.AccountTypeASSET
		isDebit := true
		if input.OperationType == model.SalesOperationTypeSalesReturn {
			isDebit = false
			item.Amount = -item.Amount
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

	isCountable := func(cat product.Category) bool {
		return cat == product.CategoryMERCHANDISE
	}

	var productMovements []*ent.ProductMovementCreate
	for i, product := range retrievedProducts {
		for _, inputProduct := range input.Products {
			if product.ID != inputProduct.ID {
				continue
			}

			// Update product movement
			unitCost := product.UnitCost
			price := math.Abs(math.Round(100*inputProduct.Amount/float64(inputProduct.Quantity))) / 100
			newProdMove := client.ProductMovement.Create().SetInput(
				ent.CreateProductMovementInput{
					ProductID:   &product.ID,
					EntryGroup:  entryCounter.Group,
					AverageCost: product.UnitCost,
					UnitCost:    unitCost,
					Price:       &price,
					Quantity:    inputProduct.Quantity,
				})
			productMovements = append(productMovements, newProdMove)

			if !isCountable(product.Category) {
				continue
			}

			// Update inventory: product price, stock, unit cost
			retrievedProducts[i].Price = int(price)
			retrievedProducts[i].Stock = product.Stock - float64(inputProduct.Quantity)
			fmt.Println("product.Category.String():", product.Category.String())
			inventoryAccount := inventoryAccounts[product.Category.String()]
			costAccount := costOfSalesAccounts[product.Category.String()]
			cost := product.UnitCost * float64(inputProduct.Quantity)

			// Reset debit and credit entries. This is done to classify COGS
			debitEntries = []entry{}
			creditEntries = []entry{}

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
					creditEntries = append(creditEntries, entry{Account: inventoryAccount, Amount: -cost, IsDebit: false, Type: accountingentry.AccountTypeASSET})
				}
				if costAccount != "" {
					debitEntries = append(debitEntries, entry{Account: costAccount, Amount: cost, IsDebit: true, Type: accountingentry.AccountTypeEXPENSE})
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
					creditEntries = append(creditEntries, entry{Account: inventoryAccount, Amount: cost, IsDebit: false, Type: accountingentry.AccountTypeASSET})
				}
				if costAccount != "" {
					debitEntries = append(debitEntries, entry{Account: costAccount, Amount: -cost, IsDebit: true, Type: accountingentry.AccountTypeEXPENSE})
				}
			}

			// Append classified COGS to the entries
			if len(debitEntries) > 0 {
				entries = append(entries, debitEntries...)
			}
			if len(creditEntries) > 0 {
				entries = append(entries, creditEntries...)
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
				AccountType: entry.Type,
				IsDebit:     entry.IsDebit,
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

	// 4. Product movements
	_, err = client.ProductMovement.CreateBulk(productMovements...).Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(input.Cash) > 0 {
		// 5. Create cash movements
		_, err := client.CashMovement.CreateBulk(treasuryMovements...).Save(ctx)
		if err != nil {
			return nil, err
		}

		// 6. Cash and cash equivalent (treasury) balances
		for _, treasury := range retrievedTreasury {
			_, err := treasury.Update().Save(ctx)
			if err != nil {
				return nil, err
			}
		}
	}

	// 7. If invoice has been requested, issue and save it
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
