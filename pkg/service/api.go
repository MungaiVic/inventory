package service

import "github.com/MungaiVic/inventory/pkg/models"

type ItemService interface {
	GetItems() ([]models.Item, error)
	GetItemByID()
	CreateItem()
	UpdateItem()
	DeleteItem()
}

type UserService interface {
	GetUsers()
	GetUserByID()
	CreateUser()
	UpdateUser()
	DeleteUser()
}
