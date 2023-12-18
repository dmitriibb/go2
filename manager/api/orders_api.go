package api

import (
	"dmbb.com/go2/common/model"
	"dmbb.com/go2/manager/orders/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func ClientOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		newOrder(w, r)
	default:
		fmt.Fprintf(w, "Error - unsupported request method %v", r.Method)
	}
}

func newOrder(w http.ResponseWriter, r *http.Request) {
	order := new(model.ClientOrderApi)
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		logger.Error(err.Error())
		fmt.Fprintf(w, "Can't create a new order because '%v'", err.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Can't process new order because '%v'", r)
			fmt.Fprintf(w, "Can't process new order because '%v'", r)
		}
	}()
	service.NewOrder(order)
	fmt.Fprintf(w, "Created new order %v", order)
}
