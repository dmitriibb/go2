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
	// TODO start transaction use Context
	order := &ClientOrder{ClientId: orderApi.ClientId}
	order, err := SaveOrderInDb(order)
	if err != nil {
		loggerService.Error("Can't save order in DB because %v", err)
		panic(fmt.Sprintf("Can't save order in DB because %v", err))
	}
	items := make([]*ClientOrderItem, 0)
	// Items with prices
	for _, itemApi := range orderApi.Items {
		for i := 0; i < itemApi.Quantity; i++ {
			item := &ClientOrderItem{
				OrderId:  order.Id,
				ClientId: order.ClientId,
				DishName: itemApi.DishName,
				// TODO add prices
				Price: price,
			}
			item, err := SaveOrderItemInDb(item)
			if err != nil {
				loggerService.Error("Can't save order item in DB because %v", err)
				panic(fmt.Sprintf("Can't save order item in DB because %v", err))
			}
			items = append(items, item)
		}
	}
	order.Items = items
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
			ItemId:   int32(item.ItemId),
			Comment:  item.Comment,
		}
		items[i] = newItem
	}
	response, err := client.PutNewOrder(context.Background(), &handler.PutNewOrderRequest{OrderId: int32(order.Id), Items: items})
	if err != nil {
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	loggerService.Debug("Sent order %v PutNewOrderRequest to Kitchen. Response = %v", order.Id, response)
}
