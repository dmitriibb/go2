package main

import (
	"context"
	"fmt"
	"github.com/dmitriibb/go2/common/constants"
	"github.com/dmitriibb/go2/common/logging"
	"github.com/dmitriibb/go2/common/utils"
	"github.com/dmitriibb/go2/waiter/receiver"
	"github.com/dmitriibb/go2/waiter/workers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var maiLogger = logging.NewLogger("Main")
var httpPort = utils.GetEnvProperty(constants.HttpPortEnv)

func main() {
	maiLogger.Info("Start")
	rootContext, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(rootContext, cancel)
	receiver.Init(rootContext)
	workers.Init()

	// -------- for testing
	//qConf, _ := rabbit.GetQueueConfig("test2")
	//go func() {
	//	var buffer bytes.Buffer
	//	buffer.WriteString("hello")
	//	for i := 0; i < 7; i++ {
	//		buffer.WriteString(".")
	//		rabbit.SendToQueue(qConf, buffer.String())
	//		time.Sleep(100 * time.Millisecond)
	//	}
	//}()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		maiLogger.Info("dummy http request")
	})
	http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
}

// TODO doesn't work
func gracefulShutdown(rootContext context.Context, rootContextCancelFunc context.CancelFunc) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	waitForCleanupFunc := context.AfterFunc(rootContext, func() {
		fmt.Println("Shut down in.")
		for i := 5; i > 0; i-- {
			fmt.Printf("Shut down in %v...\n", i)
			time.Sleep(time.Second)
		}
	})
	go func() {
		<-s
		fmt.Println("Shutting down gracefully.")
		rootContextCancelFunc()
		waitForCleanupFunc()
		os.Exit(0)
	}()
}
