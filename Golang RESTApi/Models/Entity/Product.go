package entity

import "time"

type Product struct {
	Id                       int64
	ProductName, ProductDesc string
	CreateBy                 string
	UserId                   int
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
