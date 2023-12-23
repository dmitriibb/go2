package service

import (
	"context"
	"dmbb.com/go2/common/logging"
)

var logger = logging.NewLogger("KitchenService")

type KitchenService struct {
}

func (service *KitchenService) mustEmbedUnimplementedKitchenServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (service *KitchenService) NewOrderEvent(ctx context.Context, event *OrderEvent) (*OrderEventResponse, error) {
	logger.Info("Received new order event %v", event)
	return &OrderEventResponse{OrderId: event.OrderId, Message: "Received"}, nil
}
