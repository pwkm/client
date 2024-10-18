package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/GoesToEleven/golang-web-dev/046_mongodb/15_postgres/config"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
}

// New creates a new PostgreSQL database instance
func New(ctx context.Context, config *config.DB) (*sql.DB, error) {
	ConnectString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open(config.DB.Connection, ConnectString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
