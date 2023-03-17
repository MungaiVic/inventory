package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4(); primaryKey; not null"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
}

func MigrateUsers(db gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err == nil {
		fmt.Println("User model successfully migrated!")
	}
	return err
}
