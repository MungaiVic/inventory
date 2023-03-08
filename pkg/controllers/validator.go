package controllers

import (
	"errors"

	"github.com/MungaiVic/inventory/pkg/models"
)

func ValidateItem(item models.Item) map[string]error {
	itemErrors := make(map[string]error)
	if item.Name == "" {
		itemErrors["name"] = errors.New("name must not be empty")
	} else if len(item.Name) < 3 {
		itemErrors["name"] = errors.New("name must not be less than 3 characters")
	}

	if item.Price <= 0 {
		itemErrors["price"] = errors.New("price cannot be a zero value or less")
	}

	if item.Quantity <= 0 {
		itemErrors["quantity"] = errors.New("quantity cannot be a zero value or less")
	}

	if item.Reorderlvl <= 0 {
		itemErrors["reorderlvl"] = errors.New("reorderlvl cannot be a zero value or less")
	}


	return itemErrors
}
