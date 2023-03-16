package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name,omitempty"`
	Price      uint16 `json:"price,omitempty"`
	Quantity   uint16 `json:"quantity,omitempty"`
	Reorderlvl uint16 `json:"reorderlvl,omitempty"`
}

func MigrateItems(db gorm.DB) error {
	err := db.AutoMigrate(&Item{})
	if err == nil {
		fmt.Println("Item model successfully migrated!")
	}
	return err
}
