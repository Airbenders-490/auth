package route

import (
	"github.com/gofiber/fiber/v2"
	"mocklogin/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}