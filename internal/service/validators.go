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

func ValidateRegisterUser(user *UserRegistration) map[string]string {
	userErrors := make(map[string]string)

	if user.Username == "" {
		userErrors["username"] = "username must not be empty"
	} else if len(user.Username) < 3 {
		userErrors["username"] = "username must not be less than 3 characters"
	}
	if len(user.Password) < 6 {
		userErrors["password "] = "password must not be less than 6 characters"
	}
	if len(user.FirstName) < 3 {
		userErrors["first_name"] = "first_name must not be less than 3 characters"
	}
	if len(user.LastName) < 3 {
		userErrors["last_name"] = "last_name must not be less than 3 characters"
	}
	if user.FirstName == "" {
		userErrors["first_name"] = "first_name must not be empty"
	}
	if user.LastName == "" {
		userErrors["last_name"] = "last_name must not be empty"
	}
	return userErrors
}
