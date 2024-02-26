package clientorders

import (
	"context"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/model"
	"dmbb.com/go2/kitchen/kitchenorders"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
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

	// Items for kitchen
	for _, itemApi := range orderApi.Items {
		for i := 0; i < itemApi.Quantity; i++ {
			dishItem := ClientOrderDishItem{
				ClientId:    order.ClientId,
				OrderId:     order.Id,
				DishName:    itemApi.DishName,
				TimeCreated: time.Now(),
				Status:      Created,
			}
			id := SaveOrderDishItemInDb(&dishItem)
			SaveOrderDishItemStatus(id, dishItem.TimeCreated, Created)
		}
	}
	// TODO fix grpc
	sendNewOrderEvent(order)
}

func sendNewOrderEvent(order *ClientOrder) {
	conn, err := grpc.Dial(kitchenServiceGrpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		loggerService.Error("Can't call kitchen grpc because %v", err)
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	defer conn.Close()

	client := kitchenorders.NewKitchenOrdersServiceClient(conn)
	items := make([]*kitchenorders.NewOrderItem, len(order.Items))
	for i, item := range order.Items {
		newItem := &kitchenorders.NewOrderItem{
			DishName: item.DishName,
			Quantity: int32(item.Quantity),
		}
		items[i] = newItem
	}
	response, err := client.PutNewOrder(context.Background(), &kitchenorders.PutNewOrderRequest{OrderId: int32(order.Id), Items: items})
	if err != nil {
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	loggerService.Debug("Sent order %v event to Kitchen. Response = %v", order.Id, response)
}
