package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
