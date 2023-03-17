package repository

import (
	"inv-v2/internal/models"

	"gorm.io/gorm"
)

type PgUserRepository struct {
	db *gorm.DB
}

func NewUserConnection(db *gorm.DB) UserRepository {
	return &PgUserRepository{db: db}
}

func (user PgUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	user.db.Find(&users)
	return users, nil
}

func (user PgUserRepository) CreateUser(regUser *models.User) (*models.User, error) {
	if err := user.db.Create(&regUser).Error; err != nil {
		return nil, err
	}
	return regUser, nil
}