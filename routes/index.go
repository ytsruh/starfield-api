package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"starfieldapi.com/routes/api"
	"starfieldapi.com/routes/dashboard"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"PageTitle": "Welcome to the Stafield API",
		})
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{
			"PageTitle": "Register for the Stafield API",
		})
	})
	app.Post("/register", registerUser)
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"PageTitle": "Login to the Stafield API",
		})
	})
	app.Post("/login", loginUser)
	app.Get("/logout", logoutUser)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Live Server Metrics"}))

	// Setup Dashboard & Public API routes
	api.Setup(app)
	dashboard.Setup(app)
}
