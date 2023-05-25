package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/Zumari/labora-api/API/models"
)

// GetItems extract all items from the database.
func GetItems() ([]models.Item, error) {
	items := make([]models.Item, 0)
	rows, err := Db.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details, &item.ViewCounter)
		if err != nil {
			fmt.Println(err)

			continue
		}
		items = append(items, item)
	}

	return items, nil
}

// GetItemsPerPage extract from the database all items segmented by pages and quantity per page.
func GetItemsPerPage(pages, itemsPerPage int) ([]models.Item, int, error) {
	// Calculate the initial index and the limit of elements depending on the current page and the elements per page
	start := (pages - 1) * itemsPerPage

	// Obtain the total number of rows in the Items table
	var count int
	err := Db.QueryRow("SELECT COUNT(*) FROM items").Scan(&count)
	if err != nil {

		return nil, 0, err
	}

	// Obtener la lista de elementos correspondientes a la página actual
	rows, err := Db.Query("SELECT * FROM items ORDER BY id OFFSET $1 LIMIT $2", start, itemsPerPage)
	if err != nil {

		return nil, 0, err
	}
	defer rows.Close()

	var newListItems []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details, &item.ViewCounter)
		if err != nil {

			return nil, 0, err
		}
		newListItems = append(newListItems, item)
	}

	if len(newListItems) == 0 {

		return nil, 0, fmt.Errorf("No items found for page %d", pages)
	}

	return newListItems, count, nil
}

var m sync.Mutex

// GetItemId extract from the database the item whose ID coincides with the search.
func GetItemId(id int) (models.Item, error) {
	var item models.Item

	m.Lock()
	Db.QueryRow("UPDATE items SET view_counter = view_counter + $1 WHERE id = $2 RETURNING *",
		1, id)
	m.Unlock()

	err := Db.QueryRow("SELECT * FROM items WHERE id=$1", id).Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details, &item.ViewCounter)
	if err != nil {

		return models.Item{}, err
	}

	return item, nil
}

// GetItemName extract from the database all items whose name matches the search.
func GetItemName(name string) ([]models.Item, error) {
	var items []models.Item
	var item models.Item
	query := fmt.Sprintf("SELECT * FROM items WHERE customer_name ILIKE '%%%s%%'", name)
	rows, err := Db.Query(query)
	if err != nil {

		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details, &item.ViewCounter)
		if err != nil {

			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}

// CreateItem insert a new item to the database.
func CreateItem(newItem models.Item) (models.Item, error) {

	// Insertar el nuevo item en la base de datos
	insertStatement := `INSERT INTO items (customer_name, order_date, product, quantity, price, view_counter, details)
                        VALUES ($1, $2, $3, $4, $5, 0, $6)`
	_, err := Db.Exec(insertStatement, newItem.CustomerName, newItem.OrderDate, newItem.Product, newItem.Quantity, newItem.Price, newItem.Details)
	if err != nil {

		return models.Item{}, err
	}

	return newItem, nil
}

// UpdateItem update an item in the database through its ID.
func UpdateItem(id int, item models.Item) (models.Item, error) {
	var updatedItem models.Item
	row := Db.QueryRow("UPDATE items SET customer_name = $1, order_date = $2, product = $3, quantity = $4, price = $5 WHERE id = $6 RETURNING *",
		item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, id)
	err := row.Scan(&updatedItem.ID, &updatedItem.CustomerName, &updatedItem.OrderDate, &updatedItem.Product, &updatedItem.Quantity, &updatedItem.Price, &item.Details)
	if err != nil {

		return models.Item{}, err
	}

	return updatedItem, nil
}

// DeleteItem delete an item in the database through its id.
func DeleteItem(id int) (models.Item, error) {
	var deleteItem models.Item
	query := fmt.Sprintf("DELETE FROM items WHERE id = %d", id)
	_, err := Db.Exec(query)
	if err != nil {

		return models.Item{}, err
	}

	return deleteItem, nil
}

// getWikipediaDetails obtain from the Wikipedia API information from a specific product.
func getWikipediaDetails(productName string) (string, error) {
	// Let's give the petition to the api of wikipedia
	url := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", productName)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)

		return "", err
	}
	defer resp.Body.Close()

	// Leemos la respuesta de la API
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return "", err
	}

	// Decodificamos el JSON de la respuesta
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Obtenemos el extracto del objeto
	extract := data["extract"].(string)

	return extract, nil
}

// UpdateItemDetails update the details of a product with the information obtained from getWikipediaDetails.
func UpdateItemDetails(id int) (models.Item, error) {
	item, err := GetItemId(id)
	if err != nil {
		fmt.Println(err)
		return models.Item{}, err
	}
	product := strings.ToLower(item.Product)
	// Obtenemos el extracto del objeto desde la API de Wikipedia
	extract, err := getWikipediaDetails(product)
	if err != nil {
		fmt.Println(err)
		return models.Item{}, err
	}

	// Actualizamos la columna "detalles" en la tabla "items"
	var updatedItem models.Item
	row := Db.QueryRow("UPDATE items SET details=$1 WHERE id=$2 RETURNING *",
		extract, id)
	err = row.Scan(&updatedItem.ID, &updatedItem.CustomerName, &updatedItem.OrderDate, &updatedItem.Product, &updatedItem.Quantity, &updatedItem.Price, &updatedItem.Details)
	if err != nil {
		fmt.Println(err)
		return models.Item{}, err
	}
	return updatedItem, nil
}
