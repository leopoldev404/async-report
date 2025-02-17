package store

import (
	"database/sql"
	"fmt"

	"github.com/leopoldev404/async-report/service/config"
)

func NewDb(config *config.DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Url); if err != nil {
		return nil, fmt.Errorf("failed open db connection: %w", err)
	}
	return db, nil;
}