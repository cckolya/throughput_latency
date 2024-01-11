package main

import (
	"devhands/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(cfg config.Postgres) (*Postgres, error) {
	connectionURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
	open, err := sqlx.Open("pgx", connectionURL)
	if err != nil {
		return nil, err
	}

	err = open.Ping()
	if err != nil {
		return nil, err
	}

	open.SetMaxOpenConns(cfg.Settings.MaxOpenConns)
	open.SetConnMaxLifetime(cfg.Settings.ConnMaxLifeTime * time.Second)
	open.SetMaxIdleConns(cfg.Settings.MaxIdleConns)
	open.SetConnMaxIdleTime(cfg.Settings.MaxIdleLifeTime * time.Second)

	return &Postgres{db: open}, nil
}

func (db *Postgres) SelectDual() error {
	var value int
	return db.db.QueryRow("SELECT 1").Scan(&value)
}
