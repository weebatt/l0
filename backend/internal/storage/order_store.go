package storage

import (
	"context"
	"database/sql"
	"l0/internal/models"
	"l0/pkg/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type OrderStore interface {
	CreateOrder(
		ctx context.Context,
		order_uid uuid.UUID,
		track_number string,
		entry string,
		locale string,
		customer_id string,
		delivery_service string,
		shardkey string,
		sm_id int,
		date_created string,
		oof_shard string,
		internal_signature string,
	) error
	GetOrder(
		ctx context.Context,
		order_uid uuid.UUID,
	) (*models.Order, error)
	GetAllOrders(ctx context.Context) ([]*models.Order, error)
}

type orderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) OrderStore {
	return &orderStore{db: db}
}

func (s *orderStore) CreateOrder(
	ctx context.Context,
	order_uid uuid.UUID,
	track_number string,
	entry string,
	locale string,
	customer_id string,
	delivery_service string,
	shardkey string,
	sm_id int,
	date_created string,
	oof_shard string,
	internal_signature string,
) error {
	query := `INSERT INTO orders(order_uid, track_number, entry, locale, customer_id, 
            delivery_service, shardkey, sm_id, date_created, oof_shard, internal_signature)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := s.db.ExecContext(
		ctx,
		query,
		order_uid,
		track_number,
		entry,
		locale,
		customer_id,
		delivery_service,
		shardkey,
		sm_id,
		date_created,
		oof_shard,
		internal_signature,
	)
	if err != nil {
		logger.GetFromContext(ctx).Error("troubles with creating order", zap.Error(err))
	}
	return err
}

func (s *orderStore) GetOrder(ctx context.Context, order_uid uuid.UUID) (*models.Order, error) {
	query := `SELECT * FROM orders WHERE order_uid = $1`

	var order models.Order
	err := s.db.QueryRowContext(ctx, query, order_uid).
		Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.OofShard,
			&order.InternalSignature,
		)

	if err != nil {
		logger.GetFromContext(ctx).Error("troubles with getting order", zap.Error(err))
	}

	return &order, nil
}

func (s *orderStore) GetAllOrders(ctx context.Context) ([]*models.Order, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT order_uid, track_number, entry, locale, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard, internal_signature FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
			&order.InternalSignature,
		)
		if err != nil {
			continue
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
