package storage

import (
	"context"
	"database/sql"
	"l0/internal/models"

	"github.com/google/uuid"
)

type PaymentStore interface {
	GetPayment(ctx context.Context, transaction uuid.UUID) (*models.Payment, error)
	CreatePayment(ctx context.Context, p *models.Payment) error
}

type paymentStore struct {
	db *sql.DB
}

func NewPaymentStore(db *sql.DB) PaymentStore {
	return &paymentStore{db: db}
}

func (s *paymentStore) CreatePayment(ctx context.Context, p *models.Payment) error {
	query := `INSERT INTO payment (transaction, request_id, currency, provider, amount, 
				payment_dt, bank, delivery_cost, goods_total, custom_fee)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query,
		p.Transaction,
		p.RequestID,
		p.Currency,
		p.Provider,
		p.Amount,
		p.PaymentDt,
		p.Bank,
		p.DeliveryCost,
		p.GoodsTotal,
		p.CustomFee,
	)
	return err
}

func (s *paymentStore) GetPayment(ctx context.Context, transaction uuid.UUID) (*models.Payment, error) {
	query := `SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, request_id FROM payment WHERE transaction = $1`

	var p models.Payment
	err := s.db.QueryRowContext(ctx, query, transaction).Scan(
		&p.Transaction,
		&p.RequestID,
		&p.Currency,
		&p.Provider,
		&p.Amount,
		&p.PaymentDt,
		&p.Bank,
		&p.DeliveryCost,
		&p.GoodsTotal,
		&p.CustomFee,
		&p.RequestID,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
