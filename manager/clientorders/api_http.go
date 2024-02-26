package clientorders

import (
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/model"
	"encoding/json"
	"fmt"
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
			loggerApiHttp.Error("Can't process new order because '%v'", r)
		}
	}()
	NewOrder(order)
	loggerApiHttp.Debug("Created new order %v", order)
	fmt.Fprintf(w, "Created new order %v", order)
}
