package service

import (
	"inv-v2/internal/models"
)

func ValidateItem(item *models.Item) map[string]string {
	itemErrors := make(map[string]string)
	if item.Name == "" {
		itemErrors["name"] = "name must not be empty"
	} else if len(item.Name) < 3 {
		itemErrors["name"] = "name must not be less than 3 characters"
	}

	if item.Price <= 0 {
		itemErrors["price"] = "price cannot be a zero value or less"
	}

	if item.Quantity <= 0 {
		itemErrors["quantity"] = "quantity cannot be a zero value or less"
	}

	if item.Reorderlvl <= 0 {
		itemErrors["reorderlvl"] = "reorderlvl cannot be a zero value or less"
	}

	return itemErrors
}
