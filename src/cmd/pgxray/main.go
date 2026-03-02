package main

import (
	"github.com/crispuscrew/pgxray/internal/ui"
	"github.com/crispuscrew/pgxray/internal/cli"
	"github.com/crispuscrew/pgxray/internal/cfg"
)

func main() {
	cfgByCli := cli.Execute()
	config := cfg.BuildConfig(cfgByCli)
	ui.RunUI(config)
}