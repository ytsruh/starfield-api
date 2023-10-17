package api

import (
	"github.com/gofiber/fiber/v2"
	database "starfieldapi.com/db"
)

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
