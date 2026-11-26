package artists

import (
	"API_Artiste/src/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RemoveHandler(db *gorm.DB) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil || id == 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID invalide",
			})
		}

		result := db.Delete(&models.Artist{}, id)

		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}

		if result.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{
				"error": "Artiste introuvable",
			})
		}

		return c.SendStatus(204)
	}
}
