package services

import (
	"errors"
	"github.com/VJSRE/go-backend-starter/models"
	"github.com/google/uuid"
	"time"
)

var storage []models.Item = []models.Item{}

func GetAllItems() []models.Item {
	return storage
}

func GetItemByID(id string) (models.Item, error) {

	for _, item := range storage {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Item{}, errors.New("Item not found")
}

func CreateItem(itemRequest models.ItemRequest) models.Item {

	var newItem models.Item = models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quality:   itemRequest.Quality,
		CreatedAt: time.Now(),
	}

	storage = append(storage, newItem)
	return newItem
}

func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {

	for index, item := range storage {

		if item.ID == id {
			item.Name = itemRequest.Name
			item.Price = itemRequest.Price
			item.Quality = itemRequest.Quality
			item.UpdatedAt = time.Now()

			storage[index] = item
			return item, nil
		}

	}
	return models.Item{}, errors.New("item update failed, item not found")
}

func DeleteItem(id string) bool {
	for index, item := range storage {
		if item.ID == id {
			storage = append(storage[:index], storage[index+1:]...)

		}
	}
	return true
}
