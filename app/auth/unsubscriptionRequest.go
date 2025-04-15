package auth

import (
	"fmt"
	"mazza/app/utils"
	"mazza/app/utils/email"
	"mazza/ent/generated"
	"mazza/ent/generated/token"
	"mazza/ent/generated/user"
	"mazza/inits"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Request account deletion. Create an expiring token if the user exists.
// This token will be used to confirm the account deletion
func UnsubscriptionRequest(ctx *gin.Context) {
	var input struct {
		Email string
	}
	if err := ctx.BindJSON(&input); err != nil || input.Email == "" {
		fmt.Println("err:", err, input)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	targetUser, err := inits.Client.User.Query().Where(user.Email(input.Email)).First(ctx)
	if err != nil {
		fmt.Println("err:", err, input)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Generate JWT token: Create a token with 2 hour expiry
	duration := time.Hour * 2 // valid for 30 days
	expiry := time.Now().Add(duration)
	payload := jwt.MapClaims{
		"userID": targetUser.ID,
	}

	jwtToken, err := utils.GenerateJWTToken(duration, payload)
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "an error occurred"})
		return
	}

	_, err = inits.Client.Token.Create().SetInput(generated.CreateTokenInput{
		Category: token.CategoryACCOUNT_DELETE,
		UserID:   &targetUser.ID,
		Token:    jwtToken,
		Expiry:   expiry,
	}).Save(ctx)

	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "an error occurred"})
		return
	}

	// Send email to user
	validity := fmt.Sprintf("%d horas", duration)
	go email.SendAccountDeleteConfirmation(targetUser.Name, input.Email, jwtToken, validity)

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
