package menu

import "github.com/gofiber/fiber/v3"

func Menu(c fiber.Ctx) error {
	return c.SendString("Bienvenue sur l'API des Artistes")
}
