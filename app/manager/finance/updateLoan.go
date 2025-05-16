package finance

import (
	"context"
	"fmt"
	accountingentry "mazza/app/accountingEntry"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/loan"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func UpdateLoan(ctx context.Context, client *generated.Client, input model.UpdateLoanInputData) (*generated.Loan, error) {
	_, activeCompany := utils.GetSession(&ctx)
	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	updatedLoan, err := tx.Loan.UpdateOneID(input.ID).
		Where(loan.HasCompanyWith(company.ID(activeCompany.ID))).
		SetInput(*input.LoanInput).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	if input.AccountingEntry != nil {
		input.AccountingEntry.LoanID = &updatedLoan.ID
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
	return updatedLoan, nil
}
