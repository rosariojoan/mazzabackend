package finance

import (
	"context"
	"fmt"
	accountingentry "mazza/app/accountingEntry"
	"mazza/ent/generated"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func CreateLoan(ctx context.Context, client *generated.Client, input model.CreateLoanInputData, schedule []*generated.CreateLoanScheduleInput) (*generated.Loan, error) {
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
	go CreateLoanSchedule(ctx, client, schedule, newLoan.ID)

	return newLoan, nil
}
