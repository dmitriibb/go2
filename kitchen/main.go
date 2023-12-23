package main

import (
	"dmbb.com/go2/common/logging"
	"fmt"
	"google.golang.org/grpc"
	"kitchen/service"
	"net"
	"net/http"
)

const (
	httpPort = 8090
	grpcPort = 8091
)

var logger = logging.NewLogger("KitchenMain")

func main() {
	logger.Info("start")
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
		logger.Info("http started...")
	}()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service.RegisterKitchenServiceServer(grpcServer, &service.KitchenService{})
	logger.Info("Kitchen service registered...")
	grpcServer.Serve(lis)
}
