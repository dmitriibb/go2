package clientorders

import "time"

type ClientOrderDishItemStatus string

type ClientOrderItem struct {
	ClientId string
	OrderId  int
	DishName string
	Quantity int
	Price    float32
}

type ClientOrder struct {
	Id       int
	ClientId string
	Items    []ClientOrderItem
}

type ClientOrderDishItem struct {
	ClientId    string
	OrderId     int
	DishName    string
	TimeCreated time.Time
	Status      ClientOrderDishItemStatus
}

const (
	Created    ClientOrderDishItemStatus = "Created"
	InProgress ClientOrderDishItemStatus = "InProgress"
	Ready      ClientOrderDishItemStatus = "Ready"
	Served     ClientOrderDishItemStatus = "Served"
)

const (
	StartProcessingOrder string = "StartProcessing"
	CancelOrder          string = "CancelOrder"
)
