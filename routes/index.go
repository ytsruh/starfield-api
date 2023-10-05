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
	// Star Systems
	app.Get("/starsystem", getStarSystems)
	app.Get("/starsystem/:id", getSingleStarSystem)
	app.Post("/starsystem", createStarSystem)
	app.Put("/starsystem/:id", updateStarSystem)
	app.Delete("/starsystem/:id", deleteStarSystem)
	// Planets
	app.Get("/planet", getPlanets)
	app.Get("/planet/:id", getSinglePlanet)
	app.Post("/planet", createPlanet)
	app.Put("/planet/:id", updatePlanet)
	app.Delete("/planet/:id", deletePlanet)
	// Moons
	app.Get("/moon", getMoons)
	app.Get("/moon/:id", getSingleMoon)
	app.Post("/moon", createMoon)
	app.Put("/moon/:id", updateMoon)
	app.Delete("/moon/:id", deleteMoon)
}
