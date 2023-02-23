package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name"`
	Price      string `json:"price"`
	Quantity   string `json:"quantity"`
	Reorderlvl string `json:"reorderlvl"`
}

func MigrateItems(db *gorm.DB) error {
	fmt.Println("Migrating Item model...")
	err := db.AutoMigrate(&Item{})
	if err == nil {
		fmt.Println("Item model successfully migrated!")
	}
	return err
}
