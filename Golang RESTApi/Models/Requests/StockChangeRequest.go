package requests

type StockChangeRequest struct {
	ProductId int    `json:"product_id"`
	Change    int    `json:"change"` // bisa positif (untuk menambah stok) atau negatif (untuk mengurangi stok)
	Status    string `json:"status"` // Status yang ingin diubah (AVAILABLE, BAD, LOST)
}
