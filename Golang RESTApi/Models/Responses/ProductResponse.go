package responses

import "time"

type ProductRespon struct {
	Id          int64     `json:"id"`
	ProductName string    `json:"product_name"`
	ProductDesc string    `json:"product_desc"`
	CreateBy    string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductDetailRespon struct {
	Id          int64     `json:"id"`
	ProductName string    `json:"product_name"`
	ProductDesc string    `json:"product_desc"`
	Price       float32   `json:"price"`
	Stock       int       `json:"Stock"`
	Status      string    `json:"status"`
	CreateBy    string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
