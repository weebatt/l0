package models

import (
	"github.com/google/uuid"
)

type Order struct {
	OrderUID          uuid.UUID `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       string    `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
	InternalSignature string    `json:"internal_signature"`
}
