package models

import (
	"fmt"

	"github.com/mgutz/ansi"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name" validate:"required,min=3"`
	Price      uint32 `json:"price" validate:"required"`
	Quantity   uint8  `json:"quantity" validate:"required,gte=0"`
	Reorderlvl uint8  `json:"reorderlvl" validate:"required"`
}

type ItemResponse struct {
	// gorm.Model
	ID         uint
	Name       string `gorm:"not null" json:"name" validate:"required,min=3"`
	Price      uint32 `json:"price" validate:"required"`
	Quantity   uint8  `json:"quantity" validate:"required,gte=0"`
	Reorderlvl uint8  `json:"reorderlvl" validate:"required"`
}

func MigrateItems(db *gorm.DB) error {
	orangefy := ansi.ColorFunc("yellow")
	msg := orangefy("Migrating Item model...")
	fmt.Println(msg)
	err := db.AutoMigrate(&Item{})
	if err == nil {
		greenify := ansi.ColorFunc("green+b")
		msg := greenify("Item model successfully migrated!")
		fmt.Println(msg)
	}
	return err
}
