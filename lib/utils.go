package lib

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
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

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	maxIndex := big.NewInt(int64(len(letterBytes)))

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", err
		}
		bytes[i] = letterBytes[index.Int64()]
	}

	return string(bytes), nil
}
