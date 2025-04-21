package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// Parse a JWT token.
// Check if token is valid and not expired.
// Return JWT MapClaims
func ParseJWTToken(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if token is valid
	claims, ok := token.Claims.(jwt.MapClaims); 
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check if token is expired
	if claims["exp"] == nil || float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("expired")
	}

	return claims, nil
}
