package api

import (
	"github.com/dmitriibb/go2/common/logging"
	"net/http"
)

var logger = logging.NewLogger("api")

func Hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("Hello")
}
