package repository

import (
	"errors"

	"github.com/MungaiVic/inventory/pkg/models"
	"gorm.io/gorm"
)

type DAO struct {
	db *gorm.DB
}

func New(db *gorm.DB) ItemRepository{
	return &DAO{db}
}

func (dao *DAO) GetAll() []models.Item {
	var items []models.Item
	dao.db.Find(&items)
	return items
}

func (dao *DAO) GetByID(itemID uint64) (models.Item, error) {
	itemModel := &models.Item{}
	if itemID == 0{
		return models.Item{},errors.New("ID cannot be empty")
	}
	if err := dao.db.Where("id = ?", itemID).First(itemModel).Error; err != nil{
		return *itemModel, nil
	}
	return models.Item{}, errors.New("item not found")
}

func (dao *DAO) Update(itemID int, item models.Item) {
	panic("Implement Update")
}

func (dao *DAO) Create(models.Item) {
	panic("Implement Create")
}

func (dao *DAO) Delete() {
	panic("Implement Delete")
}
