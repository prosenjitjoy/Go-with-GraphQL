package model

type ReviewInput struct {
	Date      string `json:"date"`
	Title     string `json:"title"`
	Comment   string `json:"comment"`
	Rating    int32  `json:"rating"`
	ProductID string `json:"product_id"`
}

type Review struct {
	ID string `json:"_key"`
	ReviewInput
}
