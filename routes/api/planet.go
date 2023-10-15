package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func createPlanet(c *fiber.Ctx) error {
	input := new(database.Planet)

	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{
			"message": "failed to parse request body",
		})
	}
	inputData := database.Planet{
		Name:         input.Name,
		StarSystemId: input.StarSystemId,
	}
	error := database.CreatePlanet(&inputData)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create new planet" + error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully created new star system",
		"data":    inputData,
	})
}

func updatePlanet(c *fiber.Ctx) error {
	input := new(database.Planet)
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

	err := database.UpdatePlanet(input)
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

func deletePlanet(c *fiber.Ctx) error {
	param := c.Params("id")
	err := database.DeletePlanet(param)
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
