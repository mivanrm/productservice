package product

type Product struct {
	ID          int64   `json:"product_id" db:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
}

type Variant struct {
	ID         int64   `json:"variant_id" db:"variant_id"`
	ParentID   int64   `json:"parent_id" db:"parent_id"`
	Price      float64 `json:"price" db:"price"`
	OptionName string  `json:"option_name" db:"option_name"`
	Image      string  `json:"image" db:"image"`
	Stock      int64   `json:"stock"`
}

type CreateProductParam struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Stock       int64     `json:"stock"`
	Variants    []Variant `json:"variants"`
}

type UpdateProductParam struct {
	ID          int64     `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Variants    []Variant `json:"variants"`
}

type ProductResponse struct {
	Product  Product   `json:"product"`
	Variants []Variant `json:"variants"`
	Image    []string  `json:"image"`
}
