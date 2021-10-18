package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

func Open(cfg *Config) (*sql.DB, func(), error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.DbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, fmt.Errorf("connecting to database: %w", err)
	}

	return db, func() { db.Close() }, nil
}
