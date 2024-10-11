package middlewares

import (
	"fmt"
	"mazza/inits"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println("err 2:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := inits.Client.User.Get(c, int(claims["id"].(float64)))
		if err != nil || user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		company, err := inits.Client.Company.Get(c, int(claims["companyID"].(float64)))
		if company.ID == 0 || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Set("company", company)
		c.Set("companyID", company.ID)
		c.Set("employeeID", int(claims["employeeID"].(float64)))
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
