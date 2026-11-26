package artists

import (
	"API_Artiste/src/database"
	"API_Artiste/src/models"

	"github.com/gofiber/fiber/v3"
)

func Post(c fiber.Ctx) error {
	var artist models.Artist
	if err := c.Bind().Body(&artist); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Données invalides"})
	}

	result := database.DB.Create(&artist)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(201).JSON(artist)
}
