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
		// Render index template
		return c.Render("index", fiber.Map{
			"PageTitle": "Welcome to the Stafield API",
		})
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("login", fiber.Map{
			"PageTitle": "Login to the Stafield API",
		})
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Live Server Metrics"}))
	// Create route group and add protection
	api := app.Group("/api")
	// Auth
	api.Post("/auth/register", registerUser)
	api.Post("/auth/login", loginUser)
	api.Get("/auth/me", getUser)

	// Create route group and add protection
	v1 := api.Group("/v1", protected())
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
