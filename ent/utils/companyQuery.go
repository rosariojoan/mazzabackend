package utils

import (
	"context"
	"mazza/ent/generated/company"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/user"
	"mazza/ent/generated/userrole"
)

/* User role query filter: return a filter to get the predicate.UserRole of the current user. */
func UserRoleQuery(ctx *context.Context) predicate.UserRole {
	currentUser, _ := GetSession(ctx)
	roleQ := userrole.HasUserWith(user.IDEQ(currentUser.ID))
	return roleQ
}

/*
Company query filter: return a filter to get all predicate.Company associated with the current user.
*/
func CompanyQuery(ctx *context.Context) predicate.Company {
	currentUser, _ := GetSession(ctx)
	userQ := userrole.HasUserWith(user.IDEQ(currentUser.ID))
	companyQ := company.HasAvailableRolesWith(userQ)

	return companyQ
}

/*
Current company query: return a filter to get the predicate.Company of the current company.
*/
func CurrentCompanyQuery(ctx *context.Context) predicate.Company {
	_, currentCompany := GetSession(ctx)
	return company.IDEQ(currentCompany.ID)
}
