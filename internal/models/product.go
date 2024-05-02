package models

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	UnitId      int    `json:"unit_id"`
	Description string `json:"description"`
	NumberOfProduct int    `json:"number_of_product"`
}

type ProductDTO struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Category        string `json:"category"`
	Price           int    `json:"price"`
	Unit            string `json:"unit"`
	Description     string `json:"description"`
	NumberOfProduct int    `json:"number_of_product"`
}
