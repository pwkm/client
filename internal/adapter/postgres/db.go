package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pwkm/client/internal/core/domain"
	"github.com/pwkm/client/internal/util"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// New creates a new PostgreSQL database instance
func PostgresNew(config *util.Container) (*sql.DB, error) {
	ConnectString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB.Connection,
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
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

func SaveClient(client *domain.Client) error {

	return nil
}

func FindByID(id string) (*domain.Client, error) {

	return nil, nil
}

func SaveLogin(login *domain.Login) error {

	return nil
}
