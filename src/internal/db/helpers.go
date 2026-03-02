package db

import (
	"github.com/crispuscrew/pgxray/internal/cfg"

	"net/url"
	"fmt"
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