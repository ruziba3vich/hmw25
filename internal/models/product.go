package models

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	UnitId      int    `json:"unit_id"`
	Description string `json:"description"`
}
