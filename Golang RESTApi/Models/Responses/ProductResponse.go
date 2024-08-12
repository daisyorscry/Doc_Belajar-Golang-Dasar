package responses

import "time"

type ProductRespon struct {
	Id          int64     `json:"id"`
	ProductName string    `json:"product_name"`
	ProductDesc string    `json:"product_desc"`
	CreateBy    string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
