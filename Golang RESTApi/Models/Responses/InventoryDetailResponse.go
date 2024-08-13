package responses

import "time"

type InventoryDetailResponse struct {
	Id                 int       `json:"id"`
	InventoryProductId int       `json:"inventory_product_id"`
	Stock              int       `json:"stock"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
