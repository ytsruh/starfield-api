package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	database "starfieldapi.com/db"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the Starfield API",
		})
	})
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Live Server Metrics"}))
	app.Get("/starsystem", getStarSystems)
	app.Get("/starsystem/:id", getSingleStarSystem)
	app.Post("/starsystem", createStarSystem)
}

func getStarSystems(c *fiber.Ctx) error {
	systems, err := database.GetAllStarSystems()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"data": systems,
	})
}

func getSingleStarSystem(c *fiber.Ctx) error {
	param := c.Params("id")
	system, err := database.GetStarSystem(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"data": system,
	})
}

type Input struct {
	Name string
}

func createStarSystem(c *fiber.Ctx) error {
	input := new(Input)
	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{
			"message": "failed to parse request body",
		})
	}
	inputData := database.StarSystem{
		Name: input.Name,
	}
	error := database.CreateStarSystem(&inputData)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create new url" + error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully created new url",
		"data":    inputData,
	})
}
