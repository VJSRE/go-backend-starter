package routes

import (
	"github.com/VJSRE/go-backend-starter/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/items", handlers.GetAllItems)
	app.Get("/api/v1/items:id", handlers.GetItemById)
	app.Post("/api/v1/items", handlers.CreateItem)
	app.Put("/api/v1/items:id", handlers.UpdateItem)
	app.Delete("/api/v1/items:id", handlers.DeleteItem)
}
