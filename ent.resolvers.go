package mazza

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/generated/user"
	"mazza/ent/utils"

	"entgo.io/contrib/entgql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (generated.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]generated.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// AccountingEntries is the resolver for the accountingEntries field.
func (r *queryResolver) AccountingEntries(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *generated.AccountingEntryOrder, where *generated.AccountingEntryWhereInput) (*generated.AccountingEntryConnection, error) {
	_, currentCompany, _ := utils.GetSession(&ctx)
	compFilter := accountingentry.HasCompanyWith(company.IDEQ(currentCompany.ID))
	return r.client.AccountingEntry.Query().Where(compFilter).Paginate(
		ctx, after, first, before, last,
		generated.WithAccountingEntryOrder(orderBy),
		generated.WithAccountingEntryFilter(where.Filter),
	)
}

// Files is the resolver for the files field.
func (r *queryResolver) Files(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *generated.FileOrder, where *generated.FileWhereInput) (*generated.FileConnection, error) {
	currentUser, currentCompany, _ := utils.GetSession(&ctx)
	userQ := user.IDEQ(currentUser.ID)
	companyQ := company.IDEQ(currentCompany.ID)

	return r.client.User.Query().Where(userQ).QueryCompany().Where(companyQ).QueryFiles().Paginate(
		ctx, after, first, before, last,
		generated.WithFileOrder(orderBy),
		generated.WithFileFilter(where.Filter),
	)
}

// Receivables is the resolver for the receivables field.
func (r *queryResolver) Receivables(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.ReceivableOrder, where *generated.ReceivableWhereInput) (*generated.ReceivableConnection, error) {
	currentUser, currentCompany, _ := utils.GetSession(&ctx)
	userQ := user.IDEQ(currentUser.ID)

	return r.client.User.Query().Where(userQ).
		QueryCompany().Where(company.IDEQ(currentCompany.ID)).
		QueryCustomers().QueryReceivables().
		Paginate(
			ctx, after, first, before, last,
			generated.WithReceivableFilter(where.Filter),
		)
}

// Tokens is the resolver for the tokens field.
func (r *queryResolver) Tokens(ctx context.Context) ([]*generated.Token, error) {
	panic(fmt.Errorf("not implemented: Tokens - tokens"))
}

// Workshifts is the resolver for the workshifts field.
func (r *queryResolver) Workshifts(ctx context.Context) ([]*generated.Workshift, error) {
	panic(fmt.Errorf("not implemented: Workshifts - workshifts"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
