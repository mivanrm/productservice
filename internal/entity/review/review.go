package review

type Review struct {
	ReviewID   int64  `json:"review_id"`
	ProductID  int64  `json:"product_id"`
	ReviewText string `json:"review_text"`
	Rating     int    `json:"rating"`
}
