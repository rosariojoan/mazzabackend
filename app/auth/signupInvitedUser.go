package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"

	// generated "mazza/ent/generated"

	"mazza/app/utils"
	"mazza/ent/generated"
	ent "mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/employee"
	"mazza/ent/generated/membersignuptoken"
	"mazza/ent/generated/user"
	"mazza/ent/generated/userrole"
	"mazza/firebase"
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

	// Create new user and employee
	department := "geral"
	phone := utils.GetValue(input.UserInput.Phone, "")
	userIsActive := false

	newUser, err := tx.User.Create().SetInput(*input.UserInput).
		AddCompanyIDs(companyID).
		SetActive(userIsActive).
		SetLeaderID(inviteeID).
		AddAssignedRoles(
			tx.UserRole.Create().SetInput(roleInput).SaveX(ctx), // create new role and assign it to the user
		).Save(ctx)
	if err != nil {
		fmt.Println("InvitedUserSignup create user err:", err)
		_ = tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred"})
		return
	}

	// If the user is not an auditor, create an employee record for him
	if token.Role != membersignuptoken.Role(userrole.RoleAUDITOR) {
		newUser.Update().SetEmployee(
			tx.Employee.Create().SetInput(ent.CreateEmployeeInput{
				Name:       input.UserInput.Name,
				Birthdate:  input.UserInput.Birthdate,
				Gender:     employee.GenderMale,
				HireDate:   token.CreatedAt,
				Department: &department,
				Phone:      &phone,
				CompanyID:  &companyID,
			}).SaveX(ctx),
		)
	}

	// Invalidate token
	_, err = tx.Client().MemberSignupToken.UpdateOneID(token.ID).SetAlreadyUsed(true).AddNumberAccessed(1).Save(ctx)
	if err != nil {
		fmt.Println("InvitedUserSignup token invalidation err:", err)
		_ = tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred"})
		return
	}

	// Create a new user entry in the Firestore database
	err = firebase.CreateUserEntry(ctx, companyID, newUser.FirebaseUID, userIsActive, roleInput.Role)
	if err != nil {
		fmt.Println("could not create user:", err)
		_ = tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred"})
		return
	}

	err = tx.Commit()
	// tx.Rollback()
	if err != nil {
		fmt.Println("err:", err)
		_ = tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "an error occurred"})
		return
	}

	// Notify to admin user
	go (func() {
		adminUsers, err := inits.Client.User.Query().Where(
			user.HasCompanyWith(company.ID(companyID)),
			user.Active(true),
			user.IDNEQ(newUser.ID),
			user.ExpoPushTokenNotNil(),
			user.HasAssignedRolesWith(
				userrole.RoleEQ(userrole.RoleADMIN),
				userrole.HasCompanyWith(company.ID(companyID)),
			),
		).All(ctx)

		if err == nil && len(adminUsers) > 0 {
			var tokens []expo.ExponentPushToken
			for _, admin := range adminUsers {
				tokens = append(tokens, expo.ExponentPushToken(*admin.ExpoPushToken))
			}
			_, err := inits.ExpoClient.Publish(&expo.PushMessage{
				To:    tokens,
				Title: "Novo utilizador",
				Body:  fmt.Sprintf("%s registou-se na sua empresa. Podes activar a conta do novo utilizador.", newUser.Unwrap().Name),
				Data: map[string]string{
					"type":      firebase.AlertType.InvitedUserRegistration,
					"userID":    strconv.Itoa(newUser.ID),
					"companyID": strconv.Itoa(companyID),
				},
				Sound:    "default",
				Priority: expo.HighPriority,
			})
			if err != nil {
				fmt.Println("send notif err:", err)
			}
			// response.
		}
	})()

	ctx.JSON(http.StatusCreated, gin.H{"data": "usuário registado"})
}
