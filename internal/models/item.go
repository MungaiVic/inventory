package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string `gorm:"not null"`
	Price      uint16
	Quantity   uint16
	Reorderlvl uint16
}

func MigrateItems(db gorm.DB) error {
	err := db.AutoMigrate(&Item{})
	if err == nil {
		fmt.Println("Item model successfully migrated!")
	}
	return err
}
