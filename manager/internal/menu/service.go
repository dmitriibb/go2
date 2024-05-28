package menu

import (
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"github.com/dmitriibb/go-common/utils"
	"github.com/dmitriibb/go2/manager/internal/constants"
	"net/http"
)

var kitchenServiceUrl string
var kitchenServiceHttpPort string

var logger = logging.NewLogger("MenuService")

func init() {
	kitchenServiceUrl = utils.GetEnvProperty(constants.KitchenUrlEnv)
	kitchenServiceHttpPort = utils.GetEnvProperty(constants.KitchenHttpPortEnv)
}

func getMenuFromKitchen() *model.MenuDto {
	response, err := http.Get(fmt.Sprintf("http://%v:%v/recipes/menu", kitchenServiceUrl, kitchenServiceHttpPort))
	if err != nil {
		logger.Error("can't get menu from kitchen because '%v'", err.Error())
	}

	if response.StatusCode >= 300 {
		errorResponse := model.CommonErrorResponse{}
		json.NewDecoder(response.Body).Decode(&errorResponse)
		logger.Error("can't get menu from kitchen because '%+v'", errorResponse)
		return nil
	}

	result := model.MenuDto{}
	json.NewDecoder(response.Body).Decode(&result)
	return &result
}
