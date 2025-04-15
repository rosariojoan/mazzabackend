package middlewares

import (
	"fmt"
	"mazza/app/utils"
	"mazza/inits"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/* A valid access token should have a user ID and a company ID. */
func LoginRequired(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		_tokenString, exists := c.Request.Header["Authorization"]
		if !exists {
			fmt.Println("No Authorization header provided")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// tokenType = strings.Split(_tokenString[0], " ")[0]
		tokenString = strings.Split(_tokenString[0], " ")[1]
	}

	claims, err := utils.ParseJWTToken(tokenString)
	if err != nil {
		fmt.Println("err 2:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := inits.Client.User.Get(c, int(claims["id"].(float64)))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// http.StatusForbidden
	company, err := inits.Client.Company.Get(c, int(claims["companyID"].(float64)))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Set("company", company)
	c.Set("companyID", company.ID)

	c.Next()
}
