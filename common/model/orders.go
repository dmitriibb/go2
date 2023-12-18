package model

type OrderItemApi struct {
	DishName string `json:"dishName"`
	Quantity int    `json:"quantity"`
}

type ClientOrderApi struct {
	ClientId string         `json:"clientId"`
	Items    []OrderItemApi `json:"items"`
}

type OrderItem struct {
	ClientId string
	OrderId  int
	DishName string
	Quantity int
	Price    float32
}

type ClientOrder struct {
	Id       int
	ClientId string
	Items    []OrderItem
}
