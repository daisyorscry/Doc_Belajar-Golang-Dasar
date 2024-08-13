package entity

import "time"

type InventoryDetail struct {
	Id                 int
	InventoryProductId int
	Stock              int
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
