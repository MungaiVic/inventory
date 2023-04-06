package repository

import (
	"inv-v2/internal/models"
)

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
	GetItem(id string) (models.Item, error)
	CreateItem(item *models.Item) (*models.Item, error)
	UpdateItem(item *models.Item) (*models.Item, error)
	DeleteItem(id string) error
}

type UserRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(userName string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(userupdate *models.User) (*models.User, error)
	ChangePassword(passchange *models.User) (*models.User, error)
	DeleteUser(userID string) error
}
