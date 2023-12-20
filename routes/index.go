package routes

import (
	"fmt"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	database "starfieldapi.com/db"
	"starfieldapi.com/routes/api"
	"starfieldapi.com/routes/dashboard"
)

func SetRoutes(app *fiber.App) {
	// Setup & use Swagger for documentation
	app.Use(swagger.New(swagger.Config{
		Next:     nil,
		BasePath: "/",
		FilePath: "./swagger.json",
		Path:     "docs",
		Title:    "Starfield API documentation",
		CacheAge: 3600,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"PageTitle": "Welcome to the Stafield API",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"PageTitle": "About the project and additional information",
		})
	})

	// app.Get("/register", func(c *fiber.Ctx) error {
	// 	return c.Render("register", fiber.Map{
	// 		"PageTitle": "Register for the Stafield API",
	// 	})
	// })
	//app.Post("/register", registerUser)
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"PageTitle": "Login to the Stafield API",
		})
	})
	app.Post("/login", loginUser)
	app.Get("/logout", logoutUser)

	app.Get("/request-reset", func(c *fiber.Ctx) error {
		return c.Render("requestreset", fiber.Map{
			"PageTitle": "Request a link to reset your password",
		})
	})
	app.Post("/request-reset", requestPasswordReset)

	app.Get("/reset-password", func(c *fiber.Ctx) error {
		resetQuery := c.Query("reset")
		// Check if a reset query parameter exists
		if resetQuery == "" {
			return c.Render("resetpassword", fiber.Map{
				"PageTitle": "Reset your password reset the Stafield API",
				"Error":     "Error: Password reset request not found",
			})
		}
		// Check if the id is valid
		reset, err := database.GetPasswordReset(resetQuery)
		if err != nil {
			fmt.Println(err)
			return c.Render("resetpassword", fiber.Map{
				"PageTitle": "Reset your password reset the Stafield API",
				"Error":     "Error: Password reset request not found",
			})
		}

		return c.Render("resetpassword", fiber.Map{
			"PageTitle": "Reset your password reset the Stafield API",
			"Reset":     reset,
		})
	})
	app.Post("/reset-password", resetPassword)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Live Server Metrics"}))

	// Setup Dashboard & Public API routes
	api.Setup(app)
	dashboard.Setup(app)
}
