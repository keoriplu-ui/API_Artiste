
package database

import (
    "log"

    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("artists.db"), &gorm.Config{})

    if err != nil {
        log.Fatal("Erreur de connexion à SQLite :", err)
    }

    log.Println("Connexion à SQLite réussie !")
    return db
}
