package db

import (
	"context"

	"github.com/crispuscrew/pgxray/internal/cfg"
	"github.com/jackc/pgx/v5"

	"fmt"
	"net/url"
)

func profileConnString(profile cfg.Profile) string {
	url := url.URL{
		Scheme	: "postgres",
		Host	: fmt.Sprintf("%s:%d", profile.Host.Get(), profile.Port.Get()),
		User	: url.User(profile.User.Get()),
		Path	: "/" + profile.Database.Get(),
	}
	query := url.Query()
	query.Set("sslmode", profile.SslMode.Get())

	if profile.PgpassFile.IsSet() {
		query.Set("pgpassfile", profile.PgpassFile.Get())
	}

	url.RawQuery = query.Encode()
	return url.String()
}

func withTx[T any](ctx context.Context, conn *Conn, fn func(pgx.Tx) (T, error)) (T, error) {
	tx, err := conn.conn.BeginTx(ctx, pgx.TxOptions{
		AccessMode	: pgx.ReadOnly,
		IsoLevel	: pgx.RepeatableRead,
	})

	if err != nil { var zero T; return zero, err }
	defer func() { _ = tx.Rollback(ctx) }() // Allow us to don't close queries and etc

	return fn(tx)
}

var oidName = map[uint32]string{
	16:   "bool",
	20:   "int8",
	21:   "int2",
	23:   "int4",
	25:   "text",
	700:  "float4",
	701:  "float8",
	1043: "varchar",
	1082: "date",
	1083: "time",
	1114: "timestamp",
	1184: "timestamptz",
	1700: "numeric",
	2950: "uuid",
}

func oidToName(oid uint32) string {
	if name, ok := oidName[oid]; ok { return name }
	return "unknown"
}