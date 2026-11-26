package main

import (
	"log"

	"API_Artiste/src/database"
	"API_Artiste/src/models"
	"API_Artiste/src/routes"
)

func main() {
	db := database.ConnectDatabase()

	err := db.AutoMigrate(&models.Artist{})
	if err != nil {
		log.Fatal("Erreur migration :", err)
	}
	err = db.AutoMigrate(&models.Concert{})
	if err != nil {
		log.Fatal("Erreur migration :", err)
	}

	app := routes.AllRoutes(db)

	log.Fatal(app.Listen(":3500"))
}

// DELETE
