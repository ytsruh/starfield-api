package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the Starfield API",
		})
	})
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Live Server Metrics"}))
}
