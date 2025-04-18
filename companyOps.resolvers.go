package mazza

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.52

import (
	"context"
	"fmt"
	accountingentry "mazza/app/accountingEntry"
	appUtils "mazza/app/utils"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/companydocument"
	"mazza/ent/generated/membersignuptoken"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/user"
	"mazza/ent/generated/userrole"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
	"time"
)

// InitialSetup is the resolver for the initialSetup field.
func (r *mutationResolver) InitialSetup(ctx context.Context, input model.InitialSetupInput) (*string, error) {
	if input.AccountingEntry != nil {
		tx, err := r.client.Tx(ctx)
		if err != nil {
			fmt.Println("err:", err)
			return nil, fmt.Errorf("an error occurred")
		}

		_, err = accountingentry.RegisterAccountingOperations(ctx, tx, *input.AccountingEntry, nil)
		if err != nil {
			fmt.Println("InitialSetup RegisterAccountingOperations err:", err)
			return nil, fmt.Errorf("an error occurred")
		}
		if err = tx.Commit(); err != nil {
			fmt.Println("err:", err)
			return nil, fmt.Errorf("an error occurred")
		}
	}

	user, company := utils.GetSession(&ctx)
	_ = user
	// Make sure that the user is an admin of this company
	// userIsAdmin := ...
	// if !userIsAdmin {
	// return nil, fmt.Error("you do not have privileges for this operation")
	// }
	// input.CompanyInfo.Industry
	_, err := r.client.Company.UpdateOneID(company.ID).
		SetVatRate(input.CompanyInfo.VatRate).
		SetIndustry(input.CompanyInfo.Industry).
		SetIncompleteSetup(false).
		Save(ctx)
	if err != nil {
		fmt.Println("InitialSetup company info update err:", err)
		return nil, fmt.Errorf("an error occurred")
	}

	result := "company info was updated"
	return &result, nil
}

// UploadDocument is the resolver for the uploadDocument field.
func (r *mutationResolver) UploadDocument(ctx context.Context, input generated.CreateCompanyDocumentInput) (*generated.CompanyDocument, error) {
	currentUser, currentCompany := utils.GetSession(&ctx)
	// roles, err := user.AssignedRoles(ctx)

	userIsAdmin, _ := r.client.UserRole.Query().Where(
		userrole.HasUserWith(user.ID(currentUser.ID)),
		userrole.HasCompanyWith(company.ID(currentCompany.ID)),
		userrole.RoleEQ(userrole.RoleADMIN),
	).Exist(ctx)

	builder := r.client.CompanyDocument.Create().
		SetInput(input).SetCompanyID(currentCompany.ID).SetUploadedByID(currentUser.ID)
	if userIsAdmin {
		builder = builder.SetStatus(companydocument.StatusAPPROVED).SetApprovedByID(currentUser.ID)
	} else {
		builder = builder.SetStatus(companydocument.StatusPENDING)
	}

	document, err := builder.Save(ctx)
	if err != nil {
		fmt.Println("UploadDocument err:", err)
		return nil, fmt.Errorf("an error occurred while creating the file")
	}

	return document, nil
}

// UpdateDocument is the resolver for the updateDocument field.
func (r *mutationResolver) UpdateDocument(ctx context.Context, id int, input generated.UpdateCompanyDocumentInput) (*generated.CompanyDocument, error) {
	_, currentCompany := utils.GetSession(&ctx)
	companyFilter := companydocument.HasCompanyWith(company.ID(currentCompany.ID))
	document, err := r.client.CompanyDocument.UpdateOneID(id).Where(companyFilter).
		SetInput(input).
		Save(ctx)
	if err != nil {
		fmt.Println("UpdateDocument err:", err)
		return nil, fmt.Errorf("an error occurred while updating the document")
	}
	return document, nil
}

// DeleteDocuments is the resolver for the deleteDocuments field.
func (r *mutationResolver) DeleteDocuments(ctx context.Context, ids []int) (*string, error) {
	_, currentCompany := utils.GetSession(&ctx)
	companyFilter := companydocument.HasCompanyWith(company.ID(currentCompany.ID))
	n, err := r.client.CompanyDocument.Delete().
		Where(companyFilter, companydocument.IDIn(ids...)).
		Exec(ctx)
	if err != nil {
		fmt.Println("DeleteDocuments err:", err)
		return nil, err
	}

	result := fmt.Sprintf("%d documents were deleted", n)
	return &result, nil
}

