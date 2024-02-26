package model

type OrderItemDTO struct {
	DishName string `json:"dishName"`
	Quantity int    `json:"quantity"`
}

type ClientOrderDTO struct {
	ClientId string         `json:"clientId"`
	Items    []OrderItemDTO `json:"items"`
}
