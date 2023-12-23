package service

import (
	"context"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/model"
	kitchenService "dmbb.com/go2/kitchen/service"
	"dmbb.com/go2/manager/orders/repository"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const price = 10.0
const kitchenServiceGrpcUrl = ":8091"

var logger = logging.NewLogger("ManagerService")

// NewOrder TODO add informative response
func NewOrder(orderApi *model.ClientOrderApi) {
	order := &model.ClientOrder{ClientId: orderApi.ClientId}
	order = repository.SaveOrderInDb(order)
	items := make([]model.OrderItem, 0)
	// Items with prices
	for _, itemApi := range orderApi.Items {
		item := model.OrderItem{
			OrderId:  order.Id,
			ClientId: order.ClientId,
			DishName: itemApi.DishName,
			Quantity: itemApi.Quantity,
			Price:    float32(itemApi.Quantity) * price,
		}
		repository.SaveOrderItemInDb(&item)
		items = append(items, item)
	}

	// Items for kitchen
	for _, itemApi := range orderApi.Items {
		for i := 0; i < itemApi.Quantity; i++ {
			dishItem := model.OrderDishItem{
				ClientId:    order.ClientId,
				OrderId:     order.Id,
				DishName:    itemApi.DishName,
				TimeCreated: time.Now(),
				Status:      model.Created,
			}
			id := repository.SaveOrderDishItemInDb(&dishItem)
			repository.SaveOrderDishItemStatus(id, dishItem.TimeCreated, model.Created)
		}
	}
	sendNewOrderEvent(order)
}

func sendNewOrderEvent(order *model.ClientOrder) {
	conn, err := grpc.Dial(kitchenServiceGrpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Can't call kitchen grpc because %v", err)
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	defer conn.Close()

	client := kitchenService.NewKitchenServiceClient(conn)
	orderEvent := kitchenService.OrderEvent{OrderId: int32(order.Id), Type: "Created"}
	response, err := client.NewOrderEvent(context.Background(), &orderEvent)
	if err != nil {
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	logger.Debug("Sent order %v event to Kitchen. Response = %v", order.Id, response)
}
