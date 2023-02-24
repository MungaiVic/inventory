package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ValidateDecimal (fl validator.FieldLevel) bool {
	if v, ok := fl.Field().Interface().(decimal.Decimal); ok {
		return v.IsPositive()
	}
	return false
}

type Item struct {
	gorm.Model
	Name       string          `gorm:"not null" json:"name" validate:"required,min=3"`
	Price      decimal.Decimal `json:"price" validate:"required, decimal"`
	Quantity   uint8           `json:"quantity" validate:"required,gte=0"`
	Reorderlvl string          `json:"reorderlvl" validate:"required"`
}

func MigrateItems(db *gorm.DB) error {
	fmt.Println("Migrating Item model...")
	err := db.AutoMigrate(&Item{})
	if err == nil {
		fmt.Println("Item model successfully migrated!")
	}
	return err
}
