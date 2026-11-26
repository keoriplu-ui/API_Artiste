package routes

import (
	"API_Artiste/src/handlers/artists"
	"API_Artiste/src/handlers/concerts"
	"API_Artiste/src/handlers/menu"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"gorm.io/gorm"
)

func AllRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	app.Get("/", menu.Menu)

	app.Post("/artists", artists.PostHandler(db))
	app.Get("/artists", artists.GetAllHandler(db))
	app.Get("/artists/:id", artists.GetIdHandler(db))
	app.Put("/artists/:id", artists.PutHandler(db))
	app.Delete("/artists/:id", artists.RemoveHandler(db))

	app.Post("/concerts", concerts.PostHandler(db))
	app.Get("/concerts", concerts.GetAllHandler(db))
	app.Get("/concerts/:id", concerts.GetIdHandler(db))
	app.Put("/concerts/:id", concerts.PutHandler(db))
	app.Delete("/concerts/:id", concerts.RemoveHandler(db))
	return app
}
