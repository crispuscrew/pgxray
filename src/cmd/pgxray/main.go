package main

import (
	"github.com/crispuscrew/pgxray/src/internal/ui"
	"github.com/crispuscrew/pgxray/src/internal/cli"
	"github.com/crispuscrew/pgxray/src/internal/cfg"
)

func main() {
	cfgByCli := cli.Execute()
	profile, keybinds, warnings := cfg.BuildConfig(cfgByCli)
	ui.RunUI(profile, keybinds, warnings)
}