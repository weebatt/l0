package storage

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type CacheSaverStore interface {
	AddOrderUID(ctx context.Context, orderUID uuid.UUID) error
	GetAllCachedOrderUIDs(ctx context.Context) ([]uuid.UUID, error)
}

type cacheSaverStore struct {
	db *sql.DB
}

func NewCacheSaverStore(db *sql.DB) CacheSaverStore {
	return &cacheSaverStore{db: db}
}

func (s *cacheSaverStore) AddOrderUID(ctx context.Context, orderUID uuid.UUID) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO cache_saver (order_uid) VALUES ($1)
		 ON CONFLICT (order_uid) DO NOTHING`, orderUID)
	return err
}

func (s *cacheSaverStore) GetAllCachedOrderUIDs(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT order_uid FROM cache_saver`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var uids []uuid.UUID
	for rows.Next() {
		var uid uuid.UUID
		if err := rows.Scan(&uid); err == nil {
			uids = append(uids, uid)
		}
	}
	return uids, nil
}
