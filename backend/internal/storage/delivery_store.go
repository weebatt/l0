package storage

import (
	"context"
	"database/sql"
	"l0/internal/models"

	"github.com/google/uuid"
)

type DeliveryStore interface {
	GetDelivery(ctx context.Context, order_uid uuid.UUID) (*models.Delivery, error)
	CreateDelivery(ctx context.Context, d *models.Delivery) error
}

type deliveryStore struct {
	db *sql.DB
}

func NewDeliveryStore(db *sql.DB) DeliveryStore {
	return &deliveryStore{db: db}
}

func (s *deliveryStore) CreateDelivery(ctx context.Context, d *models.Delivery) error {
	query := `INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := s.db.ExecContext(ctx, query,
		d.OrderUID,
		d.Name,
		d.Phone,
		d.Zip,
		d.City,
		d.Address,
		d.Region,
		d.Email,
	)
	return err
}

func (s *deliveryStore) GetDelivery(ctx context.Context, order_uid uuid.UUID) (*models.Delivery, error) {
	query := `SELECT order_uid, name, phone, zip, city, address, region, email, delivery_id FROM delivery WHERE order_uid = $1`

	var d models.Delivery
	err := s.db.QueryRowContext(ctx, query, order_uid).Scan(
		&d.OrderUID,
		&d.Name,
		&d.Phone,
		&d.Zip,
		&d.City,
		&d.Address,
		&d.Region,
		&d.Email,
		&d.DeliveryID,
	)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
