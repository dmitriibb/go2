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
	httpPort := utils.GetEnvProperty(constants.HttpPortEnv)
	pg.TestConnectPostgres()
	http.HandleFunc("/", api.Hello)
	http.HandleFunc("/order", clientorders.HttpHandleClientOrder)
	http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
}
