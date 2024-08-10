package requests

type CreateProductRequest struct {
	ProductName string `validate:"required,min=1,max=20" json:"productname"`
	ProductDesc string `validate:"required,min=1,max=100" json:"productdesc"`
}
