package controllers

import (
	"github.com/MungaiVic/inventory/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
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

func CreateItem(context *fiber.Ctx, db *gorm.DB) error {
	itemModel := &models.Item{}
	validate := validator.New()
	err := context.BodyParser(itemModel)
	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Request failed.",
		})
		return err
	}
	// Running validations
	if err := validate.Struct(itemModel); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err.Error())
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

func UpdateItem(context *fiber.Ctx, db *gorm.DB) error {
	var itemObj models.Item
	var item models.Item
	validate := validator.New()
	err := context.BodyParser(&itemObj)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "malformed request",
			"data":    err,
		})
	}
	// Running validations
	if err := validate.Struct(itemObj); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	db.First(&item, itemObj.ID)
	item.Name = itemObj.Name
	item.Price = itemObj.Price
	item.Quantity = itemObj.Quantity
	item.Reorderlvl = itemObj.Reorderlvl
	db.Save(&item)

	return context.Status(fiber.StatusAccepted).JSON(fiber.Map{
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
