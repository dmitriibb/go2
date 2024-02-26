package handler

import (
	"context"
	"dmbb.com/go2/common/logging"
)

var loggerService = logging.NewLogger("KitchenOrders")

type ordersHandler struct {
	NewOrders chan *PutNewOrderRequest
}

func (ko *ordersHandler) mustEmbedUnimplementedKitchenOrdersHandlerServer() {
	panic("Not implemented")
}

var OrdersHandler = &ordersHandler{NewOrders: make(chan *PutNewOrderRequest, 100)}

func (ko *ordersHandler) PutNewOrder(ctx context.Context, in *PutNewOrderRequest) (*PutNewOrderResponse, error) {
	loggerService.Debug("Received new order %v", in)
	go func() {
		ko.NewOrders <- in
	}()
	return &PutNewOrderResponse{Status: "Received"}, nil
}
