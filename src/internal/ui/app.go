// Package ui provides a terminal user interface for pgxray.
package ui

import (
	"github.com/crispuscrew/pgxray/internal/cfg"

	"github.com/crispuscrew/pgxray/internal/ui/colors"
	"github.com/crispuscrew/pgxray/internal/ui/common"

	"github.com/crispuscrew/pgxray/internal/ui/toast"

	tea "charm.land/bubbletea/v2"

	"log"
)

func RunUI(config cfg.Config) {
	model := Model{
		activeMode	: common.Init,

		profile		: config.Profile,
		keybinds	: config.Keybinds,
		theme		: colors.Default(),
	}

	model.components = map[common.CompID]common.Component {
		common.ToastID : &toast.Model{Theme : &model.theme},
	}

	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		log.Fatalf("could not start UI: %v", err)
	}
}