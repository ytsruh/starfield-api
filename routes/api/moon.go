package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func createMoon(c *fiber.Ctx) error {
	input := new(database.Moon)

	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{
			"message": "failed to parse request body",
		})
	}
	inputData := database.Moon{
		Name:     input.Name,
		PlanetId: input.PlanetId,
	}
	error := database.CreateMoon(&inputData)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create new moon" + error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully created new star system",
		"data":    inputData,
	})
}

func updateMoon(c *fiber.Ctx) error {
	input := new(database.Moon)
	id, parseErr := uuid.Parse(c.Params("id"))
	if parseErr != nil {
		return c.JSON(fiber.Map{
			"message": "failed to parse id",
		})
	}

	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{
			"message": "failed to parse request body",
		})
	}
	input.ID = id

	err := database.UpdateMoon(input)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"success": true,
	})
}

func deleteMoon(c *fiber.Ctx) error {
	param := c.Params("id")
	err := database.DeleteMoon(param)
	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"message": "An error occurred",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"success": true,
	})
}
