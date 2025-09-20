package dbconnection

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type DBConnection struct {
	Conn *pgx.Conn
}

func NewDBConnection(ctx context.Context, connString string) (*DBConnection, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return &DBConnection{Conn: conn}, nil
}
