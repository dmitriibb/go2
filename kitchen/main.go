package main

import (
	"dmbb.com/go2/common/constants"
	"dmbb.com/go2/common/db/mongo"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/utils"
	"fmt"
	"google.golang.org/grpc"
	"kitchen/manager"
	"kitchen/orders/handler"
	"kitchen/recipes"
	"kitchen/storage"
	"net"
	"net/http"
)

var logger = logging.NewLogger("KitchenMain")

func main() {
	logger.Info("start")
	httpPort := utils.GetEnvProperty(constants.HttpPortEnv)
	grpcPort := utils.GetEnvProperty(constants.GrpcPortEnv)

	// init
	mongo.Init()
	recipes.Init()
	closeManagerChan := make(chan string)
	manager.Init(handler.OrdersHandler.NewOrders, closeManagerChan)
	closeStorageChan := make(chan string)
	storage.Init(closeStorageChan)

	// http handle
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
		logger.Info("http started...")
	}()

	// grpc handle
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	handler.RegisterKitchenOrdersHandlerServer(grpcServer, handler.OrdersHandler)
	logger.Info("Kitchen service registered...")
	grpcServer.Serve(lis)
}
