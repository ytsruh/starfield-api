package api

import (
	"github.com/gofiber/fiber/v2"
	database "starfieldapi.com/db"
)

func getMoons(c *fiber.Ctx) error {
	moons, err := database.GetAllMoons()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"data": moons,
	})
}

func getSingleMoon(c *fiber.Ctx) error {
	param := c.Params("id")
	moon, err := database.GetMoon(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"data": moon,
	})
}
