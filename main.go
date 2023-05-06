package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"github.com/Zumari/labora-api/config"
	"github.com/Zumari/labora-api/controller"
	"github.com/Zumari/labora-api/models"
)

type ItemDetails struct {
	models.Item
	Details string `json:"details"`
}

// func UpdateItem(w http.ResponseWriter, r *http.Request) {
// 	// Función para actualizar un elemento existente
// 	var UpdateItem Item
// 	parametros := mux.Vars(r)
// 	id, _ := strconv.Atoi(parametros["id"])
// 	rqBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Inserte un item válido")
// 		return
// 	}
// 	json.Unmarshal(rqBody, &UpdateItem)
// 	for i, item := range Items {
// 		if item.ID == id {
// 			Items = append(Items[:i], Items[i+1:]...)
// 			UpdateItem.ID = id
// 			Items = append(Items, UpdateItem)
// 		}
// 	}
// 	// Enviar una respuesta exitosa al cliente
// 	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(UpdateItem)
// 	fmt.Println("¡Item actualizado!")
// }

// func DeleteItem(w http.ResponseWriter, r *http.Request) {
// 	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
// 	w.Header().Set("Content-Type", "application/json")
// 	// Función para eliminar un elemento
// 	parametros := mux.Vars(r)
// 	id, _ := strconv.Atoi(parametros["id"])
// 	for i, item := range Items {
// 		if item.ID == id {
// 			Items = append(Items[:i], Items[i+1:]...)
// 			fmt.Fprintf(w, "El item con el ID %s fue eliminada", item.ID)
// 			return
// 		}
// 	}
// 	fmt.Fprintf(w, "Usuario no encontrado")

// }

// func GetItemDetails(id int) ItemDetails {
// 	// Simula la obtención de detalles desde una fuente externa con un time.Sleep
// 	time.Sleep(100 * time.Millisecond)
// 	var foundItem Item
// 	for _, item := range Items {
// 		if item.ID == id {
// 			foundItem = item
// 			break
// 		}
// 	}
// 	//Obviamente, aquí iria un SELECT si es SQL o un llamado a un servicio externo
// 	//pero esta busqueda del item junto con Details, la hacemos a mano.
// 	return ItemDetails{
// 		Item:    foundItem,
// 		Details: fmt.Sprintf("Detalles para el item %d", id),
// 	}
// }

// func GetDetails(w http.ResponseWriter, r *http.Request) {
// 	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
// 	w.Header().Set("Content-Type", "application/json")

// 	wg := &sync.WaitGroup{}
// 	detailsChannel := make(chan ItemDetails, len(Items))
// 	var detailedItems []ItemDetails

// 	// for _, item := range Items {
// 	// 	wg.Add(1) // Creamos el escucha, sin aun crearse la gorutina
// 	// 	go func(id string) {
// 	// 		defer wg.Done() //Completamos el trabajo del escucha, al final de esta ejecución
// 	// 		detailsChannel <- G(id)
// 	// 	}(item.ID)
// 	// }

// 	go func() {
// 		wg.Wait()
// 		close(detailsChannel)
// 	}()

// 	for details := range detailsChannel {
// 		detailedItems = append(detailedItems, details)
// 	}

// 	fmt.Println(detailedItems)
// 	json.NewEncoder(w).Encode(detailedItems)
// }

func main() {

	router := mux.NewRouter() // aca creamos el coso para configurar el servidor, el coso se llama router

	// aca definimos el comportamiento del servidor
	router.HandleFunc("/", controller.HandleRootResource).Methods("GET")
	router.HandleFunc("/otro", controller.HandleOtherResource).Methods("GET")
	router.HandleFunc("/chatgpt", controller.ChatgptHandler).Methods("GET")
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/id/{id}", GetItemID).Methods("GET")
	router.HandleFunc("/items/name/{name}", GetItemName).Methods("GET")
	router.HandleFunc("/items", CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/items/details", GetDetails).Methods("GET")

	// levantar el servidor en un "puerto"
	var portNumber int = 9999
	fmt.Println("Listen in port ", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ADIOS")
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	dbConn, err := config.DBConnection()

	query := `SELECT id, customer_name, order_date, product, quantity, price, details FROM items`
	rows, err := dbConn.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	var items []models.Item
	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)

}

func GetItemID(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	dbConn, err := config.DBConnection()

	// Obtenemos los parámetros de la URL utilizando Gorilla Mux
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	// Iteramos sobre cada elemento en 'Items' para buscar el que tenga el mismo ID que el proporcionado
	var item models.Item
	err = dbConn.QueryRow("SELECT id, customer_name, order_date, product, quantity, price, details FROM items WHERE id=$1", id).Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details)
	if err != nil {
		fmt.Fprintf(w, "Error getting item from database")
		return
	}
	json.NewEncoder(w).Encode(item)

}

func GetItemName(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	dbConn, err := config.DBConnection()

	// Obtenemos los parámetros de la URL utilizando Gorilla Mux
	params := mux.Vars(r)
	name, _ := params["name"]
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	// Iteramos sobre cada elemento en 'Items' para buscar el que tenga el mismo ID que el proporcionado
	var item models.Item

	err = dbConn.QueryRow("SELECT id, customer_name, order_date, product, quantity, price, details FROM items WHERE customer_name ilike $1", name).Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details)
	if err != nil {
		fmt.Fprintf(w, "Error getting item from database")
		return
	}
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {

	dbConn, err := config.DBConnection()
	//Recibimos la información
	var newItem models.Item
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Ingrese un Item válido")
	}

	// Asignamos la informacion a la variable newItem
	json.Unmarshal(reqBody, &newItem)

	// Insertar el nuevo item en la base de datos
	query := `INSERT INTO items (customer_name, order_date, product, quantity, price, details)
                        VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = dbConn.Exec(query, newItem.CustomerName, newItem.OrderDate, newItem.Product, newItem.Quantity, newItem.Price, newItem.Details)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Enviar respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Item creado con éxito")
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {

	dbConn, err := config.DBConnection()

	// Función para actualizar un elemento existente
	var UpdateItem models.Item
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	rqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}
	json.Unmarshal(rqBody, &UpdateItem)

	query := `UPDATE public.items
	SET customer_name=$1, order_date$2, product=$3, quantity=$4, price=$5, details=$6
	WHERE id=$7`

	_, err = dbConn.Exec(query, UpdateItem.CustomerName, UpdateItem.OrderDate, UpdateItem.Product, UpdateItem.Quantity, UpdateItem.Price, UpdateItem.Details, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar una respuesta exitosa al cliente
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UpdateItem)
	fmt.Println("¡Item actualizado!")
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {

	dbConn, err := config.DBConnection()
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Obtenemos los parámetros de la URL utilizando Gorilla Mux
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	// Iteramos sobre cada elemento en 'Items' para buscar el que tenga el mismo ID que el proporcionado
	_, err = dbConn.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		fmt.Fprintf(w, "Error getting item from database")
		return
	}
	//Enviar respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Item eliminado con éxito")
}

func GetDetails(w http.ResponseWriter, r *http.Request) {

}
