package product

type Product struct {
	ID          int64   `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}
