package postgres

import (
	"database/sql"
	"fmt"
	"l0/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.DBName,
			cfg.Postgres.SSLMode,
		))
	if err != nil {
		return nil, err
	}

	return db, nil
}
