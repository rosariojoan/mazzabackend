package auth

import (
	"fmt"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/user"
	"mazza/ent/generated/userrole"

	"mazza/inits"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	FirebaseUID   string
	ExpoPushToken *string
}

type LoginOutput struct {
	User         ent.User        `json:"user"`
	Roles        []*ent.UserRole `json:"userRoles"`
	CompanyID    int             `json:"activeCompanyId"`
	Companies    []*ent.Company  `json:"companies"`
	AccessToken  string          `json:"accessToken"`
	RefreshToken string          `json:"refreshToken"`
	TTL          int             `json:"ttl"`
}

/* User login via JWT token */
// func Login(ctx *context.Context, username string, password string, ExpoPushToken *string) (data *LoginOutput, err error) {
func Login(ctx *gin.Context) {
	// client := ent.FromContext(ctx)
	var input LoginInput

	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	currentUser, err := inits.Client.User.Query().
		Where(user.And(
			user.FirebaseUID(input.FirebaseUID),
			user.DeletedAtIsNil(),
		)).
		// WithAssignedRoles().
		First(ctx)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// err = utils.CompareHashAndPassword(currentUser.Password, input.Password)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
	// 	return
	// }

	// Get the active company
	companies, err := inits.Client.User.Query().Where(user.IDEQ(currentUser.ID)).QueryCompany().All(ctx)
	if err != nil || len(companies) == 0 {
		// return nil, fmt.Errorf("there are no companies associated with this user")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	activeCompanyID := companies[0].ID

	// Generate JWT token
	duration := time.Hour * 24 * 30 // valid for 30 days
	payload := jwt.MapClaims{
		"id":        currentUser.ID,
		"companyID": activeCompanyID,
	}

	token, err := utils.GenerateJWTToken(duration, payload)
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	// fmt.Println(currentUser)
	// fmt.Println(companies)

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*24*30, "", "localhost", false, true)

	ttl, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRE"))
	if err != nil {
		ttl = 2592000 // 30 days
	}

	// Update user ExpoPushToken
	if input.ExpoPushToken != nil {
		err = inits.Client.User.UpdateOneID(currentUser.ID).SetExpoPushToken(*input.ExpoPushToken).Exec(ctx)
		if err != nil {
			fmt.Println("err:", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ocorreu um erro ao fazer o login"})
			return
		}
		inits.Client.User.Update().Where(
			user.IDNEQ(currentUser.ID),
			user.ExpoPushToken(*input.ExpoPushToken),
		).ClearExpoPushToken().Exec(ctx)
	} else {
		err = inits.Client.User.UpdateOneID(currentUser.ID).ClearExpoPushToken().Exec(ctx)
		if err != nil {
			fmt.Println("err:", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ocorreu um erro ao fazer o login"})
			return
		}
	}

	roles, err := currentUser.QueryAssignedRoles().Where(userrole.HasCompanyWith(company.ID(activeCompanyID))).All(ctx)
	if err != nil {
		fmt.Println("roles err:", err)
		roles = []*ent.UserRole{}
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		// return
	}

	currentUser, err = currentUser.Update().SetLastLogin(time.Now()).Save(ctx)
	if err != nil {
		fmt.Println("Login user SetLastLogin update err:", err)
	}

	response := LoginOutput{
		User:         *currentUser,
		Roles:        roles,
		CompanyID:    activeCompanyID,
		Companies:    companies,
		AccessToken:  token,
		RefreshToken: "",
		TTL:          ttl,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
