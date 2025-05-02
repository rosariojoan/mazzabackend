package suppliers

import (
	"context"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/payable"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func AggregatePayables(ctx context.Context, client *generated.Client, where *generated.PayableWhereInput, groupBy []model.PayablesGroupBy) ([]*model.PayableAggregationOutput, error) {
	_, currentCompany := utils.GetSession(&ctx)
	var result []*model.PayableAggregationOutput
	var agg []struct {
		Company int     `json:"company_customers"`
		Count   int     `json:"count"`
		Sum     float64 `json:"sum"`
	}

	query, err := where.Filter(client.Payable.Query())
	if err != nil {
		return nil, err
	}
	err = query.Where(payable.HasCompanyWith(company.IDEQ(currentCompany.ID))).
		// GroupBy().
		Aggregate(generated.Count(), generated.Sum(payable.FieldOutstandingBalance)).Scan(ctx, &agg)

	if err == nil {
		for _, item := range agg {
			result = append(result, &model.PayableAggregationOutput{
				Company: &item.Company,
				Count:   &item.Count,
				Sum:     &item.Sum,
			})
		}
	}

	return result, err
}
