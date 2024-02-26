package model

type DishItem struct {
	OrderId int
	Name    string
	Status  DishItemStatus
}

type DishItemStatus string

const (
	DishItemNew        DishItemStatus = "DishItemNew"
	DishItemInProgress DishItemStatus = "DishItemInProgress"
	DishItemReady      DishItemStatus = "DishItemReady"
)
