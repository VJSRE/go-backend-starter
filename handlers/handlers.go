package handlers

import (
	"github.com/VJSRE/go-backend-starter/models"
	"github.com/VJSRE/go-backend-starter/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item = services.GetAllItems()

	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Data:    items,
		Message: "All Items data",
	})
}

func GetItemById(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	item, err := services.GetItemByID(itemID)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[models.Item]{
		Success: true,
		Data:    item,
		Message: "Item found",
	})
}

func CreateItem(c *fiber.Ctx) error {

	var itemInput *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := itemInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Validation errors",
			Data:    errors,
		})
	}

	var createdItem models.Item = services.CreateItem(*itemInput)

	return c.Status(http.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "Item created",
		Data:    createdItem,
	})
}

func UpdateItem(c *fiber.Ctx) error {
	var inputItem *models.ItemRequest = new(models.ItemRequest)
	if err := c.BodyParser(&inputItem); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := inputItem.ValidateStruct()
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "Validation errors",
		})
	}

	var itemID string = c.Params("id")
	item, err := services.UpdateItem(*inputItem, itemID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response[models.Item]{
		Success: true,
		Data:    item,
		Message: "Item updated",
	})

}

func DeleteItem(c *fiber.Ctx) error {
	var itemID string = c.Params("id")
	status := services.DeleteItem(itemID)
	if status == false {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: "failed to delete item",
		})
	}
	return c.JSON(models.Response[any]{
		Success: true,
		Message: "Item deleted",
	})
}
