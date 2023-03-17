package service

import (
	"fmt"
	"inv-v2/internal/models"
	"inv-v2/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type ItemImpl struct {
	db repository.ItemRepository
}

func NewItemService(db repository.ItemRepository) ItemService {
	return &ItemImpl{db: db}
}

func (item *ItemImpl) GetItems(c *fiber.Ctx) error {
	items, _ := item.db.GetAllItems()
	return c.Status(fiber.StatusOK).JSON(items)
}

func (item *ItemImpl) GetItem(c *fiber.Ctx) error {
	itemID := c.Params("id")
	itemObj, _ := item.db.GetItem(itemID)
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "itemID cannot be empty"})
	}
	if itemObj.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "item not found"})
	}
	return c.Status(fiber.StatusOK).JSON(itemObj)
}

func (item *ItemImpl) CreateItem(c *fiber.Ctx) error {
	itemModel := &models.Item{}
	err := c.BodyParser(itemModel)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"data":    err,
		})
	}
	if err := ValidateItem(itemModel); len(err) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"data":    err,
		})
	}
	itemObj, _ := item.db.CreateItem(itemModel)
	return c.Status(fiber.StatusCreated).JSON(itemObj)
}

func (item *ItemImpl) UpdateItem(c *fiber.Ctx) error {
	itemModel := &models.Item{}
	err := c.BodyParser(&itemModel)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"data":    err,
		})
	}
	// validate item
	if err := ValidateItem(itemModel); len(err) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"data":    err,
		})
	}
	itemObj, _ := item.db.UpdateItem(itemModel)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "item updated successfully",
		"data":    itemObj,
	})
}

func (item *ItemImpl) DeleteItem(c *fiber.Ctx) error {
	itemID := c.Params("id")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "itemID cannot be empty"})
	}
	err := item.db.DeleteItem(itemID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "item not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "item deleted successfully"})
}
