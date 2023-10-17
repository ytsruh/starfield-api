package api

import (
	"github.com/gofiber/fiber/v2"
	database "starfieldapi.com/db"
)

func getPlanets(c *fiber.Ctx) error {
	planets, err := database.GetAllPlanets()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"data": planets,
	})
}

func getSinglePlanet(c *fiber.Ctx) error {
	param := c.Params("id")
	planet, err := database.GetPlanet(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"data": planet,
	})
}
