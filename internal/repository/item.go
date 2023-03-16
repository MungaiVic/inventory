package repository

import (
	"gorm.io/gorm"
	"inv-v2/internal/models"
	"strconv"
)

type PgItemRepo struct {
	db *gorm.DB
}

func NewItemConnection(db *gorm.DB) ItemRepository {
	return &PgItemRepo{db: db}
}

func (item PgItemRepo) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	item.db.Find(&items)

	return items, nil
}

func (item PgItemRepo) GetItem(id string) (models.Item, error) {
	var itemObj models.Item
	if err := item.db.Where("id = ?", id).First(&itemObj).Error; err != nil {
		return models.Item{}, err
	}

	return itemObj, nil
}

func (item PgItemRepo) CreateItem(itemObj *models.Item) (*models.Item, error) {
	if err := item.db.Create(&itemObj).Error; err != nil {
		return nil, err
	}

	return itemObj, nil
}

func (item PgItemRepo) UpdateItem(itemObj *models.Item) (*models.Item, error) {
	var itemValue models.Item
	item.db.First(&itemValue, strconv.Itoa(int(itemObj.ID)))
	if itemValue.ID == 0 {
		return nil, nil
	}
	itemValue.Name = itemObj.Name
	itemValue.Quantity = itemObj.Quantity
	itemValue.Price = itemObj.Price
	itemValue.Reorderlvl = itemObj.Reorderlvl
	item.db.Save(&itemValue)
	return &itemValue, nil

}

func (item PgItemRepo) DeleteItem(id string) error {
	var itemObj models.Item
	if err := item.db.Where("id = ?", id).First(&itemObj).Error; err != nil {
		return err
	}
	item.db.Where("id = ?", id).Delete(&itemObj)
	return nil
}
