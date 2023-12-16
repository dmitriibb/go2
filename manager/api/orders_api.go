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
	order, err := jsonToOrder(r)
	if err != nil {
		logger.Error(err.Error())
		fmt.Fprintf(w, "Can't create a new order because '%v'", err.Error())
		return
	}
	service.NewOrder(order)
	fmt.Fprintf(w, "Created new order %v", order)
}

func jsonToOrder(r *http.Request) (*model.ClientOrder, error) {
	order := new(model.ClientOrder)
	return order, json.NewDecoder(r.Body).Decode(order)
}
