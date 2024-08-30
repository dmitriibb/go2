package testapi

import (
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/utils/webUtils"
	"io"
	"net/http"
)

var logger = logging.NewLogger("testapi")

func Hello(w http.ResponseWriter, r *http.Request) {
	webUtils.EnableCors(w)
	logger.Info("Hello")
	_, err := io.WriteString(w, "Hello world")
	if err != nil {
		logger.Error("can't send hello world to http request because %v", err.Error())
	}
}
