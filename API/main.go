package main

import (
	"log"

	"github.com/Zumari/labora-api/API/config"
	"github.com/Zumari/labora-api/API/controller"
	"github.com/Zumari/labora-api/API/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	services.Connect_DB()
	router := mux.NewRouter()

	// RUTAS
	router.HandleFunc("/", controller.RootReturn).Methods("GET")
	router.HandleFunc("/items", controller.GetAllItems).Methods("GET")
	router.HandleFunc("/items/page", controller.GetItemsPaginated).Methods("GET")
	router.HandleFunc("/items/details/{id}", controller.ItemDetails).Methods("GET")
	router.HandleFunc("/items/id/{id}", controller.GetItemById).Methods("GET")
	router.HandleFunc("/items/name/{name}", controller.GetItemByName).Methods("GET")

	router.HandleFunc("/items", controller.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controller.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controller.DeleteItem).Methods("DELETE")

	// Configura las opciones de CORS. Por ejemplo, permite todas las origenes:
	// corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsOptions := handlers.AllowedMethods([]string{"GET"})

	// Envolviendo tus rutas con CORS.
	handler := handlers.CORS(corsOptions)(router)

	services.Db.PingOrDie()
	if err := config.StartServer(handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
