package requests

type StockChangeRequest struct {
	ProductId int    `json:"product_id"`
	Change    int    `json:"change"`
	Status    string `json:"status"`
}
