package client

import (
	"bytes"
	"client/internal/constants"
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/httputils"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"github.com/dmitriibb/go-common/restaurant-common/model/clientmodel"
	"github.com/dmitriibb/go-common/utils"
	"net/http"
)

var managerServiceUrl = fmt.Sprintf("http://%v:%v",
	utils.GetEnvProperty(constants.ManagerServiceUrl),
	utils.GetEnvProperty(constants.ManagerServiceHttpPort))

var logger = logging.NewLogger("restClient")

//TODO create or find convenient http client or wrapper

func enterRestaurantViaRest(clientName string, clientId string) (*clientmodel.EnterRestaurantResponse, error) {
	request := clientmodel.EnterRestaurantRequest{clientName, clientId}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(request)
	response, err := http.Post(fmt.Sprintf("%v/hostes/enter", managerServiceUrl), "application/json", &buf)
	if err != nil {
		return nil, err
	}

	responseBody := &clientmodel.EnterRestaurantResponse{}
	err = json.NewDecoder(response.Body).Decode(responseBody)
	return responseBody, err
}

func askForMenuViaRest() (*model.MenuDto, error) {
	response, err := http.Get(fmt.Sprintf("%v/menu", managerServiceUrl))
	if err != nil {
		return nil, err
	}

	res := &model.MenuDto{}
	err = json.NewDecoder(response.Body).Decode(res)
	return res, err
}

func makeAnOrderViaRest(clientId string, orderItems []string) string {
	orderItemsDto := make([]model.OrderItemDTO, 0)
	for _, item := range orderItems {
		orderItemsDto = append(orderItemsDto, model.OrderItemDTO{DishName: item, Quantity: 1})
	}
	newOrderDto := model.ClientOrderDTO{ClientId: clientId, Items: orderItemsDto}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(newOrderDto)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/client-orders", managerServiceUrl), &buf)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode > 299 {
		errorResponse := httputils.GetCommonErrorResponse(resp)
		logger.Error("client %s can't make an order because '%s'", clientId, errorResponse.Message)
		return errorResponse.Message
	} else {
		orderResponse := model.ClientOrderResponseDTO{}
		json.NewDecoder(resp.Body).Decode(&orderResponse)
		return orderResponse.Comment
	}
}
