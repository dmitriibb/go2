package kitchenorders

import (
	"context"
	"dmbb.com/go2/common/logging"
)

var loggerService = logging.NewLogger("KitchenOrders")

type KitchenOrders struct {
}

func (ko *KitchenOrders) mustEmbedUnimplementedKitchenOrdersServiceServer() {
	panic("Not implemented")
}

var KitchenOrdersService = &KitchenOrders{}

func (ko *KitchenOrders) PutNewOrder(ctx context.Context, in *PutNewOrderRequest) (*PutNewOrderResponse, error) {
	loggerService.Debug("Received new order %v", in)
	return &PutNewOrderResponse{Status: "Received dummy"}, nil
}
