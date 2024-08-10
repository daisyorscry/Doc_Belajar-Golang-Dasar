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
	}
}

func HandleProductResponses(products []entity.Product) []responses.ProductRespon {
	var productResponses []responses.ProductRespon
	for _, product := range products {
		productResponses = append(productResponses, HandleProductResponse(product))
	}
	return productResponses
}
