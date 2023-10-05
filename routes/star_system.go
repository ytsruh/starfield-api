package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func createStarSystem(c *fiber.Ctx) error {
	input := new(database.StarSystem)

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
			"message": "could not create new star system" + error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully created new star system",
		"data":    inputData,
	})
}

func updateStarSystem(c *fiber.Ctx) error {
	input := new(database.StarSystem)
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

	err := database.UpdateStarSystem(input)
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

func deleteStarSystem(c *fiber.Ctx) error {
	param := c.Params("id")
	err := database.DeleteStarSystem(param)
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
