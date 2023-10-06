package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	jwtware "github.com/gofiber/jwt/v2"
	"starfieldapi.com/lib"
)

// Protected protect routes
func protected() fiber.Handler {
	// Default behaviour is that token is found in 'Authorization' header with the prefix of 'Bearer'
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(lib.GetSecretKey()),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT"})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT"})
}

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the Starfield API",
		})
	})
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Live Server Metrics"}))
	// Auth
	app.Post("/auth/register", registerUser)
	app.Post("/auth/login", loginUser)
	app.Get("/auth/me", getUser)

	// Create route group and add protection
	v1 := app.Group("/v1", protected())
	// Star Systems
	v1.Get("/starsystem", getStarSystems)
	v1.Get("/starsystem/:id", getSingleStarSystem)
	v1.Post("/starsystem", createStarSystem)
	v1.Put("/starsystem/:id", updateStarSystem)
	v1.Delete("/starsystem/:id", deleteStarSystem)
	// Planets
	v1.Get("/planet", getPlanets)
	v1.Get("/planet/:id", getSinglePlanet)
	v1.Post("/planet", createPlanet)
	v1.Put("/planet/:id", updatePlanet)
	v1.Delete("/planet/:id", deletePlanet)
	// Moons
	v1.Get("/moon", getMoons)
	v1.Get("/moon/:id", getSingleMoon)
	v1.Post("/moon", createMoon)
	v1.Put("/moon/:id", updateMoon)
	v1.Delete("/moon/:id", deleteMoon)
}
