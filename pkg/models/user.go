package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mgutz/ansi"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4(); primaryKey; not null"`
	FirstName string    `json:"first_name" gorm:"not null" check:"length(first_name) >= 3; size:255" validate:"required,gte=3"`
	LastName  string    `json:"last_name" gorm:"not null" check:"length(last_name) >= 3;size:255" validate:"required,gte=3"`
	Email     string    `json:"email" validate:"required, email"`
	Password  string    `json:"password" gorm:"not null"`
}

func MigrateUsers(db *gorm.DB) error {
	orangefy := ansi.ColorFunc("yellow")
	msg := orangefy("Migrating User model...")
	fmt.Println(msg)
	err := db.AutoMigrate(&User{})
	if err == nil {
		greenify := ansi.ColorFunc("green+b")
		msg := greenify("User model successfully migrated!")
		fmt.Println(msg)
	}
	return err
}
