package auth

import (
	"context"
	"fmt"
	"mazza/app/utils"
	"mazza/ent/generated/company"
	"mazza/ent/generated/token"
	"mazza/ent/generated/user"
	"mazza/inits"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ConfirmDeleteUser(ctx *gin.Context) {
	var input struct {
		Token string
	}
	if err := ctx.BindJSON(&input); err != nil || input.Token == "" {
		fmt.Println("err:", err, input)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Parse JWT token
	claims, err := utils.ParseJWTToken(input.Token)
	if err != nil {
		fmt.Println("err 2:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid input"})
		return
	}

	if claims["userID"] == nil {
		fmt.Println("err 3:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid input"})
		return
	}

	// Retrieve token and user
	userID := int(claims["userID"].(float64))
	activeToken, err := inits.Client.Token.Query().Where(
		token.Token(input.Token),
		token.HasUserWith(user.ID(userID)),
		token.ExpiryGTE(time.Now()),
		token.CategoryEQ(token.CategoryAccountDelete),
	).WithUser().First(ctx)
	if err != nil {
		fmt.Println("err 4:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid input"})
		return
	}

	activeUser := activeToken.Edges.User
	activeCompany, err := inits.Client.Company.Query().Where(
		company.HasUsersWith(user.ID(activeUser.ID)),
	).First(ctx)
	if err != nil {
		fmt.Println("err 5:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "an error occurred"})
		return
	}

	// Delete account
	deleteCtx := context.Background()
	err = Unsubscribe(&deleteCtx, inits.Client, activeUser, activeCompany)
	if err != nil {
		fmt.Println("err 5:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "an error occurred"})
		return
	}

	// Invalidate token entry
	_ = inits.Client.Token.DeleteOneID(activeToken.ID).Exec(deleteCtx)

	// Return success
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
