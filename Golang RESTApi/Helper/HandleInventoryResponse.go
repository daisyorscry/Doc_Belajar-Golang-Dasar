package helper

import (
	entity "RESTApi/Models/Entity"
	responses "RESTApi/Models/Responses"
)

func HandleProductInventory(product entity.InventoryProduct) responses.InventoryProductResponse {
	return responses.InventoryProductResponse{
		Id:        product.Id,
		ProductId: product.ProductId,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func HandleProductInventories(products []entity.InventoryProduct) []responses.InventoryProductResponse {
	var inventoryResponses []responses.InventoryProductResponse

	for _, inventory := range products {
		inventoryResponse := responses.InventoryProductResponse{
			Id:        inventory.Id,
			ProductId: inventory.ProductId,
			Price:     inventory.Price,
			CreatedAt: inventory.CreatedAt,
			UpdatedAt: inventory.UpdatedAt,
		}
		inventoryResponses = append(inventoryResponses, inventoryResponse)
	}
	return inventoryResponses
}
