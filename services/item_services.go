package services

import "github.com/Zumari/labora-api/models"

func GetAllItems() []models.Item {
	return models.GetAll()
}
