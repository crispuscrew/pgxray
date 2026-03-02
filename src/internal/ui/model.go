package ui

import (
	"github.com/crispuscrew/pgxray/internal/cfg"
	"github.com/crispuscrew/pgxray/internal/opt"
	
	"github.com/crispuscrew/pgxray/internal/ui/colors"
	"github.com/crispuscrew/pgxray/internal/ui/common"

	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	loading 	bool
	critical 	bool

	initParams 	common.InitParams
	initCmd		tea.Cmd
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
	return model.initCmd
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd, updateCmd tea.Cmd
    switch msg := msg.(type) {
	case common.AddCriticalToast:
		model.critical = true
		updateCmd = func() tea.Msg {
			return common.AddInfoToast{
				Item: fmt.Sprintf("Critical failure, press %s to exit", 
					model.initParams.Keybinds.Quit.Get()[0].Help().Key),
				Timeout: opt.Set(time.Duration(0)),
			}
		}
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

	var updated common.Component
	for id, component := range model.components {
		updated, cmd = component.Update(msg);
		model.components[id] = updated
		if cmd != nil {
			updateCmd = tea.Batch(updateCmd, cmd)
		}
	}
    return model, updateCmd
}