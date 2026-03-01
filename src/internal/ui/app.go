// Package ui provides a terminal user interface for pgxray.
package ui

import (
	"github.com/crispuscrew/pgxray/src/internal/cfg"

	"github.com/crispuscrew/pgxray/src/internal/ui/colors"
	"github.com/crispuscrew/pgxray/src/internal/ui/common"

	"github.com/crispuscrew/pgxray/src/internal/ui/loader"
	"github.com/crispuscrew/pgxray/src/internal/ui/toast"
	"github.com/crispuscrew/pgxray/src/internal/ui/db"

	tea "charm.land/bubbletea/v2"
	"log"
)

func RunUI(config cfg.Config) {
	model := Model{
		loading		: true,
		initParams	: fromConfig(config),
		theme		: colors.Default(),
	}
	
	model.components = map[common.ComponentID]common.Component {
		common.ToastID	: toast.Model{},
		common.LoaderID	: loader.Model{},
		common.DBID		: db.Model{},
	}
	var update common.Component; var cmd tea.Cmd
	for id, component := range model.components {
		update, cmd = component.Init(model.initParams, model.theme)
		model.components[id] = update
		model.initCmd = tea.Batch(model.initCmd, cmd)
	}

	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		log.Fatalf("could not start UI: %v", err)
	}
}