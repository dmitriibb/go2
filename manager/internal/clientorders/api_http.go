package clientorders

import (
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"net/http"
)

var loggerApiHttp = logging.NewLogger("api.http.client.orders")

func HttpHandleClientOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		newOrder(w, r)
	default:
		fmt.Fprintf(w, "Error - unsupported request method %v", r.Method)
	}
}

func newOrder(w http.ResponseWriter, r *http.Request) {
	order := new(model.ClientOrderDTO)
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		loggerApiHttp.Error("Can't create a new order because '%v'", err.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			erMsg := fmt.Sprintf("Can't process new order because '%v'", r)
			loggerApiHttp.Error(erMsg)
			fmt.Fprintf(w, erMsg)
		}
	}()
	res := NewOrder(order)
	loggerApiHttp.Debug("Created new order %v", order)
	json.NewEncoder(w).Encode(res)
	//fmt.Fprintf(w, "Created new order %v", order)
}
