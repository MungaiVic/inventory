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
		"message": "Item fetched successfully",
		"data":    itemModel,
	})
	return nil
}

func CreateItem(c *fiber.Ctx, db *gorm.DB) error {
	itemModel := &models.Item{}
	err := c.BodyParser(itemModel)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Request failed.",
		})
		return err
	}
	// Running validations
	if err := ValidateItem(*itemModel); len(err) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request body",
			"errors":  err,
		})
	}

	err = db.Create(&itemModel).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not create book",
		})
		return err
	}
	c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "Created successfully",
		"data":    itemModel,
	})
	return nil
}

func UpdateItem(c *fiber.Ctx, db *gorm.DB) error {
	var itemObj models.Item
	var item models.Item
	err := c.BodyParser(&itemObj)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "malformed request",
			"data":    err,
		})
	}
	// Running validations
	if err := ValidateItem(itemObj); len(err) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request body",
			"errors":  err,
		})
	}
	db.First(&item, itemObj.ID)
	item.Name = itemObj.Name
	item.Price = itemObj.Price
	item.Quantity = itemObj.Quantity
	item.Reorderlvl = itemObj.Reorderlvl
	db.Save(&item)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Update Successful",
		"data":    item,
	})
}

func DeleteItem(context *fiber.Ctx, db *gorm.DB) error {
	itemID := context.Params("id")
	var itemModel models.Item
	// get the item
	db.First(&itemModel, itemID)
	if itemModel.Name == "" {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No item found with supplied ID.",
		})
	}
	db.Delete(&itemModel)
	return context.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Item deleted successfully.",
		"data":    itemModel,
	})
}
