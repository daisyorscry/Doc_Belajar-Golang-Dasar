package entity

import "time"

type InventoryProduct struct {
	Id        int
	ProductId int
	Price     float32
	CreatedAt time.Time
	UpdatedAt time.Time
}
