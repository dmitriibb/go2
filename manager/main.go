package main

import (
	"dmbb.com/go2/common/constants"
	"dmbb.com/go2/common/db/pg"
	"dmbb.com/go2/common/utils"
	"dmbb.com/go2/manager/api"
	"dmbb.com/go2/manager/clientorders"
	"fmt"
	"net/http"
)

func main() {
	pg.Init()

	httpPort := utils.GetEnvProperty(constants.HttpPortEnv)
	http.HandleFunc("/", api.Hello)
	http.HandleFunc("/order", clientorders.HttpHandleClientOrder)
	http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
}
