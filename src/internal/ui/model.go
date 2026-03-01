package ui

import (
	"github.com/crispuscrew/pgxray/src/internal/cfg"
	"github.com/crispuscrew/pgxray/src/internal/ui/colors"
	"github.com/crispuscrew/pgxray/src/internal/ui/common"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	//conn 		*db.Conn
	loading 	bool

	initParams 	common.InitParams
	initCmds	[]tea.Cmd
	theme		colors.Palette

	components 	map[common.ComponentID]common.Component

	width		int
	height		int
}


func fromConfig(config cfg.Config) common.InitParams {
	return common.InitParams{
		Profile:  config.Profile,
		Keybinds: config.Keybinds,
		Warnings: config.Warnings,
	}
}

func (model Model) Init() tea.Cmd {
	return tea.Batch(model.initCmds...)
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.width, model.height = msg.Width, msg.Height
    case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return model, tea.Quit
		}
	case common.CompleteLoadingMsg:
		model.loading = false
    }

	var cmds []tea.Cmd; var cmd tea.Cmd
	var updated common.Component
	for id, component := range model.components {
		updated, cmd = component.Update(msg);
		model.components[id] = updated
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}
    return model, tea.Batch(cmds...)
}