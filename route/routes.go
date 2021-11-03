package route

import (
	"github.com/gofiber/fiber/v2"
	"mocklogin/controller"
)

// Setup auth routes
func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.RetrieveUser)
}