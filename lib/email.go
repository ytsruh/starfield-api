package lib

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func EmailCreateUser(email string) error {
	payload := strings.NewReader(fmt.Sprintf("{\n  \"event\": \"registration\",\n  \"email\": \"%s\"\n}", email))

	req, reqErr := http.NewRequest("POST", "https://api.useplunk.com/v1/track", payload)

	if reqErr != nil {
		return reqErr
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("PLUNK_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, resErr := http.DefaultClient.Do(req)

	if resErr != nil {
		return resErr
	}

	defer res.Body.Close()

	return nil

}

func EmailPasswordResetLink(id string) {
	fmt.Println("ID : " + id)
	fmt.Println("Sending Password Reset Link")
}
