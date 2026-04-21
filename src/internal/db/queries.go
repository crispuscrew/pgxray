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
			DBInstance: DBInstance{Name: conn.dbName, IsLoaded: true},
			Schemas: schemas,
			}, nil
	})
}

func (conn *Conn) LoadSchemaInfo(ctx context.Context, schemaName string) (Schema, error) {
	return withTx(ctx, conn, func(tx pgx.Tx) (Schema, error) {
		var rows pgx.Rows; var err error

		rows, err = tx.Query(ctx, `
			SELECT tablename
			FROM pg_tables
			WHERE schemaname = $1
			ORDER BY tablename
		`, schemaName)
		if err != nil { return Schema{}, err }

		tables, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Table, error) {
			var tabl Table
			if err := row.Scan(&tabl.Name); err != nil { return Table{}, err }
			return tabl, nil
		})
		if err != nil { return Schema{}, err }

		rows, err = tx.Query(ctx, `
			SELECT viewname, definition
			FROM pg_views
			WHERE schemaname = $1
			ORDER BY viewname
		`, schemaName)
		if err != nil { return Schema{}, err }

		views, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (View, error) {
			var view View
			if err := row.Scan(&view.Name, &view.Definition); err != nil { return View{}, err }
			view.IsLoaded = true
			return view, nil
		})
		if err != nil { return Schema{}, err }

		rows, err = tx.Query(ctx, `
			SELECT sequencename, last_value, min_value, max_value, increment_by
			FROM pg_sequences
			WHERE schemaname = $1
			ORDER BY sequencename
		`, schemaName)
		if err != nil { return Schema{}, err }

		sequences, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Sequence, error) {
			var sequence Sequence
			err := row.Scan(&sequence.Name, &sequence.CurrentValue, &sequence.MinValue, &sequence.MaxValue, &sequence.Increment)
			if err != nil { return Sequence{}, err }
			sequence.IsLoaded = true
			return sequence, nil
		})
		if err != nil { return Schema{}, err }

		return Schema{
			DBInstance	: DBInstance{Name: schemaName, IsLoaded: true},
			Tables		: tables,
			Views		: views,
			Sequences	: sequences,
			}, nil
	})
}

func (conn *Conn) LoadTableInfo(ctx context.Context, schemaName, tableName string) (Table, error) {
	return withTx(ctx, conn, func(tx pgx.Tx) (Table, error) {
		var rows pgx.Rows; var err error

		rows, err = tx.Query(ctx, `
			SELECT 	pg_attribute.attname,
					pg_type.typname,
					pg_attribute.attnotnull,
					pg_get_expr(pg_attrdef.adbin, pg_attrdef.adrelid),				-- from binary to strings (takes oid)
					col_description(pg_attribute.attrelid, pg_attribute.attnum)		-- column comment
			FROM pg_attribute
			JOIN pg_class 		ON pg_class.oid = pg_attribute.attrelid				-- oid = Object ID (Internal)
			JOIN pg_namespace 	ON pg_namespace.oid = pg_class.relnamespace
			JOIN pg_type 		ON pg_type.oid = pg_attribute.atttypid
			LEFT JOIN pg_attrdef ON pg_attrdef.adrelid 	= pg_attribute.attrelid		-- same table
								AND pg_attrdef.adnum 	= pg_attribute.attnum		-- same column (attnum = column num)
			WHERE 	pg_namespace.nspname 	= $1
				AND pg_class.relname 		= $2
				AND pg_attribute.attnum 	> 0										-- reject internal
				AND NOT pg_attribute.attisdropped
			ORDER BY pg_attribute.attnum
		`, schemaName, tableName)

		if err != nil { return Table{}, err }

		columns, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Column, error) {
			var col Column; var tempDefault *string; var tempComment *string
			err := row.Scan(&col.Name, &col.DataType, &col.notNullable, &tempDefault, &tempComment)
			if err != nil { return Column{}, err }
			col.IsLoaded = true
			if tempComment != nil { col.Comment = *tempComment }
			if tempDefault != nil { col.Default = *tempDefault }
			return col, nil
		})
		if err != nil { return Table{}, err }

		rows, err = tx.Query(ctx, `
			SELECT 	pg_class.relname,
					ARRAY (
						SELECT pg_attribute.attname
						FROM pg_attribute
						WHERE 	pg_attribute.attrelid 	= pg_index.indrelid
							AND	pg_attribute.attnum 	= ANY(pg_index.indkey)
						ORDER BY pg_attribute.attnum
					),
					pg_am.amname, 															-- kind : btree, hash and etc
					pg_index.indisunique,
					pg_get_expr(pg_index.indpred, pg_index.indrelid)						-- from bin to str
			FROM pg_index
			JOIN pg_class 		ON pg_class.oid = pg_index.indexrelid
			JOIN pg_am 			ON pg_am.oid 	= pg_class.relam
			JOIN pg_class 		AS indexed_table ON indexed_table.oid = pg_index.indrelid 	-- indexed_table = index from pg_class
			JOIN pg_namespace 	ON pg_namespace.oid = indexed_table.relnamespace
			WHERE 	pg_namespace.nspname 	= $1
				AND indexed_table.relname 	= $2
			ORDER BY pg_class.relname
		`, schemaName, tableName)
		if err != nil { return Table{}, err }

		indexes, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Index, error) {
			var index Index; var tempPartial *string
			err := row.Scan(&index.Name, &index.Columns, &index.Kind, &index.IsUnique, &tempPartial)
			if err != nil { return Index{}, err }
			index.IsLoaded = true
			if tempPartial != nil { index.Partial = *tempPartial }
			return index, nil
		})
		if err != nil { return Table{}, err }

		rows, err = tx.Query(ctx, `
			SELECT 	pg_constraint.conname,
					CASE pg_constraint.contype
						WHEN 'p' THEN 'PRIMARY KEY'
						WHEN 'f' THEN 'FOREIGN KEY'
						WHEN 'c' THEN 'CHECK'
						WHEN 'u' THEN 'UNIQUE'
					END,
					pg_get_constraintdef(pg_constraint.oid)
			FROM pg_constraint
			JOIN pg_class		ON pg_class.oid = pg_constraint.conrelid
			JOIN pg_namespace	ON pg_namespace.oid = pg_class.relnamespace
			WHERE 	pg_namespace.nspname	= $1
				AND pg_class.relname 		= $2
			ORDER BY pg_constraint.conname
		`, schemaName, tableName)
		if err != nil { return Table{}, err }

		constraints, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Constraint, error) {
			var constraint Constraint
			err := row.Scan(&constraint.Name, &constraint.Kind, &constraint.Definition)
			if err != nil { return Constraint{}, err }
			constraint.IsLoaded = true
			return constraint, nil
		})
		if err != nil { return Table{}, err }


		return Table{
			DBInstance: DBInstance{Name: tableName, IsLoaded: true},
			Columns: columns,
			Indexes: indexes,
			Constraints: constraints,
			}, nil
	})
}