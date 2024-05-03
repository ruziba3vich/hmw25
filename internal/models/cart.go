package models

type Cart struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
}
