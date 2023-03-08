package repository

import (
	"github.com/MungaiVic/inventory/pkg/models"
	"gorm.io/gorm"
)

// Data Access Object == DAO

type DAO struct{
	db *gorm.DB
}

func New(db *gorm.DB) ItemRepository{
	return &DAO{db}
}

func (dao *DAO) GetAll() ([]models.Item, error){
	var items []models.Item
	dao.db.Find(&items)
	return items, nil
}
