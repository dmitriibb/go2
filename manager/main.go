package main

import (
	"dmbb.com/go2/common/db"
	"dmbb.com/go2/manager/api"
	"dmbb.com/go2/manager/clientorders"
	"fmt"
	"net/http"
)

const (
	port = 8080
)

func main() {
	db.TestConnectPostgres()
	http.HandleFunc("/", api.Hello)
	http.HandleFunc("/order", clientorders.HttpHandleClientOrder)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
