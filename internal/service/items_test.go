package service

import (
	"inv-v2/internal/repository"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestItemImpl_CreateItem(t *testing.T) {
	type fields struct {
		db repository.ItemRepository
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
			item := &ItemImpl{
				db: tt.fields.db,
			}
			if err := item.CreateItem(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ItemImpl.CreateItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItemImpl_GetItem(t *testing.T) {
	type fields struct {
		db repository.ItemRepository
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
			item := &ItemImpl{
				db: tt.fields.db,
			}
			if err := item.GetItem(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ItemImpl.GetItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItemImpl_UpdateItem(t *testing.T) {
	type fields struct {
		db repository.ItemRepository
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
			item := &ItemImpl{
				db: tt.fields.db,
			}
			if err := item.UpdateItem(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ItemImpl.UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItemImpl_DeleteItem(t *testing.T) {
	type fields struct {
		db repository.ItemRepository
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
			item := &ItemImpl{
				db: tt.fields.db,
			}
			if err := item.DeleteItem(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ItemImpl.DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
