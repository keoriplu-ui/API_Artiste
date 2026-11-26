package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v3"

	"API_Artiste/src/database"
	"API_Artiste/src/handlers"
	"API_Artiste/src/models"
)

func main() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(&models.Artist{})
	if err != nil {
		log.Fatal("Erreur migration :", err)
	}

	app := fiber.New()

	app.Post("/artists", handlers.Post)
	app.Get("/artists", handlers.GetAll)
	app.Get("/artists/:id", handlers.GetId)
	app.Put("/artists/:id", handlers.Put)
	app.Delete("/artists/:id", handlers.Remove)

	log.Fatal(app.Listen(":3500"))
}








// DELETE

