package main

import (
	"github.com/gofiber/fiber/v2"
	"mocklogin/database"
	"mocklogin/route"
)

func main() {

	database.Connect()

	app := fiber.New()

	route.Setup(app)

	app.Listen(":3000")
}

