package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
)

func main() {
	key := make([]byte, 64)
	_, err := rand.Read(key)

	if err != nil {
		panic("Error generating secret key")
	}

	decodedKey := b64.URLEncoding.EncodeToString(key)
	fmt.Println(decodedKey)
}
