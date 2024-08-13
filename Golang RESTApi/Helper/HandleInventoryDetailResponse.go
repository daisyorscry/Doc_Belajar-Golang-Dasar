package helper

import (
	entity "RESTApi/Models/Entity"
	responses "RESTApi/Models/Responses"
)

func HandleInventoryDetail(detail entity.InventoryDetail) responses.InventoryDetailResponse {
	return responses.InventoryDetailResponse{
		Id:                 detail.Id,
		InventoryProductId: detail.InventoryProductId,
		Stock:              detail.Stock,
		Status:             detail.Status,
		CreatedAt:          detail.CreatedAt,
		UpdatedAt:          detail.UpdatedAt,
	}
}
