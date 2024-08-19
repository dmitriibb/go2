package main

import (
	"fmt"
	"github.com/dmitriibb/go-common/constants"
	"github.com/dmitriibb/go-common/db/pg"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/utils"
	"github.com/dmitriibb/go2/manager/internal/clientorders"
	"github.com/dmitriibb/go2/manager/internal/hostes"
	"github.com/dmitriibb/go2/manager/internal/menu"
	"github.com/dmitriibb/go2/manager/internal/testapi"
	"github.com/dmitriibb/go2/manager/internal/ws"
	"net/http"
)

var logger = logging.NewLogger("ManagerMain")

func main() {
	pg.Init()
	clientorders.Init()

	httpPort := utils.GetEnvProperty(constants.HttpPortEnv)
	http.HandleFunc("/", testapi.Hello)
	clientorders.HandleMapping("/client-orders")
	hostes.HandleMapping("/hostes")
	menu.HandleMapping("/menu")
	ws.HandleMapping("/ws")

	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
	}()
	logger.Debug("http.Listening And Serving...")

	forever := make(chan int)
	<-forever
}
