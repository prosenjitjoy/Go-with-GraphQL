package model

type CategoryInput struct {
	Name string `json:"name"`
}

type Category struct {
	ID string `json:"_key"`
	CategoryInput
}
