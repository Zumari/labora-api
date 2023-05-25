package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/Zumari/labora-api/API/models"
	"github.com/Zumari/labora-api/API/services"

	"github.com/gorilla/mux"
)

// JsonResponse returns a JSON type response.
func JsonResponse(response http.ResponseWriter, status int, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("error while marshalling object %v, trace: %+v", data, err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, err = response.Write(bytes)
	if err != nil {

		return fmt.Errorf("error while writing bytes to response writer: %+v", err)
	}

	return nil
}

// GetAllItems returns all elements in JSON format.
func GetAllItems(response http.ResponseWriter, _ *http.Request) {
	items, err := services.GetItems()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error getting items"))

		return
	}

	JsonResponse(response, http.StatusOK, items)
}

// GetItemsPaginated returns the elements in JSON format segmented on pages.
func GetItemsPaginated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pageUser := r.URL.Query().Get("page")
	itemsUser := r.URL.Query().Get("itemsPerPage")

	page, err := strconv.Atoi(pageUser)
	if err != nil || page < 1 {
		page = 1
	}
	itemsPerPage, err := strconv.Atoi(itemsUser)
	if err != nil || itemsPerPage < 1 {
		itemsPerPage = 5
	}

	// Obtener la lista de elementos paginada
	newList, count, err := services.GetItemsPerPage(page, itemsPerPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(itemsPerPage)))

	// Crear un mapa que contiene información sobre la paginación
	paginationInfo := map[string]interface{}{
		"totalPages":  totalPages,
		"currentPage": page,
	}

	// Crear un mapa que contiene la lista de elementos y la información de paginación
	responseData := map[string]interface{}{
		"items":      newList,
		"pagination": paginationInfo,
	}

	// Codificar el mapa de respuesta en formato JSON y enviar en la respuesta HTTP
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Write(jsonData)
}

// GetItemById returns a item sought by su id.
func GetItemById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var item models.Item
	parameters := mux.Vars(request)
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		// Manejar el error de la conversión
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("ID must be a number"))

		return
	}

	item, err = services.GetItemId(id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(response).Encode(item)
}

// GetItemByName returns the items found that have the name to search.
func GetItemByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	parameters := mux.Vars(request)
	name := parameters["name"]

	items, err := services.GetItemName(name)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(response).Encode(items)
}

// CreateItem create a new item.
func CreateItem(response http.ResponseWriter, request *http.Request) {
	var newItem models.Item
	var items []models.Item

	err := json.NewDecoder(request.Body).Decode(&newItem)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Error when processing the application"))

		return
	}

	newItem, err = services.CreateItem(newItem)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Error when processing the application"))

		return
	}

	items, err = services.GetItems()
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Error when processing the application"))

		return
	}

	JsonResponse(response, http.StatusOK, items)
}

// UpdateItem update an item for your id.
func UpdateItem(response http.ResponseWriter, request *http.Request) {
	items, err := services.GetItems()
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}
	parameters := mux.Vars(request)
	var itemUpdate models.Item

	err = json.NewDecoder(request.Body).Decode(&itemUpdate)
	defer request.Body.Close()
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	itemUpdate, err = services.UpdateItem(id, itemUpdate)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}
	items, err = services.GetItems()
	JsonResponse(response, http.StatusOK, items)
}

// DeleteItem delete an item for your id.
func DeleteItem(response http.ResponseWriter, request *http.Request) {
	items, err := services.GetItems()
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}
	parameters := mux.Vars(request)
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	_, err = services.DeleteItem(id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}
	items, err = services.GetItems()
	JsonResponse(response, http.StatusOK, items)

}

// ItemDetails shows the details of the items for their id.
func ItemDetails(response http.ResponseWriter, request *http.Request) {

	var updateItem models.Item
	parameters := mux.Vars(request)
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		fmt.Println(err)
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	updateItem, err = services.UpdateItemDetails(id)
	if err != nil {
		fmt.Println(err)
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}
	JsonResponse(response, http.StatusOK, updateItem)

}

// RootReturn indicates that he is in the root directory.
func RootReturn(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("You are on the root path"))
}
