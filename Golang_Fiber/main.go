package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
		ReadTimeout: time.Second * 10,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
	
	err := app.Listen("localhost:3000")


	if err != nil {
		panic(err)
	}
}