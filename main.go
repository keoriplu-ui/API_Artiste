package main

import (
	"log"

	"API_Artiste/src/database"
	"API_Artiste/src/models"
	"API_Artiste/src/routes"
)

func main() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(&models.Artist{})
	if err != nil {
		log.Fatal("Erreur migration :", err)
	}
	err = database.DB.AutoMigrate(&models.Concert{})
	if err != nil {
		log.Fatal("Erreur migration :", err)
	}

	app := routes.AllRoutes()

	log.Fatal(app.Listen(":3500"))
}

// DELETE
