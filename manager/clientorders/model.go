package clientorders

type ClientOrderItemStatus string

const (
	Created    ClientOrderItemStatus = "Created"
	InProgress ClientOrderItemStatus = "InProgress"
	Ready      ClientOrderItemStatus = "Ready"
	Served     ClientOrderItemStatus = "Served"
)

type ClientOrderItem struct {
	ClientId string
	OrderId  int
	ItemId   int
	DishName string
	Comment  string
	Price    float32
}

type ClientOrder struct {
	Id       int
	ClientId string
	Items    []*ClientOrderItem
}

const (
	StartProcessingOrder string = "StartProcessing"
	CancelOrder          string = "CancelOrder"
)
