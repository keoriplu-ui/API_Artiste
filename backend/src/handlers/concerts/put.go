package concerts

import (
	"API_Artiste/src/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func PutHandler(db *gorm.DB) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil || id == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "ID invalide"})
		}
		var concert models.Concert
		result := db.First(&concert, id)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Concert introuvable"})
		}
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		var nouvConcert models.Concert
		if err := c.Bind().Body(&nouvConcert); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Données invalides"})
		}
		result = db.Model(&concert).Updates(map[string]interface{}{
			"date":      nouvConcert.Date,
			"artist_id": nouvConcert.ArtistId,
		})
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(concert)
	}
}
