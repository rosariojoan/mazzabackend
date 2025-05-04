package trialbalance

import (
	"context"
	"fmt"
	"mazza/app/utils"
	"mazza/mazza/generated/model"
	"strings"
	"time"

	ent "mazza/ent/generated"
)

func BuildReport(client *ent.Client, ctx context.Context, user ent.User, currentCompany ent.Company, endDate time.Time) (output *model.FileDetailsOutput, err error) {
	result, err := GetTrialBalance(client, ctx, user, currentCompany, endDate, []string{})
	if err != nil {
		return nil, err
	}

	startDate := utils.StartOfYear(endDate)
	file, _, err := generatePDFDoc(result, &currentCompany, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	output = &model.FileDetailsOutput{
		Message: "balancete",
		File: &model.FileOutput{
			Encoding: "base64",
			Kind:     "application/pdf",
			Name:     "balancete.pdf",
			Data:     *file,
		},
	}

	return output, err
}

func GetTrialBalance(
	client *ent.Client,
	ctx context.Context,
	user ent.User,
	currentCompany ent.Company,
	endDate time.Time,
	excludeAccountTypes []string,
) (output []*model.TrialBalanceRowItem, err error) {

	excluded := `'` + strings.Join(excludeAccountTypes, `','`) + `'` // e.g. excluded = ('a','b')
	country, lang := "MZ", "pt"

	total := struct {
		Debit  float64 `json:"debit"`
		Credit float64 `json:"credit"`
	}{Debit: 0, Credit: 0}

	// sqlStr := fmt.Sprintf(`
	// 	SELECT
	// 		account,
	// 		sum(debit) AS debit,
	// 		sum(credit) AS credit,
	// 		CASE
	// 			WHEN account_type IN ('asset', 'expense')
	// 			THEN sum(debit) - sum(credit)
	// 			ELSE sum(credit) - sum(debit)
	// 		END AS balance
	// 	FROM
	// 	(
	// 		SELECT
	// 			account,
	// 			account_type,
	// 			CASE WHEN debit < 0 THEN -debit ELSE debit END AS debit,
	// 			CASE WHEN credit < 0 THEN -credit ELSE credit END AS credit
	// 		FROM
	// 		(
	// 			SELECT
	// 				account AS account,
	// 				account_type,
	// 				CASE WHEN is_debit THEN amount ELSE 0 END AS debit,
	// 				CASE WHEN is_debit THEN 0 ELSE amount END AS credit
	// 			FROM accounting_entries
	// 			WHERE company_accounting_entries = %d AND date < '%s' AND account_type NOT IN (%s)
	// 		) AS intermediate_result
	// 	) AS result

	// 	GROUP BY account, account_type
	// 	ORDER BY account ASC
	// `, currentCompany.ID, endDate.Format(time.RFC3339), excluded)

	sqlStr := fmt.Sprintf(`
		SELECT
			account,
			account_type,
			sum(debit) AS debit,
			sum(credit) AS credit,
			CASE 
				WHEN account_type IN ('ASSET', 'EXPENSE') 
				THEN sum(debit) - sum(credit)
				ELSE sum(credit) - sum(debit)
			END AS balance
		FROM (
			SELECT 
				account,
				account_type,
				CASE WHEN is_debit THEN amount ELSE 0 END AS debit,
				CASE WHEN is_debit THEN 0 ELSE amount END AS credit
			FROM accounting_entries
			WHERE company_accounting_entries = %d AND date < '%s' AND account_type NOT IN (%s)
		) AS result
		
		GROUP BY account, account_type
		ORDER BY account ASC
	`, currentCompany.ID, endDate.Format(time.RFC3339), excluded)

	fmt.Println("TB excluded:", excluded)
	rows, err := client.QueryContext(ctx, sqlStr)
	if err != nil {
		fmt.Println("GetTrialBalance err:", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var item model.TrialBalanceRowItem
		if err := rows.Scan(&item.Account, &item.AccountType, &item.Debit, &item.Credit, &item.Balance); err != nil {
			// Check for a scan error. Query rows will be closed with defer.
			fmt.Println("GetTrialBalance err:", err)
			return nil, err
		}
		output = append(output, &item)
	}

	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	err = rows.Close()
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	accountNames, err := utils.LoadAccountNames(country, lang)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}
	for i, item := range output {
		output[i].Label = accountNames[item.Account]
		total.Debit += item.Debit
		total.Credit += item.Credit
	}

	// if len(excludeAccountTypes) > 0 {
	// 	return &result, nil
	// }
	return output, nil
}
