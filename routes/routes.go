package routes

import (
	"github.com/VJSRE/go-backend-starter/handlers"
	"github.com/VJSRE/go-backend-starter/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	var publicRoutes = app.Group("/api/v1")

	publicRoutes.Post("/signup", handlers.Signup)
	publicRoutes.Post("/login", handlers.Login)
	publicRoutes.Get("/items", handlers.GetAllItems)
	publicRoutes.Get("/items/:id", handlers.GetItemById)

	var privateRoutes = app.Group("/api/v1", middlewares.CreateMiddleware())
	privateRoutes.Post("/items", handlers.CreateItem)
	privateRoutes.Put("/items:id", handlers.UpdateItem)
	privateRoutes.Delete("/items:id", handlers.DeleteItem)

}
