package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	database "starfieldapi.com/db"
	"starfieldapi.com/lib"
)

type CustomClaims struct {
	User                 string `json:"user"`
	Id                   string `json:"id"`
	jwt.RegisteredClaims `json:"claims"`
}

func registerUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "issue getting input data",
		})
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "issue creating password",
		})
	}
	user := database.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	error := database.CreateUser(&user)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create new url" + error.Error(),
		})
	}

	return c.JSON(user)
}

func loginUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := database.GetUserByEmail(data["email"])
	if user.Email == "" || err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := CustomClaims{
		user.Email,
		user.Id.String(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "search-engine",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	secretKey := lib.GetSecretKey()
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   signedToken,
	})
}

func getUser(c *fiber.Ctx) error {
	details := c.Locals("user").(*CustomClaims) // Get local context & type assertion to custom claims struct
	user, err := database.GetUserByEmail(details.User)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to look up user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"email":   user.Email,
		"id":      user.Id,
		"name":    user.Name,
	})
}
