package requests

// InventoryProductResponse defines the response structure for inventory products.
type CreateInventoryProductRequest struct {
	ProductId int     `json:"product_id"`
	Price     float32 `json:"price"`
}

// InventoryDetailResponse defines the response structure for inventory details.
type InventoryDetailResponse struct {
	Id                 int    `json:"id"`
	InventoryProductId int    `json:"inventory_product_id"`
	Quantity           int    `json:"quantity"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"` // or time.Time depending on your usage
	UpdatedAt          string `json:"updated_at"` // or time.Time depending on your usage
}
