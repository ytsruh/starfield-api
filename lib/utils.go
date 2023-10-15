package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func PrintJSON(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error printing JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}

func GetSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable is not set")
	}
	return secretKey
}

type CustomClaims struct {
	User                 string `json:"user"`
	Id                   string `json:"id"`
	jwt.RegisteredClaims `json:"claims"`
}
