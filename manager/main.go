package main

import (
	"fmt"
	"github.com/dmitriibb/go2/common/constants"
	"github.com/dmitriibb/go2/common/db/pg"
	"github.com/dmitriibb/go2/common/utils"
	"github.com/dmitriibb/go2/manager/api"
	"github.com/dmitriibb/go2/manager/clientorders"
	"net/http"
)

func main() {
	pg.Init()

	httpPort := utils.GetEnvProperty(constants.HttpPortEnv)
	http.HandleFunc("/", api.Hello)
	http.HandleFunc("/order", clientorders.HttpHandleClientOrder)
	http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
}
