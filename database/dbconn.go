package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func DatabaseConnection() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://bob:admin@localhost:5432/kitabwala")
	if err != nil {
		return err
	}

	DB = conn
	
	return nil
	// return pool, nil
}
