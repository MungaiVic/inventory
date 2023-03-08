package svc

import "github.com/MungaiVic/inventory/pkg/models"

func ConvertItemModelToItemService(items []models.Item) []ItemResponse{
	var itemResponses []ItemResponse
	for _, item := range items {
		itemResponse := ItemResponse{
			ID: item.ID,
			Name: item.Name,
			Price: item.Price,
			Reorderlvl: item.Reorderlvl,
			Quantity: item.Quantity,
		}
		itemResponses = append(itemResponses, itemResponse)
	}
	return itemResponses
}
