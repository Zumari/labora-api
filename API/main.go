package main

import (
	"log"

	"github.com/Zumari/labora-api/API/config"
	"github.com/Zumari/labora-api/API/controller"
	"github.com/Zumari/labora-api/API/services"

	"github.com/gorilla/mux"
)

func main() {

	services.Connect_DB()
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Root).Methods("GET")
	router.HandleFunc("/items", controller.GetAllItems).Methods("GET")
	router.HandleFunc("/items/page", controller.GetItemsPaginated).Methods("GET")
	router.HandleFunc("/items/details/{id}", controller.ItemDetails).Methods("GET")
	router.HandleFunc("/items/id/{id}", controller.GetItemById).Methods("GET")
	router.HandleFunc("/items/name/{name}", controller.GetItemByName).Methods("GET")

	router.HandleFunc("/items", controller.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controller.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controller.DeleteItem).Methods("DELETE")

	services.Db.PingOrDie()
	if err := config.StartServer(router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
