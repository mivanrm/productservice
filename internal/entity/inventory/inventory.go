package inventory

type Inventory struct {
	StockID   int64  `json:"stock_id"`
	VariantID int64  `json:"variant_id"`
	Amount    string `json:"amount"`
}
