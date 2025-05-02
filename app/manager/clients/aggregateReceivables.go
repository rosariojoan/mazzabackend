package clients

import (
	"context"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/receivable"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

func AggregateReceivables(ctx context.Context, client *generated.Client, where *generated.ReceivableWhereInput, groupBy []model.ReceivablesGroupBy) ([]*model.ReceivableAggregationOutput, error) {
	_, currentCompany := utils.GetSession(&ctx)
	var result []*model.ReceivableAggregationOutput
	var agg []struct {
		Company int     `json:"company_customers"`
		Count   int     `json:"count"`
		Sum     float64 `json:"sum"`
	}

	query, err := where.Filter(client.Receivable.Query())
	if err != nil {
		return nil, err
	}
	err = query.Where(receivable.HasCompanyWith(company.IDEQ(currentCompany.ID))).
		// GroupBy().
		Aggregate(generated.Count(), generated.Sum(receivable.FieldOutstandingBalance)).Scan(ctx, &agg)

	if err == nil {
		for _, item := range agg {
			result = append(result, &model.ReceivableAggregationOutput{
				Company: &item.Company,
				Count:   &item.Count,
				Sum:     &item.Sum,
			})
		}
	}

	return result, err
}
