package controller

import (
	"encoding/json"
	"net/http"

	// "github.com/Zumari/labora-api/config"
	// "github.com/Zumari/labora-api/models"
	"github.com/Zumari/labora-api/services"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items := services.GetAllItems()

	json.NewEncoder(w).Encode(items)
}
