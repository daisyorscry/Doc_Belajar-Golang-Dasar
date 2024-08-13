package responses

import "time"

type InventoryProductResponse struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"created_by"`
}
