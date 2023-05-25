package models

import (
	"time"
)

type Item struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"CurstomerName"`
	OrderDate    time.Time `json:"orderDate"`
	Product      string    `json:"product"`
	Quantity     int       `json:"quantity"`
	Price        float32   `json:"price"`
	Details      string    `json:"details"`
	TotalPrice   float32   `json:"totalPrice"`
	ViewCounter  int       `json:"viewCounter"`
}

// GeneratorTotalPrice it generates the total price of items based on quantity and price.
func (item Item) GeneratorTotalPrice() float32 {
	totalPrice := item.Price * float32(item.Quantity)
	return totalPrice

}
