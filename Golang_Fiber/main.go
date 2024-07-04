package main

import (
	"belajar-golang-fiber/database"
	"belajar-golang-fiber/database/migrations"
	"belajar-golang-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
 app := fiber.New()


 // Inisialisasi Database
 database.DatabaseInit()

 // Inisialisasi Migration
 migrations.Migration()

 app.Get("/", func(c *fiber.Ctx) error {
  return c.Status(200).JSON(fiber.Map{
   "message": "Hello",
  })
 })

 // Inisialisasi Rute
 routes.RouteInit(app)

 app.Listen(":8080")
}