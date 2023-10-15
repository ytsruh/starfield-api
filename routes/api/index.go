package api

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	// Star Systems
	api.Get("/starsystem", getStarSystems)
	api.Get("/starsystem/:id", getSingleStarSystem)
	api.Post("/starsystem", createStarSystem)
	api.Put("/starsystem/:id", updateStarSystem)
	api.Delete("/starsystem/:id", deleteStarSystem)
	// Planets
	api.Get("/planet", getPlanets)
	api.Get("/planet/:id", getSinglePlanet)
	api.Post("/planet", createPlanet)
	api.Put("/planet/:id", updatePlanet)
	api.Delete("/planet/:id", deletePlanet)
	// Moons
	api.Get("/moon", getMoons)
	api.Get("/moon/:id", getSingleMoon)
	api.Post("/moon", createMoon)
	api.Put("/moon/:id", updateMoon)
	api.Delete("/moon/:id", deleteMoon)
}
