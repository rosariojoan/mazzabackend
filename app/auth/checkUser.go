package auth

import (
	"context"
	"fmt"
	"mazza/ent/generated/user"
	"mazza/inits"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserCheck struct {
	Email string `json:"email"`
}

// Check if user exists
func CheckUser(ctx *gin.Context) {
	var input UserCheck

	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "not available"})
		return
	}

	_, err := inits.Client.User.Query().Where(user.EmailEqualFold(input.Email)).FirstID(ctx)
	if err != nil {
		fmt.Println("err:", err)

		// If user does not exist, check if it is registered on Firebase.
		// If it does, remove it from Firebase
		if strings.Contains(err.Error(), "user not found") {
			ctx := context.Background()
			// Remove user from Firebase
			firebaseUser, err := inits.Auth.GetUserByEmail(ctx, input.Email)
			if err == nil {
				inits.Auth.DeleteUser(ctx, firebaseUser.UID)
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"email_already_taken": false}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"email_already_taken": true}})
}
