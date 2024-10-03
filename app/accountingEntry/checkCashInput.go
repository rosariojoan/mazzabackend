package accountingentry

import (
	"context"
	"fmt"
	ent "mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/treasury"
	"mazza/mazza/generated/model"
)

// Check if cash input is valid. Prepare treasury balance update
func CheckAndUpdateCashInput(ctx context.Context, client *ent.Client, companyID int, input []*model.EntryItem, retrievedTreasury []*ent.Treasury, balanceRestriction bool) error {
	cashIds := []int{}
	for _, p := range input {
		if p.TreasuryID != nil {
			cashIds = append(cashIds, *p.TreasuryID)
		}
	}

	if len(cashIds) > 0 {
		retrievedTreasury, err := client.Treasury.Query().
			Where(
				treasury.HasCompanyWith(company.IDEQ(companyID)),
				treasury.IDIn(cashIds...),
			).All(ctx)
		if err != nil {
			return err
		}
		if len(cashIds) != len(retrievedTreasury) {
			return fmt.Errorf("invalid treasury account input")
		}

		// Check if the cash input amount is less than the balance.
		// Prepare cash balance update
		for _, account := range retrievedTreasury {
			for _, p := range input {
				if p.TreasuryID != nil && *p.TreasuryID == account.ID {
					if balanceRestriction && p.Amount > account.Balance {
						return fmt.Errorf("not enough cash balance")
					}

					if p.IsDebit {
						account.Balance += p.Amount
					} else {
						account.Balance += -p.Amount
					}
					continue
				}
			}
		}
	}

	return nil
}
