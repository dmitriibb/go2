package main

import (
	"context"
	"fmt"
	"github.com/dmitriibb/go-common/constants"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/utils"
	"github.com/dmitriibb/go2/waiter/receiver"
	"github.com/dmitriibb/go2/waiter/workers"
	"github.com/dmitriibb/go2/waiter/ws"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var mainLogger = logging.NewLogger("Waiter - Main")
var httpPort = utils.GetEnvProperty(constants.HttpPortEnv)

func main() {
	mainLogger.Info("Start")
	rootContext, rootCancel := context.WithCancel(context.Background())
	defer gracefulShutdown2()
	receiver.Init(rootContext)
	workers.Init(rootContext)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

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
		mainLogger.Info("dummy http request")
	})
	ws.HandleMapping("/ws")
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
	}()
	mainLogger.Debug("http.Listening And Serving...")

	forever := make(chan int)
	for {
		select {
		case <-forever:
			mainLogger.Debug("finished from forever.....")
			return
		case <-interrupt:
			mainLogger.Debug("finished from interrupt.....")
			rootCancel()
			return

		}
	}

}

// TODO doesn't work
//func gracefulShutdown(rootContext context.Context, rootContextCancelFunc context.CancelFunc) {
//	s := make(chan os.Signal, 1)
//	signal.Notify(s, os.Interrupt)
//	signal.Notify(s, syscall.SIGTERM)
//	waitForCleanupFunc := context.AfterFunc(rootContext, func() {
//		fmt.Println("Shut down in.")
//		for i := 5; i > 0; i-- {
//			fmt.Printf("Shut down in %v...\n", i)
//			time.Sleep(time.Second)
//		}
//	})
//	go func() {
//		<-s
//		fmt.Println("Shutting down gracefully.")
//		rootContextCancelFunc()
//		waitForCleanupFunc()
//		os.Exit(0)
//	}()
//}

func gracefulShutdown2() {
	fmt.Println("Shut down in.")
	for i := 5; i > 0; i-- {
		fmt.Printf("Shut down in %v...\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Shutting down gracefully.")
}
