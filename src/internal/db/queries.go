package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func withTx[T any](ctx context.Context, conn *Conn, fn func(pgx.Tx) (T, error)) (T, error) {
	tx, err := conn.conn.BeginTx(ctx, pgx.TxOptions{
		AccessMode	: pgx.ReadOnly,
		IsoLevel	: pgx.RepeatableRead,
	})

	if err != nil { var zero T; return zero, err }
	defer func() { _ = tx.Rollback(ctx) }()

	return fn(tx)
}

// Sample stub
func (conn *Conn) LoadDatabaseInfo(ctx context.Context) (Database, error) {
	return withTx(ctx, conn, func(tx pgx.Tx) (Database, error) {
		var db Database
		return db, nil
	})
}
