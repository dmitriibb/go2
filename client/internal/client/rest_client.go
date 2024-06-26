package client

import (
	"bytes"
	"client/internal/constants"
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"github.com/dmitriibb/go-common/restaurant-common/model/clientmodel"
	"github.com/dmitriibb/go-common/utils"
	"net/http"
)

var managerServiceUrl = fmt.Sprintf("http://%v:%v",
	utils.GetEnvProperty(constants.ManagerServiceUrl),
	utils.GetEnvProperty(constants.ManagerServiceHttpPort))

func enterRestaurantRest(clientName string, clientId string) (*clientmodel.EnterRestaurantResponse, error) {
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

func askForMenu() (*model.MenuDto, error) {
	response, err := http.Get(fmt.Sprintf("%v/menu", managerServiceUrl))
	if err != nil {
		return nil, err
	}

	res := &model.MenuDto{}
	err = json.NewDecoder(response.Body).Decode(res)
	return res, err
}
