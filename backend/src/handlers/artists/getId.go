package artists

import (
	"API_Artiste/src/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetIdHandler(db *gorm.DB) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil || id == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "ID invalide"})
		}
		var artist models.Artist
		result := db.First(&artist, id)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Artiste introuvable"})
		}
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.JSON(artist)
	}
}
