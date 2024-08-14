package entity

import "time"

type InventoryDetail struct {
	Id                 int
	InventoryProductId int
	Stock              int
	Status             string
	Version            int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
