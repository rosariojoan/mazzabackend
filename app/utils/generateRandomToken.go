package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomToken() string {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return "123456"
	}

	// Generate a number between 0 - 999_999
	tokenNum := (int(b[0])<<16 | int(b[1])<<8 | int(b[2])) % 1_000_000

	// Format the stoken as xxx-xxx string with leading zeros if needed
	return fmt.Sprintf("%03d-%03d", tokenNum/1000, tokenNum%1000)
}
