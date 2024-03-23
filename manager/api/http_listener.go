package api

import (
	"github.com/dmitriibb/go-common/logging"
	"net/http"
)

var logger = logging.NewLogger("api")

func Hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("Hello")
}
