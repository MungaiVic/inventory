package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name" validate:"required,min=3"`
	Price      string `json:"price" validate:"required"`
	Quantity   string `json:"quantity" validate:"required"`
	Reorderlvl string `json:"reorderlvl" validate:"required"`
}

func MigrateItems(db *gorm.DB) error {
	fmt.Println("Migrating Item model...")
	err := db.AutoMigrate(&Item{})
	if err == nil {
		fmt.Println("Item model successfully migrated!")
	}
	return err
}
