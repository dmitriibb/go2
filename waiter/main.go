package main

import (
	"dmbb.com/go2/common/constants"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/queue/rabbit"
	"dmbb.com/go2/common/utils"
	"fmt"
	"net/http"
)

var maiLogger = logging.NewLogger("Main")
var httpPort = utils.GetEnvProperty(constants.HttpPortEnv)

func main() {
	maiLogger.Info("Start")

	rabbit.TestConnection()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		maiLogger.Info("dummy http request")
	})
	http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
}
