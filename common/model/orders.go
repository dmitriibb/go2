package model

type ClientOrder struct {
	ClientId int    `json:"clientId"`
	DishName string `json:"dishName"`
	Quantity int    `json:"quantity"`
}
