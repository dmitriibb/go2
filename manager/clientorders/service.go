package clientorders

import (
	"context"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/model"
	"dmbb.com/go2/kitchen/orders/handler"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const price = 10.0
const kitchenServiceGrpcUrl = ":8091"

var loggerService = logging.NewLogger("ManagerService")

// NewOrder TODO add informative response
func NewOrder(orderApi *model.ClientOrderDTO) {
	// TODO start transaction
	order := &ClientOrder{ClientId: orderApi.ClientId}
	order, err := SaveOrderInDb(order)
	if err != nil {
		loggerService.Error("Can't save order in DB because %v", err)
		panic(fmt.Sprintf("Can't save order in DB because %v", err))
	}
	items := make([]ClientOrderItem, 0)
	// Items with prices
	for _, itemApi := range orderApi.Items {
		item := ClientOrderItem{
			OrderId:  order.Id,
			ClientId: order.ClientId,
			DishName: itemApi.DishName,
			Quantity: itemApi.Quantity,
			Price:    float32(itemApi.Quantity) * price,
		}
		SaveOrderItemInDb(&item)
		items = append(items, item)
	}
	order.Items = items

	// Items for kitchen
	//for _, itemApi := range orderApi.Items {
	//	for i := 0; i < itemApi.Quantity; i++ {
	//		dishItem := ClientOrderDishItem{
	//			ClientId:    order.ClientId,
	//			OrderId:     order.Id,
	//			DishName:    itemApi.DishName,
	//			TimeCreated: time.Now(),
	//			Status:      Created,
	//		}
	//		id := SaveOrderDishItemInDb(&dishItem)
	//		SaveOrderDishItemStatus(id, dishItem.TimeCreated, Created)
	//	}
	//}
	sendNewOrderEvent(order)
}

func sendNewOrderEvent(order *ClientOrder) {
	conn, err := grpc.Dial(kitchenServiceGrpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		loggerService.Error("Can't call kitchen grpc because %v", err)
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	defer conn.Close()

	client := handler.NewKitchenOrdersHandlerClient(conn)
	items := make([]*handler.NewOrderItem, len(order.Items))
	for i, item := range order.Items {
		newItem := &handler.NewOrderItem{
			DishName: item.DishName,
			Quantity: int32(item.Quantity),
		}
		items[i] = newItem
	}
	response, err := client.PutNewOrder(context.Background(), &handler.PutNewOrderRequest{OrderId: int32(order.Id), Items: items})
	if err != nil {
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	loggerService.Debug("Sent order %v PutNewOrderRequest to Kitchen. Response = %v", order.Id, response)
}
