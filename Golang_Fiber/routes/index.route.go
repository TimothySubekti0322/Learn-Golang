package routes

import (
	"belajar-golang-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

// User Routes
userGroup := r.Group("/users")

userGroup.Get("/", controllers.GetAllUsers)
userGroup.Post("/", controllers.CreateUsers)
userGroup.Get("/:id", controllers.GetUserById)
userGroup.Patch("/:id", controllers.UpdateUser)
userGroup.Delete("/:id", controllers.DeleteUser)



}