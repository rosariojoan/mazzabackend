package finance

import (
	"context"
	"fmt"
	accountingentry "mazza/app/accountingEntry"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/customer"
	"mazza/ent/generated/supplier"
	"mazza/ent/utils"
	"mazza/inits"
	"mazza/mazza/generated/model"
)

func CreateLoan(ctx context.Context, client *generated.Client, input model.CreateLoanInputData) (*generated.Loan, error) {
	_, activeCompany := utils.GetSession(&ctx)
	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	newLoan, err := tx.Loan.Create().
		SetInput(*input.LoanInput).
		SetCompanyID(activeCompany.ID).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	if *input.LoanInput.IsLending && input.LoanCounterpartyID != nil {
		exists, err := inits.Client.Customer.Query().
			Where(
				customer.ID(*input.LoanCounterpartyID),
				customer.HasCompanyWith(company.ID(activeCompany.ID)),
			).Exist(ctx)
		if err != nil || !exists {
			_, err = tx.Customer.Create().
				SetInput(*input.Borrower).
				SetCompanyID(activeCompany.ID).
				AddLoanIDs(newLoan.ID).
				Save(ctx)
			if err != nil {
				_ = tx.Rollback()
				fmt.Println(err)
				return nil, fmt.Errorf("an error occurred")
			}
		} else {
			_, err = tx.Customer.UpdateOneID(*input.LoanCounterpartyID).
				Where(customer.HasCompanyWith(company.ID(activeCompany.ID))).
				AddLoanIDs(newLoan.ID).
				Save(ctx)
			if err != nil {
				_ = tx.Rollback()
				fmt.Println(err)
				return nil, fmt.Errorf("an error occurred")
			}
		}
	} else if *input.LoanInput.IsLending {
		// Check if this client already exists
		var taxId string
		if input.Borrower.TaxID != nil {
			taxId = *input.Borrower.TaxID
		} else {
			taxId = ""
		}
		clientID, err := tx.Customer.Query().Where(
			customer.Name(input.Borrower.Name),
			customer.HasCompanyWith(company.ID(activeCompany.ID)),
			customer.TaxID(taxId),
		).FirstID(ctx)
		if err != nil {
			_, err = tx.Customer.Create().
				SetInput(*input.Borrower).
				SetCompanyID(activeCompany.ID).
				AddLoanIDs(newLoan.ID).
				Save(ctx)
			if err != nil {
				_ = tx.Rollback()
				return nil, fmt.Errorf("customer already")
			}
		} else {
			input.LoanCounterpartyID = &clientID
			_, err = tx.Customer.UpdateOneID(clientID).
				AddLoanIDs(newLoan.ID).
				Save(ctx)
			if err != nil {
				_ = tx.Rollback()
				return nil, fmt.Errorf("customer already")
			}
		}
	} else if input.LoanCounterpartyID != nil {
		// Borrowing
		exists, err := inits.Client.Supplier.Query().
			Where(
				supplier.ID(*input.LoanCounterpartyID),
				supplier.HasCompanyWith(company.ID(activeCompany.ID)),
			).Exist(ctx)
		if err != nil || !exists {
			_, err = tx.Supplier.Create().
				SetInput(*input.Lender).
				SetCompanyID(activeCompany.ID).
				AddLoanIDs(newLoan.ID).
				Save(ctx)
			if err != nil {
				_ = tx.Rollback()
				fmt.Println(err)
				return nil, fmt.Errorf("an error occurred")
			}
		} else {
			_, err = tx.Supplier.UpdateOneID(*input.LoanCounterpartyID).
				Where(supplier.HasCompanyWith(company.ID(activeCompany.ID))).
				AddLoanIDs(newLoan.ID).
				Save(ctx)
			if err != nil {
				_ = tx.Rollback()
				fmt.Println(err)
				return nil, fmt.Errorf("an error occurred")
			}
		}
	} else {
		// Also borrowing
		_, err = tx.Supplier.Create().
			SetInput(*input.Lender).
			SetCompanyID(activeCompany.ID).
			AddLoanIDs(newLoan.ID).
			Save(ctx)
		if err != nil {
			_ = tx.Rollback()
			fmt.Println(err)
			return nil, fmt.Errorf("an error occurred")
		}
	}

	if input.AccountingEntry != nil {
		input.AccountingEntry.LoanID = &newLoan.ID
		_, err = accountingentry.RegisterAccountingOperations(ctx, tx, *input.AccountingEntry, nil)
		if err != nil {
			_ = tx.Rollback()
			fmt.Println(err)
			return nil, fmt.Errorf("an error occurred")
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	// Create schedule
	_, _ = CreateLoanSchedule(ctx, client, input.Schedule, newLoan.ID)

	return newLoan, nil
}
