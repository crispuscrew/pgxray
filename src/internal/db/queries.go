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

func (conn *Conn) LoadDatabaseInfo(ctx context.Context) (Database, error) {
	return withTx(ctx, conn, func(tx pgx.Tx) (Database, error) {
		rows, err := tx.Query(ctx, `
			SELECT nspname
			FROM pg_namespace
			WHERE nspname NOT LIKE 'pg_%' 
				AND nspname != 'information_schema'
			ORDER BY nspname
		`)
		if err != nil {return Database{}, err}

		schemas, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Schema, error) {
			var name string
			if err := row.Scan(&name); err != nil { return Schema{}, err }
			return Schema{DBInstance : DBInstance{Name: name}}, nil
		})

		if err != nil { return Database{}, err }

		return Database{
			DBInstance: DBInstance{Name: conn.dbName, ChildrenLoaded: true},
			Schemas: schemas,
			}, nil
	})
}

func (conn *Conn) LoadSchemaInfo(ctx context.Context, schemaName string) (Schema, error) {
	return withTx(ctx, conn, func(tx pgx.Tx) (Schema, error) {
		rows, err := tx.Query(ctx, `
			SELECT tablename
			FROM pg_tables
			WHERE schemaname = $1
			ORDER BY tablename
		`, schemaName)
		if err != nil { return Schema{}, err }

		tables, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Table, error) {
			var name string
			if err := row.Scan(&name); err != nil { return Table{}, err }
			return Table{DBInstance : DBInstance{Name: name}}, nil
		})

		if err != nil { return Schema{}, err }

		return Schema{
			DBInstance: DBInstance{Name: schemaName, ChildrenLoaded: true},
			Tables: tables,
			}, nil
	})
}