package main

import (
	"github.com/VJSRE/go-backend-starter/database"
	"github.com/VJSRE/go-backend-starter/routes"
	"github.com/VJSRE/go-backend-starter/utils"
	"github.com/gofiber/fiber/v2"

	"fmt"
)

var DEFAULT_PORT = "3000"

func NewFiberApp() *fiber.App {
	var app *fiber.App = fiber.New()
	routes.SetupRoutes(app)
	return app
}

func main() {

	var app *fiber.App = NewFiberApp()

	database.InitDatabase(utils.GetValue("DB_NAME"))

	var PORT string = utils.GetValue("PORT")
	if PORT == "" {
		PORT = DEFAULT_PORT
	}

	app.Listen(fmt.Sprintf(":%s", PORT))

}
