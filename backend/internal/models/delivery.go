package models

import "github.com/google/uuid"

type Delivery struct {
	OrderUID   uuid.UUID `json:"order_uid"`
	DeliveryID int       `json:"delivery_id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Zip        string    `json:"zip"`
	City       string    `json:"city"`
	Address    string    `json:"address"`
	Region     string    `json:"region"`
	Email      string    `json:"email"`
}
