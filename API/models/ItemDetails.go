package models

type ItemDetails struct {
	Item
	Details string `json:"details"`
}
