package routes

import (
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
	error := database.CreateUser(&user)
	if error != nil {
		c.Status(500)
		return c.Render("register", fiber.Map{
			"Message": "There was error creating your account",
		})
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
