package storage

import (
	"context"
	"database/sql"
	"l0/internal/models"
)

type ItemsStore interface {
	GetItems(ctx context.Context, order_uid string) ([]models.Item, error)
	CreateItems(ctx context.Context, items []models.Item) error
}

type itemsStore struct {
	db *sql.DB
}

func NewItemsStore(db *sql.DB) ItemsStore {
	return &itemsStore{db: db}
}

func (s *itemsStore) CreateItems(ctx context.Context, items []models.Item) error {
	query := `INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, 
				size, total_price, nm_id, brand, status)
			  	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	for _, it := range items {
		_, err := s.db.ExecContext(ctx, query,
			it.OrderUID,
			it.ChrtID,
			it.TrackNumber,
			it.Price,
			it.Rid,
			it.Name,
			it.Sale,
			it.Size,
			it.TotalPrice,
			it.NmID,
			it.Brand,
			it.Status,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *itemsStore) GetItems(ctx context.Context, order_uid string) ([]models.Item, error) {
	query := `SELECT * FROM items WHERE order_uid = $1`

	rows, err := s.db.QueryContext(ctx, query, order_uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var it models.Item
		err := rows.Scan(
			&it.OrderUID,
			&it.ChrtID,
			&it.TrackNumber,
			&it.Price,
			&it.Rid,
			&it.Name,
			&it.Sale,
			&it.Size,
			&it.TotalPrice,
			&it.NmID,
			&it.Brand,
			&it.Status,
		)
		if err == nil {
			items = append(items, it)
		}
	}
	return items, nil
}
