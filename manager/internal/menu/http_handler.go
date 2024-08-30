package menu

import (
	"encoding/json"
	"github.com/dmitriibb/go-common/restaurant-common/httputils"
	"github.com/dmitriibb/go-common/utils/webUtils"
	"net/http"
)

func HandleMapping(apiPrefix string) {
	http.HandleFunc(apiPrefix, getMenu)
}

// Temporal solution. Manager should update the menu in the startup and save it in it's own DB.
func getMenu(w http.ResponseWriter, _ *http.Request) {
	webUtils.EnableCors(w)
	menu := getMenuFromKitchen()
	w.Header().Set("content-Type", "application/json")
	if menu == nil {
		httputils.ReturnResponseWithError(w, http.StatusInternalServerError, logger, "Something went wrong")
		return
	}
	json.NewEncoder(w).Encode(menu)
}
