package routes

import (
	"API_Artiste/src/handlers/artists"
	"API_Artiste/src/handlers/concerts"
	"API_Artiste/src/handlers/menu"

	"github.com/gofiber/fiber/v3"
)

func AllRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/", menu.Menu)

	app.Post("/artists", artists.Post)
	app.Get("/artists", artists.GetAll)
	app.Get("/artists/:id", artists.GetId)
	app.Put("/artists/:id", artists.Put)
	app.Delete("/artists/:id", artists.Remove)

	app.Post("/concerts", concerts.Post)
	app.Get("/concerts", concerts.GetAll)
	app.Get("/concerts/:id", concerts.GetId)
	app.Put("/concerts/:id", concerts.Put)
	app.Delete("/concerts/:id", concerts.Remove)
	return app
}
