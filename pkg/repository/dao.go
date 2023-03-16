/*
Package repository defines the methods that database access methods need to implement.
*/

package repository

import "github.com/MungaiVic/inventory/pkg/models"

type ItemRepository interface {
	GetAll() []models.Item
	GetByID(itemID uint64) (models.Item, error)
	// Create(models.Item)
	// Update(itemID int, item models.Item)
	// Delete()
}
