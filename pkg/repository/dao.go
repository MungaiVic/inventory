package repository

import "github.com/MungaiVic/inventory/pkg/models"

type ItemRepository interface {
	GetAll() ([]models.Item, error)
	// FindByID(id uint) (*models.Item, error)
	// Create(item *models.Item) error
	// Update(item *models.Item) error
	// Delete(item *models.Item) error
}
