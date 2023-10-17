package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	database "starfieldapi.com/db"
)

func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	// API Auth Middleware
	api.Use(func(c *fiber.Ctx) error {
		// Get all keys from database
		keys, err := database.GetAllKeys()
		if err != nil {
			return c.SendStatus(500)
		}
		// Find the API Key in the request
		requestKey, reqErr := getKeyFromRequest(c)
		if reqErr != nil {
			c.Status(401)
			return c.JSON(fiber.Map{
				"success": false,
				"message": reqErr,
			})
		}
		// Check if the API key is in the list of active keys from DB
		authed := slices.Contains(keys, requestKey)
		if authed {
			go createRequestLog(c, requestKey)
			return c.Next()
		}
		c.Status(401)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "invalid api key",
		})
	})

	// Star Systems
	api.Get("/starsystem", getStarSystems)
	api.Get("/starsystem/:id", getSingleStarSystem)
	// Planets
	api.Get("/planet", getPlanets)
	api.Get("/planet/:id", getSinglePlanet)
	// Moons
	api.Get("/moon", getMoons)
	api.Get("/moon/:id", getSingleMoon)
}

func createRequestLog(ctx *fiber.Ctx, key string) {
	path := ctx.Path()
	fmt.Println("Path:" + path)
	fmt.Println("Key:" + key)
}

func getKeyFromRequest(ctx *fiber.Ctx) (string, error) {
	headers := ctx.GetReqHeaders()
	headerToken := headers["Api-Key"]
	if len(headerToken) > 0 {
		return headerToken, nil
	}
	queryToken := ctx.Query("apikey")
	if len(queryToken) > 0 {
		return queryToken, nil
	}
	return "", errors.New("no api key found")
}
