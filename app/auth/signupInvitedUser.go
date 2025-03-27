package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// generated "mazza/ent/generated"

	"mazza/ent/generated"
	"mazza/ent/generated/userrole"
	"mazza/inits"
	"mazza/mazza/generated/model"
)

func SignupInvitedUser(ctx *gin.Context) {
	var input model.InvitedUserSignupInput

	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "houve um erro ao registar usuário"})
		return
	}
	// if input.InvitationToken

	tx, err := inits.Client.Tx(ctx)
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "houve um erro ao registar usuário"})
		return
	}

	token, err := inits.Client.MemberSignupToken.Get(ctx, input.InvitationToken)
	if err != nil {
		fmt.Println("InvitedUserSignup get token err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	companyID, err := token.QueryCompany().FirstID(ctx)
	if err != nil {
		fmt.Println("InvitedUserSignup get company err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	inviteeID, err := token.QueryCreatedBy().FirstID(ctx)
	if err != nil {
		fmt.Println("InvitedUserSignup get invitee err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	roleInput := generated.CreateUserRoleInput{
		Role:      userrole.Role(token.Role),
		Notes:     token.Note,
		CompanyID: &companyID,
	}

	// Create new user
	_, err = tx.User.Create().SetInput(*input.UserInput).
		AddCompanyIDs(companyID).
		SetActive(false).SetLastLogin(time.Now()).SetLeaderID(inviteeID).
		AddAssignedRoles(
			tx.UserRole.Create().SetInput(roleInput).SaveX(ctx),
		). // create new role and assign it to the user
		Save(ctx)
	if err != nil {
		fmt.Println("InvitedUserSignup create user err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred"})
		return
	}

	// Invalidate token
	_, err = tx.Client().MemberSignupToken.UpdateOneID(token.ID).SetAlreadyUsed(true).AddNumberAccessed(1).Save(ctx)
	if err != nil {
		fmt.Println("InvitedUserSignup token invalidation err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred"})
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "an error occurred"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": "usuário registado"})
}
