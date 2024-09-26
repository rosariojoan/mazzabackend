package ledger

import (
	ae "mazza/ent/generated/accountingentry"
)

func isDebitable(typeStr string) bool {
	return typeStr == ae.AccountTypeASSET.String() || typeStr == ae.AccountTypeEXPENSE.String() || typeStr == ae.AccountTypeCONTRA_ASSET.String()
	// return typeStr == AccountType.Asset || typeStr == AccountType.Expense || typeStr == AccountType.ContraAsset
}
