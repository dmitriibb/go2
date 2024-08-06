package testapi

import (
	"github.com/dmitriibb/go-common/logging"
	"net/http"
)

var logger = logging.NewLogger("testapi")

func Hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("Hello")
}
