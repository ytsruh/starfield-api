package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	database "starfieldapi.com/db"
	"starfieldapi.com/lib"
)

func registerUser(c *fiber.Ctx) error {
	type Payload struct {
		Name     string `json:"name" xml:"name" form:"name"`
		Email    string `json:"email" xml:"email" form:"email"`
		Password string `json:"password" xml:"password" form:"password"`
	}
	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		c.Status(500)
		return c.Render("register", fiber.Map{
			"Message": "There was an error with the form",
		})
	}
	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		c.Status(500)
		return c.Render("register", fiber.Map{
			"Message": "There was issue creating your password",
		})
	}
	user := database.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: password,
	}
	createErr := database.CreateUser(&user)
	if createErr != nil {
		c.Status(500)
		return c.Render("register", fiber.Map{
			"Message": "There was error creating your account",
		})
	}
	emailErr := lib.EmailCreateUser(payload.Email)
	if emailErr != nil {
		fmt.Print(emailErr)
	}
	c.Status(200)
	return c.Render("login", fiber.Map{
		"Message": "Success",
	})
}

func loginUser(c *fiber.Ctx) error {
	type Payload struct {
		Email    string `json:"email" xml:"email" form:"email"`
		Password string `json:"password" xml:"password" form:"password"`
	}
	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		c.Status(500)
		return c.Render("login", fiber.Map{
			"Message": "There was an error with the form",
		})
	}

	user, err := database.GetUserByEmail(payload.Email)
	if user.Email == "" || err != nil {
		c.Status(401)
		return c.Render("login", fiber.Map{
			"Message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		c.Status(401)
		return c.Render("login", fiber.Map{
			"Message": "Incorrect password provided",
		})
	}

	claims := lib.CustomClaims{
		User: user.Email,
		Id:   user.Id.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "starfield-api",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	secretKey := lib.GetSecretKey()
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.Status(500)
		return c.Render("login", fiber.Map{
			"Message": "Something went wrong logging in, please try again.",
		})
	}

	// Create and set the cookie
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true, // Meant only for the server
	}
	c.Cookie(&cookie)

	return c.Redirect("/dashboard")
}

func logoutUser(c *fiber.Ctx) error {
	// Create & set new cookie with expired date
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true, // Meant only for the server
	}
	c.Cookie(&cookie)
	return c.Redirect("/login")
}

func requestPasswordReset(c *fiber.Ctx) error {
	type Payload struct {
		Email string `json:"email" xml:"email" form:"email"`
	}
	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		c.Status(500)
		return c.Render("requestreset", fiber.Map{
			"Message": "There was an error with the form",
		})
	}
	reset := database.PasswordReset{
		Email: payload.Email,
	}

	createErr := database.NewPasswordReset(&reset)

	if createErr != nil {
		c.Status(500)
		return c.Render("requestreset", fiber.Map{
			"Message": "There was error resetting your password",
		})
	}

	lib.EmailPasswordResetLink(reset.Id.String(), payload.Email)

	c.Status(200)
	return c.Render("requestreset", fiber.Map{
		"Message": "Success. An Email will be sent with a link to reset your password",
	})
}

func resetPassword(c *fiber.Ctx) error {
	type Payload struct {
		ResetId         string `json:"resetid" xml:"resetid" form:"resetid"`
		NewPassword     string `json:"newpassword" xml:"newpassword" form:"newpassword"`
		ConfirmPassword string `json:"confirmpassword" xml:"confirmpassword" form:"confirmpassword"`
	}
	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		c.Status(500)
		return c.Render("resetpassword", fiber.Map{
			"Message": "There was an error with the form",
			"Reset": fiber.Map{
				"Id": payload.ResetId,
			},
		})
	}

	if payload.NewPassword != payload.ConfirmPassword {
		c.Status(500)
		return c.Render("resetpassword", fiber.Map{
			"Message": "The passwords do not match",
			"Reset": fiber.Map{
				"Id": payload.ResetId,
			},
		})
	}

	reset, resetErr := database.GetPasswordReset(payload.ResetId)
	if resetErr != nil {
		c.Status(500)
		return c.Render("resetpassword", fiber.Map{
			"Error": "Error: Password reset request not found",
			"Reset": fiber.Map{
				"Id": payload.ResetId,
			},
		})
	}

	password, passwordErr := bcrypt.GenerateFromPassword([]byte(payload.NewPassword), 14)
	if passwordErr != nil {
		c.Status(500)
		return c.Render("resetpassword", fiber.Map{
			"Error": "Error: There was issue creating your new password",
			"Reset": fiber.Map{
				"Id": payload.ResetId,
			},
		})
	}

	user := database.User{
		Email:    reset.Email,
		Password: password,
	}

	updateErr := database.UpdateUserPassword(&user)
	if updateErr != nil {
		c.Status(500)
		return c.Render("resetpassword", fiber.Map{
			"Error": "Error: There was issue updating your new password",
			"Reset": fiber.Map{
				"Id": payload.ResetId,
			},
		})
	}

	deleteErr := database.DeletePasswordReset(string(reset.Id.String()))
	if deleteErr != nil {
		fmt.Println("Password reset was not deleted: " + reset.Id.String())
	}

	return c.Redirect("/login")
}
