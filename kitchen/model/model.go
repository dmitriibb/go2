package model

type OrderItem struct {
	OrderId int
	ItemId  int
	Name    string
	Comment string
	Status  OrderItemStatus
}

type OrderItemStatus string

const (
	OrderItemNew        OrderItemStatus = "New"
	OrderItemInProgress OrderItemStatus = "InProgress"
	OrderItemReady      OrderItemStatus = "Ready"
	OrderItemError      OrderItemStatus = "Error"
)
