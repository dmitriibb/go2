package clientorders

import (
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/httputils"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"net/http"
	"strconv"
)

var loggerApiHttp = logging.NewLogger("testapi.http.client.orders")

func HandleMapping(apiPrefix string) {
	http.HandleFunc(apiPrefix, handleClientOrder)
}

func handleClientOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		newOrder(w, r)
	case http.MethodGet:
		getClientIdByOrderId(w, r)
	default:
		message := fmt.Sprintf("Error - unsupported request method %v", r.Method)
		httputils.ReturnResponseWithError(w, http.StatusBadRequest, loggerApiHttp, message)
	}
}

func newOrder(w http.ResponseWriter, r *http.Request) {
	order := new(model.ClientOrderDTO)
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		erMsg := fmt.Sprintf("Can't create a new order because '%v'", err.Error())
		httputils.ReturnResponseWithError(w, http.StatusInternalServerError, loggerApiHttp, erMsg)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			erMsg := fmt.Sprintf("Can't process new order because '%v'", r)
			httputils.ReturnResponseWithError(w, http.StatusInternalServerError, loggerApiHttp, erMsg)
		}
	}()
	res := NewOrder(order)
	loggerApiHttp.Debug("Created new order %v", order)
	json.NewEncoder(w).Encode(res)
}

func getClientIdByOrderId(w http.ResponseWriter, r *http.Request) {
	orderId := r.Header.Get("OrderId")
	loggerApiHttp.Debug("search client id for order %v", orderId)

	if len(orderId) == 0 {
		loggerApiHttp.Warn("OrderId header is empty")
		return
	}
	orderIdInt, _ := strconv.Atoi(orderId)

	clientId := getClientIdByOrderIdFromDb(orderIdInt)
	if len(clientId) == 0 {
		loggerApiHttp.Warn("clientId is empty for orderId %v", orderIdInt)
		return
	}
	w.Write([]byte(clientId))
}
