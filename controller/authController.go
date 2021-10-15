package controller

import "github.com/gofiber/fiber/v2"

func Register(context *fiber.Ctx) error {

	// extract the http request data
	var data map[string]string // declare var data of type map (key:string, value:string)

	// handle error if any
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	// return http body content
	return context.JSON(data)
}
