package helper

import (
	entity "RESTApi/Models/Entity"
	responses "RESTApi/Models/Responses"
)

func HandleProductResponse(product entity.Product) responses.ProductRespon {
	return responses.ProductRespon{
		Id:          product.Id,
		ProductName: product.ProductName,
		ProductDesc: product.ProductDesc,
		CreateBy:    product.CreateBy,
		CreatedAt:   product.CreatedAt,
	}
}

func HandleProductDetailResponse(product entity.Product, inventoryDetail entity.InventoryDetail) responses.ProductDetailRespon {
	return responses.ProductDetailRespon{
		Id:          product.Id,
		ProductName: product.ProductName,
		ProductDesc: product.ProductDesc,
		Stock:       inventoryDetail.Stock,
		Status:      inventoryDetail.Status,
		CreateBy:    product.CreateBy,
		CreatedAt:   product.CreatedAt,
	}
}

func HandleProductResponses(products []entity.Product) []responses.ProductRespon {
	var productResponses []responses.ProductRespon

	for _, product := range products {
		productResponse := responses.ProductRespon{
			Id:          product.Id,
			ProductName: product.ProductName,
			ProductDesc: product.ProductDesc,
			CreateBy:    product.CreateBy,
			CreatedAt:   product.CreatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses
}
