package main

import (
	"fmt"
	"github.com/dmitriibb/go-common/constants"
	"github.com/dmitriibb/go-common/db/pg"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/utils"
	"github.com/dmitriibb/go2/manager/api"
	"github.com/dmitriibb/go2/manager/clientorders"
	"net/http"
)

var logger = logging.NewLogger("ManagerMain")

func main() {
	pg.Init()

	httpPort := utils.GetEnvProperty(constants.HttpPortEnv)
	http.HandleFunc("/", api.Hello)
	http.HandleFunc("/order", clientorders.HttpHandleClientOrder)

	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
	}()
	logger.Debug("http.Listening And Serving...")

	forever := make(chan int)
	<-forever
}
