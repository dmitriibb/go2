package main

import (
	"dmbb.com/go2/common/db"
	"dmbb.com/go2/manager/api"
	"fmt"
	"net/http"
)

const (
	port = 8080
)

func main() {
	db.TestConnectPostgres()
	http.HandleFunc("/", api.Hello)
	http.HandleFunc("/order", api.ClientOrder)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
