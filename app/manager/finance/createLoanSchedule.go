package finance

import (
	"context"
	"fmt"
	"mazza/ent/generated"
)

func CreateLoanSchedule(ctx context.Context, client *generated.Client, input []*generated.CreateLoanScheduleInput, loanID int) ([]*generated.LoanSchedule, error) {
	// _, activeCompany := utils.GetSession(&ctx)

	if len(input) == 0 {
		return nil, fmt.Errorf("no data provided")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	var builder []*generated.LoanScheduleCreate

	for _, data := range input {
		builder = append(builder, tx.LoanSchedule.Create().
			SetInput(*data).
			SetLoanID(loanID),
		)
	}

	newSchedule, err := tx.LoanSchedule.CreateBulk(builder...).Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	return newSchedule, nil
}
