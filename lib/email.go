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

func EmailPasswordResetLink(id string, email string) {
	baseUrl := os.Getenv("BASE_URL")
	link := fmt.Sprintf("<div><p>Your password reset link is <a href='%s/reset-password?reset=%s'>here</a></p></div>", baseUrl, id)

	payload := strings.NewReader(fmt.Sprintf("{\n  \"to\": \"%s\",\n  \"subject\": \"Starfield API Password Reset\",\n  \"body\": \"%s\"\n}", email, link))

	req, reqErr := http.NewRequest("POST", "https://api.useplunk.com/v1/send", payload)

	if reqErr != nil {
		fmt.Println("Req err")
		fmt.Println(reqErr)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("PLUNK_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, resErr := http.DefaultClient.Do(req)

	if resErr != nil {
		fmt.Println("Res err")
		fmt.Println(resErr)
	}

	res.Body.Close()

}
