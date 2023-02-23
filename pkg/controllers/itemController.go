package controllers

import (
	"github.com/MungaiVic/inventory/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetItems(c *fiber.Ctx, db *gorm.DB) error {
	var items []models.Item

	db.Find(&items)
	return c.JSON(items)
}

func GetItem(context *fiber.Ctx, db *gorm.DB) error {
	itemModel := &models.Item{}
	itemID := context.Params("id")

	if err := db.Where("id = ?", itemID).First(itemModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"message": "Item not found.",
			})
			return nil
		}
		context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not get the item.",
		})
		return err
	}

	context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Book id fetched successfully",
		"data":    itemModel,
	})
	return nil
}

func CreateItem(context *fiber.Ctx, db *gorm.DB) error {
	itemModel := &models.Item{}
	err := context.BodyParser(itemModel)
	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Request failed.",
		})
		return err
	}
	err = db.Create(&itemModel).Error
	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not create book",
		})
		return err
	}
	context.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "Created successfully",
		"data":    itemModel,
	})
	return nil
}
