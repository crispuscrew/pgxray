package common

import (
	"github.com/crispuscrew/pgxray/src/internal/cfg"
)

type InitParams struct {
	Profile 	cfg.Profile
	Keybinds 	cfg.Keybinds
	Warnings	[]string
}