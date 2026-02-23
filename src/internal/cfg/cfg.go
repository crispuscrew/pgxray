// Package cfg provides configuration by config file and CLI overrides for pgxray.
package cfg

import (
	"github.com/crispuscrew/pgxray/src/internal/cli"

	"github.com/BurntSushi/toml"
	"log"
)

func BuildConfig(cliOverride CliConfig) Profile, Keybinds {

	return cfg
}