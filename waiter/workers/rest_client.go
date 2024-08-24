package workers

import (
	"fmt"
	"github.com/dmitriibb/go-common/restaurant-common/httputils"
	"github.com/dmitriibb/go-common/utils"
	"github.com/dmitriibb/go2/waiter/constants"
	"io"
	"net/http"
)

var managerServiceUrl = fmt.Sprintf("http://%v:%v",
	utils.GetEnvProperty(constants.ManagerServiceUrl),
	utils.GetEnvProperty(constants.ManagerServiceHttpPort))

func getClientIdByOrderIdFromManager(orderId int) string {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/client-orders", managerServiceUrl), nil)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("OrderId", fmt.Sprintf("%v", orderId))

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode > 299 {
		errorResponse := httputils.GetCommonErrorResponse(resp)
		logger.Error("can't get clientId by orderId because %v", errorResponse.Message)
		return ""
	} else {
		respBytes, _ := io.ReadAll(resp.Body)
		return string(respBytes)
	}
}
