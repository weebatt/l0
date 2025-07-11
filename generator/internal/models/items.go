package models

import "github.com/google/uuid"

type Item struct {
	OrderUID    uuid.UUID `json:"order_uid"`
	ChrtID      int       `json:"chrt_id"`
	TrackNumber string    `json:"track_number"`
	Price       int       `json:"price"`
	Rid         string    `json:"rid"`
	Name        string    `json:"name"`
	Sale        int       `json:"sale"`
	Size        string    `json:"size"`
	TotalPrice  int       `json:"total_price"`
	NmID        int       `json:"nm_id"`
	Brand       string    `json:"brand"`
	Status      int       `json:"status"`
}
