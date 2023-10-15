package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	database "starfieldapi.com/db"
	"starfieldapi.com/lib"
)

func Setup(app *fiber.App) {
	secretKey := lib.GetSecretKey()
	dash := app.Group("/dashboard")

	dash.Use(func(c *fiber.Ctx) error {
		// Get the cookie by name
		cookie := c.Cookies("auth")
		// Parse the cookie & check for errors
		token, err := jwt.ParseWithClaims(cookie, &lib.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.Status(401)
			return c.Redirect("login")
		}
		// Parse the custom claims & check jwt is valid
		claims, ok := token.Claims.(*lib.CustomClaims)
		if ok && token.Valid {
			c.Locals("user", claims)
			return c.Next()
		}
		// Return unauthorized if jwt is not valid
		c.Status(401)
		return c.Redirect("/login")
	})

	dash.Use("*", func(c *fiber.Ctx) error {
		c.Bind(fiber.Map{
			"LoggedIn": true,
		})
		return c.Next()
	})

	dash.Get("/", func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*lib.CustomClaims)
		if !ok {
			return c.Redirect("/500")
		}
		user, err := database.GetUserByEmail(claims.User)
		if err != nil {
			return c.Redirect("/login")
		}
		return c.Render("dashboard", fiber.Map{
			"PageTitle": "Dashboard",
			"User":      user,
		})
	})

	dash.Post("/", func(c *fiber.Ctx) error {
		var payload database.User
		if err := c.BodyParser(&payload); err != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		claims, ok := c.Locals("user").(*lib.CustomClaims)
		if !ok {
			c.Status(500)
			return c.Redirect("/500")
		}
		id, err := uuid.Parse(claims.Id)
		if err != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		updatedUser := database.User{
			Id:    id,
			Name:  payload.Name,
			Email: claims.User,
		}
		dbErr := database.UpdateUser(&updatedUser)
		if dbErr != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		c.Status(200)
		return c.Render("dashboard", fiber.Map{
			"Message": "Success",
			"User":    updatedUser,
		})
	})

	dash.Get("/keys", func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*lib.CustomClaims)
		if !ok {
			return c.Redirect("/500")
		}
		keys, err := database.GetKeysByUser(claims.Id)
		if err != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		return c.Render("keys", fiber.Map{
			"PageTitle": "API Keys",
			"Keys":      keys,
		})
	})

	dash.Post("/keys", func(c *fiber.Ctx) error {
		type Payload struct {
			Name string
		}
		var payload Payload
		if err := c.BodyParser(&payload); err != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		claims, ok := c.Locals("user").(*lib.CustomClaims)
		if !ok {
			c.Status(500)
			return c.Redirect("/500")
		}
		id, err := uuid.Parse(claims.Id)
		if err != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		key, err := lib.GenRandomString(48)
		if err != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		createErr := database.CreateKey(&database.APIKey{
			Name:   payload.Name,
			Key:    key,
			UserId: id,
		})
		if createErr != nil {
			c.Status(500)
			return c.Redirect("/500")
		}
		return c.Redirect("/dashboard/keys")
	})

}
