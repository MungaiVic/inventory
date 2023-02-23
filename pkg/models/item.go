package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string `json:"name"`
	Price      string `json:"price"`
	Quantity   string `json:"quantity"`
	Reorderlvl string `json:"reorderlvl"`
}

func MigrateItems(db *gorm.DB) error {
	err := db.AutoMigrate(&Item{})
	return err
}
