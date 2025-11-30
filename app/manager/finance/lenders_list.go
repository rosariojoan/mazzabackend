package finance

import (
	"context"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/loan"
	"mazza/ent/generated/supplier"
)

// Get the list of all the lenders of a given company
func GetLendersList(ctx context.Context, client *generated.Client, companyID int) ([]*generated.Supplier, error) {
	// _, activeCompany := utils.GetSession(&ctx)
	lenders, err := client.Supplier.Query().Where(
		supplier.HasCompanyWith(company.ID(companyID)),
		supplier.HasLoansWith(loan.IsLending(false)),
	).All(ctx)
	if err != nil {
		return []*generated.Supplier{}, nil
	}

	return lenders, nil
}
