package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

/*
Generate JWT token to be used for various purposes.

Normally, you want the payload to contain: user ID, company ID, and expiration date
*/
func GenerateJWTToken(duration time.Duration, payload jwt.MapClaims) (token string, err error) {
	payload["exp"] = time.Now().Add(time.Duration(duration)).Unix()

	// Generate JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err = jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return token, err
}