// GenerateMemberSignupToken is the resolver for the generateMemberSignupToken field.
func (r *mutationResolver) GenerateMemberSignupToken(ctx context.Context, input generated.CreateMemberSignupTokenInput) (*model.CreateMemberSignupTokenOutput, error) {
	currentUser, currentCompany := utils.GetSession(&ctx)

	token := (func() string {
		token := appUtils.GenerateRandomToken()

		// Check if token is unique. Try 15 times then give up.
		N := 15
		for i := range N {
			exists, _ := r.client.MemberSignupToken.Query().Where(membersignuptoken.Token(token)).Exist(ctx)
			if exists && i < 14 {
				token = appUtils.GenerateRandomToken()
			} else if exists {
				token = ""
			} else {
				break
			}
		}
		return token
	})()

	if token == "" {
		return nil, fmt.Errorf("cannot generate token. Try again")
	}

	newToken, err := r.client.MemberSignupToken.Create().
		SetInput(input).
		SetToken(token).
		SetCompanyID(currentCompany.ID).SetCreatedByID(currentUser.ID).
		// SetNumberAccessed(0).
		SetExpiresAt(time.Now().Add(time.Duration(time.Hour * 48))). // token expires in 48 hours
		Save(ctx)
	if err != nil {
		fmt.Println("GenerateMemberSignupToken err:", err)
		return nil, fmt.Errorf("ocorreu um erro")
	}

	output := model.CreateMemberSignupTokenOutput{
		Token:     newToken.Token,
		ExpiresAt: newToken.ExpiresAt,
	}
	return &output, nil
}

// DeleteMemberSignupToken is the resolver for the deleteMemberSignupToken field.
func (r *mutationResolver) DeleteMemberSignupToken(ctx context.Context, id int) (*string, error) {
	_, currentCompany := utils.GetSession(&ctx)
	companyFilter := membersignuptoken.HasCompanyWith(company.ID(currentCompany.ID))
	err := r.client.MemberSignupToken.DeleteOneID(id).Where(companyFilter).Exec(ctx)
	if err != nil {
		fmt.Println("DeleteMemberSignupToken err:", err)
		return nil, fmt.Errorf("ocorreu um erro")
	}

	result := "token deleted"
	return &result, nil
}

// RemoveUser is the resolver for the removeUser field.
func (r *mutationResolver) RemoveUser(ctx context.Context, id int) (*string, error) {
	currentUser, currentCompany := utils.GetSession(&ctx)

	// Check if active user is an admin of the company. If not, return;
	_, err := currentUser.QueryAssignedRoles().Where(
		userrole.HasCompanyWith(company.ID(currentCompany.ID)),
		userrole.RoleIn(userrole.RoleADMIN, userrole.RoleSUPERUSER),
	).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	targetUser, err := r.client.User.Query().
		Where(user.ID(id), user.HasCompanyWith(company.ID(currentCompany.ID))).
		WithCompany(func(cq *generated.CompanyQuery) {
			cq.Select(company.FieldID)
		}).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	// If the target user is only connect to this company, remove his account.
	// Otherwise, only disconnect him from this company
	if len(targetUser.Edges.Company) <= 1 {
		err = r.client.User.DeleteOneID(targetUser.ID).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("user not deleted")
		}
	} else {
		_, err = targetUser.Update().RemoveCompanyIDs(currentCompany.ID).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("user not deleted")
		}
	}

	// Notify target user about his removal from this company
	// ....

	result := fmt.Sprintf("user removed from %s", currentCompany.Name)
	return &result, nil
}

// CountDocuments is the resolver for the countDocuments field.
func (r *queryResolver) CountDocuments(ctx context.Context) ([]*model.DocumentCount, error) {
	_, currentCompany := utils.GetSession(&ctx)
	var output []*model.DocumentCount
	companyFilter := companydocument.HasCompanyWith(company.ID(currentCompany.ID))

	err := r.client.CompanyDocument.Query().Where(companyFilter).
		GroupBy(companydocument.FieldCategory).
		Aggregate(generated.Count()).
		Scan(ctx, &output)
	if err != nil {
		fmt.Println("CountDocuments err:", err)
		return nil, err
	}

	return output, nil
}

// VerifyMemberSignupToken is the resolver for the verifyMemberSignupToken field.
func (r *queryResolver) VerifyMemberSignupToken(ctx context.Context, input *model.VerifyMemberSignupTokenInput) (*generated.MemberSignupToken, error) {
	filter := []predicate.MemberSignupToken{
		membersignuptoken.Token(input.Token),
		membersignuptoken.AlreadyUsed(false),
		membersignuptoken.ExpiresAtGTE(time.Now()),
	}
	if input.Email == nil {
		filter = append(filter, membersignuptoken.EmailIsNil())
	} else {
		filter = append(filter, membersignuptoken.Or(
			membersignuptoken.EmailIsNil(),
			membersignuptoken.EmailEQ(*input.Email),
		))
	}

	currentToken, err := r.client.MemberSignupToken.Query().
		Where(filter...).
		First(ctx)
	if err != nil {
		fmt.Println("VerifyMemberSignupToken update err:", err)
		return nil, fmt.Errorf("invalid token")
	}
	_, _ = currentToken.Update().AddNumberAccessed(1).Save(ctx)

	return currentToken, nil
}
