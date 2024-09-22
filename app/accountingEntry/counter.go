package accountingentry

import (
	"context"
	ent "mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
)

type entryCounter struct {
	Group  int
	Number int
}

// Return counter:
// counter.Group: next entry group number
// counter.Number: last entry number
func GetEntryCounter(ctx context.Context, client *ent.Client, companyID int) entryCounter {
	var counter = entryCounter{}
	entry, err := client.AccountingEntry.Query().
		Where(accountingentry.HasCompanyWith(company.IDEQ(int(companyID)))).
		Order(ent.Desc(accountingentry.FieldID)).
		Select(accountingentry.FieldGroup, accountingentry.FieldNumber).
		First(ctx)

	if err != nil {
		counter.Group = 1
		counter.Number = 0
	} else {
		counter.Group = int(entry.Group) + 1
		counter.Number = int(entry.Number)
	}

	return counter
}
