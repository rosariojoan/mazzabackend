package finance

import (
	"context"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/customer"
	"mazza/ent/generated/loan"
)

// Get the list of all the debtors of a given company
func GetDebtorsList(ctx context.Context, client *generated.Client, companyID int) ([]*generated.Customer, error) {
	// _, activeCompany := utils.GetSession(&ctx)
	debtors, err := client.Customer.Query().Where(
		customer.HasCompanyWith(company.ID(companyID)),
		customer.HasLoansWith(loan.IsLending(true)),
	).All(ctx)
	if err != nil {
		return []*generated.Customer{}, nil
	}

	return debtors, nil
}
