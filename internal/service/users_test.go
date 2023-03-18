package service

import (
	"inv-v2/internal/repository"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestUserImpl_GetUsers(t *testing.T) {
	type fields struct {
		db repository.UserRepository
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := UserImpl{
				db: tt.fields.db,
			}
			if err := user.GetUsers(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserImpl.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserImpl_Register(t *testing.T) {
	type fields struct {
		db repository.UserRepository
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := UserImpl{
				db: tt.fields.db,
			}
			if err := user.Register(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserImpl.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
