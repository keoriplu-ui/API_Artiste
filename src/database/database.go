
package database

import (
    "log"

    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    db, err := gorm.Open(sqlite.Open("artists.db"), &gorm.Config{})

    if err != nil {
        log.Fatal("Erreur de connexion à SQLite :", err)
    }

    DB = db

    log.Println("Connexion à SQLite réussie !")
}
