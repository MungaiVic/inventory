package repository

import (
	"errors"
	"inv-v2/internal/models"

	"gorm.io/gorm"
)

type PgUserRepository struct {
	db *gorm.DB
}

func NewUserConnection(db *gorm.DB) UserRepository {
	return &PgUserRepository{db: db}
}

// Get all users in the system
func (user PgUserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	user.db.Find(&users)
	return users, nil
}

// Get user by ID
func (user PgUserRepository) GetUserByID(userID string) (*models.User, error) {
	var userObj models.User
	user.db.First(&userObj, "user_id = ?", userID)
	if userObj.FirstName == "" {
		return &userObj, errors.New("no user found")
	}
	return &userObj, nil
}

// Get user by email
func (user PgUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var userObj models.User
	user.db.First(&userObj, "Email = ?", email)
	if userObj.FirstName == "" {
		return &userObj, errors.New("no user found")
	}
	return &userObj, nil
}

// Get user by username
func (user PgUserRepository) GetUserByUsername(userName string) (*models.User, error) {
	var userObj models.User
	user.db.First(&userObj, "Username = ?", userName)
	if userObj.FirstName == "" {
		return &userObj, errors.New("no user found")
	}
	return &userObj, nil
}

// Create new user
func (user PgUserRepository) CreateUser(regUser *models.User) (*models.User, error) {
	if err := user.db.Create(&regUser).Error; err != nil {
		return nil, err
	}
	return regUser, nil
}

func (user PgUserRepository) UpdateUser(userupdate *models.User) (*models.User, error) {
	user.db.Model(&userupdate).Updates(models.User{
		Username:  userupdate.Username,
		FirstName: userupdate.FirstName,
		LastName:  userupdate.LastName,
		Email:     userupdate.Email})
	return userupdate, nil
}


func (user PgUserRepository) ChangePassword(passchange *models.User) (*models.User, error) {
	user.db.Model(&passchange).Updates(models.User{
		Password:  passchange.Password,
})
	return passchange, nil
}
