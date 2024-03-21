package clientorders

import (
	"context"
	"dmbb.com/go2/manager/constants"
	"errors"
	"fmt"
	"github.com/dmitriibb/go2/common/db/pg"
	"github.com/dmitriibb/go2/common/logging"
	"github.com/dmitriibb/go2/common/model"
	"github.com/dmitriibb/go2/common/utils"
	"github.com/dmitriibb/go2/kitchen/orders/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

const price = 10.0

var kitchenServiceGrpcUrl = utils.GetEnvProperty(constants.KitchenGrpcPortEnv)

var loggerService = logging.NewLogger("ManagerService")

// NewOrder TODO add informative response
func NewOrder(orderApi *model.ClientOrderDTO) *model.ClientOrderResponseDTO {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()
	txWrapper := pg.StartTransaction(ctx)
	order := &ClientOrder{ClientId: orderApi.ClientId}
	order, err := SaveOrderInDb(txWrapper, order)
	if err != nil {
		loggerService.Error("Can't save order in DB because %v", err)
		ctxCancel()
		return &model.ClientOrderResponseDTO{"Fail"}
	}
	items := make([]*ClientOrderItem, 0)
	// Items with prices
	for _, itemApi := range orderApi.Items {
		for i := 0; i < itemApi.Quantity; i++ {
			item := &ClientOrderItem{
				OrderId:  order.Id,
				ClientId: order.ClientId,
				DishName: itemApi.DishName,
				Comment:  itemApi.Comment,
				// TODO add prices
				Price: price,
			}
			item, err := SaveOrderItemInDb(txWrapper, item)
			if err != nil {
				loggerService.Error("Can't save order item in DB because %v", err)
				ctxCancel()
				return &model.ClientOrderResponseDTO{"Fail"}
			}

			if strings.Contains(item.Comment, "error manager") {
				loggerService.Error("error manager")
				ctxCancel()
				return &model.ClientOrderResponseDTO{"Fail"}
			}

			items = append(items, item)
		}
	}
	order.Items = items
	sendNewOrderEvent(ctx, ctxCancel, order)

	err = ctx.Err()
	if err != nil && errors.Is(err, context.Canceled) {
		loggerService.Error("Failed save %s because %s", order, err.Error())
		return &model.ClientOrderResponseDTO{"Fail"}
	}

	txWrapper.Commit()
	return &model.ClientOrderResponseDTO{"Success"}
}

func sendNewOrderEvent(ctx context.Context, ctxCancel context.CancelFunc, order *ClientOrder) {
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
	response, err := client.PutNewOrder(ctx, &handler.PutNewOrderRequest{OrderId: int32(order.Id), Items: items})
	if err != nil {
		panic(fmt.Sprintf("Can't call kitchen grpc because %v", err))
	}
	loggerService.Debug("Sent order %v PutNewOrderRequest to Kitchen. Response = %s", order.Id, response)
	if strings.Contains(response.Status, "error") {
		loggerService.Error("Error from kitchen service. Cancel ctx")
		ctxCancel()
	}
}
