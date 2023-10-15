package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
		return c.Render("dashboard", fiber.Map{
			"PageTitle": "Dashboard",
		})
	})

	dash.Get("/keys", func(c *fiber.Ctx) error {
		return c.Render("keys", fiber.Map{
			"PageTitle": "API Keys",
		})
	})

}
