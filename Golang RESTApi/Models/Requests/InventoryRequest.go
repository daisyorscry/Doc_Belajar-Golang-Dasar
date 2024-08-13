package requests

type CreateInventoryProductRequest struct {
	ProductId int     `json:"product_id"`
	Price     float32 `json:"price"`
}
