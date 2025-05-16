package finance

import (
	"context"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/loan"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func AggregateLoans(ctx context.Context, client *generated.Client, where *generated.LoanWhereInput, groupBy []model.LoansGroupBy) ([]*model.LoanAggregationOutput, error) {
	_, currentCompany := utils.GetSession(&ctx)
	var result []*model.LoanAggregationOutput
	var agg []struct {
		Company int     `json:"company_loans"`
		Count   int     `json:"count"`
		Sum     float64 `json:"sum"`
	}

	query, err := where.Filter(client.Loan.Query())
	if err != nil {
		return nil, err
	}
	err = query.Where(loan.HasCompanyWith(company.IDEQ(currentCompany.ID))).
		// GroupBy().
		Aggregate(generated.Count(), generated.Sum(loan.FieldOutstandingBalance)).Scan(ctx, &agg)

	if err == nil {
		for _, item := range agg {
			result = append(result, &model.LoanAggregationOutput{
				Company: &item.Company,
				Count:   &item.Count,
				Sum:     &item.Sum,
			})
		}
	}

	return result, err
}
