package ledger

import (
	"context"
	ent "mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/mazza/generated/model"
	"time"

	"entgo.io/ent/dialect/sql"
)

func BuildReport(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, startDate time.Time, endDate time.Time) (output *model.FileDetailsOutput, err error) {
	var ledger []ledgerStruct
	debitQuery := sql.Expr("CASE WHEN is_debit THEN amount ELSE 0 END AS debit")
	creditQuery := sql.Expr("CASE WHEN is_debit THEN 0 ELSE amount END AS credit")

	err = client.AccountingEntry.Query().
		Where(accountingentry.HasCompanyWith(company.IDEQ(currentCompany.ID))).
		Modify(func(s *sql.Selector) {
			s.Select(
				"id", "date", "account_type", "account", "label", "description", "group", "is_reversal", "number", "reversed",
			)
			s.AppendSelectExpr(debitQuery, creditQuery).
			OrderBy(sql.Desc(accountingentry.FieldNumber))
		}).Scan(ctx, &ledger)

	if err != nil {
		return nil, err
	}
	
	file, _, err := generatePDFDoc(&ledger, &currentCompany, startDate, endDate)

	output = &model.FileDetailsOutput{
		Message: "ledger",
		File: &model.FileOutput{
			Encoding: "base64",
			Kind:     "application/pdf",
			Name:     "diario_contabilistico.pdf",
			Data:     *file,
		},
	}

	return output, err
}
