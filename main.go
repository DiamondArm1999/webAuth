package main

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:password")))
}

func hashPassword(password string) []byte {
	bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
