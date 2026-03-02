package common

import (
	"github.com/crispuscrew/pgxray/internal/cfg"
)

type InitParams struct {
	Profile 	cfg.Profile
	Keybinds 	cfg.Keybinds
	Warnings	[]string
}