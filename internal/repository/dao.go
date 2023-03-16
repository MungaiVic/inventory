package repository

import "inv-v2/internal/models"

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
	GetItem(id string) (models.Item, error)
	CreateItem(item *models.Item) (*models.Item, error)
	UpdateItem(item *models.Item) (*models.Item, error)
	DeleteItem(id string) error
}
