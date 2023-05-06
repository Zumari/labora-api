package models

import (
	"log"
	"time"

	"github.com/Zumari/labora-api/config"
)

type Item struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"CurstomerName"`
	OrderDate    time.Time `json:"orderDate"`
	Product      string    `json:"product"`
	Quantity     int       `json:"quantity"`
	Price        float32   `json:"price"`
	Details      string    `json:"details"`
}

var db, error = config.DBConnection()

func GetAll() []Item {
	rows, err := db.Query("SELECT id, customer_name, product, quantity, price  FROM items")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var items []Item = make([]Item, 0)

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.Product, &item.Quantity, &item.Price)

		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)

	}

	return items
}
