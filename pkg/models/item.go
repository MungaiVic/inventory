package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name" validate:"required,min=3"`
	Price      uint32 `json:"price" validate:"required"`
	Quantity   uint8  `json:"quantity" validate:"required,gte=0"`
	Reorderlvl uint8  `json:"reorderlvl" validate:"required"`
}

func MigrateItems(db *gorm.DB) error {
	fmt.Println("Migrating Item model...")
	err := db.AutoMigrate(&Item{})
	if err == nil {
		fmt.Println("Item model successfully migrated!")
	}
	return err
}
