package mazza

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.52

import (
	"context"
	"fmt"
	"mazza/ent/generated"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/generated/companydocument"
	"mazza/ent/generated/membersignuptoken"
	"mazza/ent/generated/payable"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projecttask"
	"mazza/ent/generated/receivable"
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
func (r *queryResolver) AccountingEntries(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.AccountingEntryOrder, where *generated.AccountingEntryWhereInput) (*generated.AccountingEntryConnection, error) {
	_, currentCompany := utils.GetSession(&ctx)
	companyFilter := accountingentry.HasCompanyWith(company.IDEQ(currentCompany.ID))
	return r.client.AccountingEntry.Query().Where(companyFilter).Paginate(
		ctx, after, first, before, last,
		generated.WithAccountingEntryFilter(where.Filter),
		generated.WithAccountingEntryOrder(orderBy),
	)
}

// CompanyDocuments is the resolver for the companyDocuments field.
func (r *queryResolver) CompanyDocuments(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.CompanyDocumentOrder, where *generated.CompanyDocumentWhereInput) (*generated.CompanyDocumentConnection, error) {
	_, currentCompany := utils.GetSession(&ctx)
	companyFilter := companydocument.HasCompanyWith(company.IDEQ(currentCompany.ID))

	// result, err := inits.Firestore.Collection("users").Doc("XddvScdAef").Set(ctx, map[string]any{
	// 	"active": false,
	// 	"role":   "accounting",
	// })
	// fmt.Println("err:", err, "\nresult:", result)
	return r.client.CompanyDocument.Query().Where(companyFilter).Paginate(
		ctx, after, first, before, last,
		generated.WithCompanyDocumentFilter(where.Filter),
		generated.WithCompanyDocumentOrder(orderBy),
	)
}

// Files is the resolver for the files field.
func (r *queryResolver) Files(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *generated.FileOrder, where *generated.FileWhereInput) (*generated.FileConnection, error) {
	currentUser, currentCompany := utils.GetSession(&ctx)
	userQ := user.IDEQ(currentUser.ID)
	companyQ := company.IDEQ(currentCompany.ID)

	return r.client.User.Query().Where(userQ).QueryCompany().Where(companyQ).QueryFiles().Paginate(
		ctx, after, first, before, last,
		generated.WithFileOrder(orderBy),
		generated.WithFileFilter(where.Filter),
	)
}

// MemberSignupTokens is the resolver for the memberSignupTokens field.
func (r *queryResolver) MemberSignupTokens(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.MemberSignupTokenOrder, where *generated.MemberSignupTokenWhereInput) (*generated.MemberSignupTokenConnection, error) {
	_, currentCompany := utils.GetSession(&ctx)
	companyFilter := company.IDEQ(currentCompany.ID)

	return r.client.MemberSignupToken.Query().
		Where(membersignuptoken.HasCompanyWith(companyFilter)).
		Paginate(
			ctx, after, first, before, last,
			generated.WithMemberSignupTokenFilter(where.Filter),
			generated.WithMemberSignupTokenOrder(orderBy),
		)
}

// Payables is the resolver for the payables field.
func (r *queryResolver) Payables(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.PayableOrder, where *generated.PayableWhereInput) (*generated.PayableConnection, error) {
	_, currentCompany := utils.GetSession(&ctx)

	return r.client.Payable.Query().
		Where(payable.HasCompanyWith(company.IDEQ(currentCompany.ID))).
		Paginate(
			ctx, after, first, before, last,
			generated.WithPayableFilter(where.Filter),
		)
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.ProjectOrder, where *generated.ProjectWhereInput) (*generated.ProjectConnection, error) {
	_, currentCompany := utils.GetSession(&ctx)
	return r.client.Project.Query().
		Where(project.HasCompanyWith(company.IDEQ(currentCompany.ID))).
		Paginate(
			ctx, after, first, before, last,
			generated.WithProjectFilter(where.Filter),
			generated.WithProjectOrder(orderBy),
		)
}

// ProjectTasks is the resolver for the projectTasks field.
func (r *queryResolver) ProjectTasks(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.ProjectTaskOrder, where *generated.ProjectTaskWhereInput) (*generated.ProjectTaskConnection, error) {
	_, currentCompany := utils.GetSession(&ctx)
	projectFilter := project.HasCompanyWith(company.IDEQ(currentCompany.ID))

	return r.client.ProjectTask.Query().
		Where(projecttask.HasProjectWith(projectFilter)).
		Paginate(
			ctx, after, first, before, last,
			generated.WithProjectTaskFilter(where.Filter),
			generated.WithProjectTaskOrder(orderBy),
		)
}

// Receivables is the resolver for the receivables field.
func (r *queryResolver) Receivables(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*generated.ReceivableOrder, where *generated.ReceivableWhereInput) (*generated.ReceivableConnection, error) {
	_, currentCompany := utils.GetSession(&ctx)

	return r.client.Receivable.Query().
		Where(receivable.HasCompanyWith(company.IDEQ(currentCompany.ID))).
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
