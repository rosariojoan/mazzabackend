package finance

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/loan"
	"mazza/ent/generated/loanschedule"
	"mazza/ent/utils"
)

func UpdateLoanSchedule(ctx context.Context, client *generated.Client, loanID int, input []*generated.UpdateLoanScheduleInput) ([]*generated.LoanSchedule, error) {
	_, activeCompany := utils.GetSession(&ctx)

	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	for _, data := range input {
		_, err := tx.LoanSchedule.Update().
			Where(
				loanschedule.HasLoanWith(
					loan.ID(loanID),
					loan.HasCompanyWith(
						company.ID(activeCompany.ID),
					),
				),
			).SetInput(*data).Save(ctx)

		if err != nil {
			return nil, fmt.Errorf("an error occurred")
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		fmt.Println(err)
		return nil, fmt.Errorf("an error occurred")
	}

	updatedLoanSchedule, _ := client.LoanSchedule.Query().Where(
		loanschedule.HasLoanWith(loan.ID(loanID)),
	).All(ctx)

	return updatedLoanSchedule, nil
}
