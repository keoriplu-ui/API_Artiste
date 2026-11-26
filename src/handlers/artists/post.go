package artists

import (
	"API_Artiste/src/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func PostHandler(db *gorm.DB) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var artist models.Artist
		if err := c.Bind().Body(&artist); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Données invalides"})
		}

		result := db.Create(&artist)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}

		return c.Status(201).JSON(artist)
	}
}
