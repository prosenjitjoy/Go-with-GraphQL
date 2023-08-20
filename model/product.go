package model

type ProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int32   `json:"quantity"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	OnSale      bool    `json:"on_sale"`
	CategoryID  string  `json:"category_id"`
}

type Product struct {
	ID string `json:"_key"`
	ProductInput
}
