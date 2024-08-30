package hostes

import (
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"github.com/dmitriibb/go-common/restaurant-common/model/clientmodel"
	"github.com/dmitriibb/go-common/utils/webUtils"
	"net/http"
)

var logger = logging.NewLogger("Hostes")

func HandleMapping(apiPrefix string) {
	http.HandleFunc(fmt.Sprintf("%s/enter", apiPrefix), handleClientEnterRequest)
}

func handleClientEnterRequest(w http.ResponseWriter, r *http.Request) {
	webUtils.EnableCors(w)
	switch r.Method {
	case http.MethodOptions:
		webUtils.HandleOptionsRequest(w, "*", "OPTIONS, POST")
		return
	case http.MethodPost:
		handleClientEnter(w, r)
	}
}

func handleClientEnter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := json.NewEncoder(w).Encode(model.CommonErrorResponse{
			Type:    model.CommonErrorTypeWrongRequest,
			Message: "Only method POST is supported",
		})
		if err != nil {
			logger.Error(err.Error())
		}
	}
	request := new(clientmodel.EnterRestaurantRequest)
	json.NewDecoder(r.Body).Decode(request)

	if len(request.ClientId) < len(request.ClientName) {
		request.ClientId = generateIdForClient(request.ClientName)
	}

	tableNumber := getAvailableTableNumber(request.ClientId)
	if tableNumber < 1 {
		response := clientmodel.EnterRestaurantResponse{
			ClientId: request.ClientId,
			Status:   clientmodel.EnterRestaurantResponseStatusReject,
			Message:  "No available tables",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &clientmodel.EnterRestaurantResponse{
		ClientId:    request.ClientId,
		Status:      clientmodel.EnterRestaurantResponseStatusWelcome,
		TableNumber: tableNumber,
		Message:     fmt.Sprintf("Welcome to the restaurant. Pleasse follow to the table %v", tableNumber),
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Welcome %v and forward him to the table %v", response.ClientId, tableNumber)
	}
}
