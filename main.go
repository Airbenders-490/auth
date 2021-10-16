package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"mocklogin/database"
	"mocklogin/route"
)

func main() {

	database.Connect()

	app := fiber.New()

	// when backend & frontend run on different ports, the browser will
	// throw an error and reject our request
	// using cors will make browser accept our request
	app.Use(cors.New(cors.Config{
		AllowCredentials: true, // important for sending/receiving httpOnly cookies
	}))

	route.Setup(app)

	app.Listen(":3000")
}

