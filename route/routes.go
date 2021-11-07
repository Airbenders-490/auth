package route

import (
	"github.com/airbenders/auth/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup auth routes
func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/validate", controller.ValidateLogin)
}
