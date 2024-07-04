package migrations

import (
	"belajar-golang-fiber/database"
	"belajar-golang-fiber/models"
	"fmt"
	"log"
)

func Migration() {
 err := database.DB.AutoMigrate(
  &models.User{},

 )
 if err != nil {
  log.Fatal("Failed to migrate...")
 }

 fmt.Println("Migrated successfully")
}