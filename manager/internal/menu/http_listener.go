package menu

import (
	"encoding/json"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"net/http"
)

// Temporal solution. Manager should update the menu in the startup and save it in it's own DB.
func HandleMenuRequest(w http.ResponseWriter, _ *http.Request) {
	menu := getMenuFromKitchen()
	w.Header().Set("content-Type", "application/json")
	if menu == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.CommonErrorResponse{
			Type:    model.CommonErrorTypeInvalidData,
			Message: "Something went wrong",
		})
		return
	}
	json.NewEncoder(w).Encode(menu)
}
