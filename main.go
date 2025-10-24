package main

import (
	"github.com/VJSRE/go-backend-starter/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	var app *fiber.App = fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")

}
