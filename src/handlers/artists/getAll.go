package artists

import (
	"API_Artiste/src/database"
	"API_Artiste/src/models"

	"github.com/gofiber/fiber/v3"
)

func GetAll(c fiber.Ctx) error {
	var artists []models.Artist
	result := database.DB.Find(&artists)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(artists)
}
