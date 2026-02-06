package ledger

import (
	ae "mazza/ent/generated/accountingentry"
)

func isDebitable(typeStr string) bool {
	return typeStr == ae.AccountTypeAsset.String() || typeStr == ae.AccountTypeExpense.String() || typeStr == ae.AccountTypeContraAsset.String()
	// return typeStr == AccountType.Asset || typeStr == AccountType.Expense || typeStr == AccountType.ContraAsset
}
