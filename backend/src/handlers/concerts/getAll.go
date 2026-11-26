package concerts

import (
	"API_Artiste/src/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetAllHandler(db *gorm.DB) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var concerts []models.Concert
		result := db.Find(&concerts)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(concerts)
	}
}
