package auth

import (
	"fmt"
	ent "mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/user"
	"mazza/ent/generated/userrole"
	"mazza/inits"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSession(ctx *gin.Context) {
	_user, _ := ctx.Get("user")
	if _user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	_companyID, _ := ctx.Get("companyID")
	if _companyID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	activeUser := _user.(*ent.User)
	companyID := _companyID.(int)

	// Get the active company
	companies, err := inits.Client.User.Query().Where(user.IDEQ(activeUser.ID)).QueryCompany().All(ctx)
	if err != nil || len(companies) == 0 {
		// return nil, fmt.Errorf("there are no companies associated with this user")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	roles, err := activeUser.QueryAssignedRoles().Where(userrole.HasCompanyWith(company.ID(companyID))).All(ctx)
	if err != nil {
		fmt.Println("roles err:", err)
		roles = []*ent.UserRole{}
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		// return
	}

	response := LoginOutput{
		User:         *activeUser,
		Roles:        roles,
		CompanyID:    companyID,
		Companies:    companies,
		AccessToken:  "",
		RefreshToken: "",
		TTL:          30000,
	}
	ctx.JSON(http.StatusOK, response)
}
