package db

import (
	"github.com/crispuscrew/pgxray/internal/cfg"

	"context"

	"github.com/jackc/pgx/v5"
)

type Conn struct {
	conn *pgx.Conn 
}

func Connect(profile cfg.Profile) (*Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultConnectionTimeout)
	defer cancel()

	conn, err := pgx.Connect(ctx, profileConnString(profile))
	
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn}, nil
}