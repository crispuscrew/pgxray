package ui

import (
	"github.com/crispuscrew/pgxray/internal/cfg"

	"github.com/crispuscrew/pgxray/internal/ui/colors"
	"github.com/crispuscrew/pgxray/internal/ui/common"

	//"fmt"

	tea "charm.land/bubbletea/v2"
)

var _ tea.Model = Model{}
type Model struct {
	activeMode 	common.UIMode

	profile 	cfg.Profile
	keybinds 	cfg.Keybinds
	theme		colors.Palette

	components 	map[common.CompID]common.Component

	width		int
	height		int
}

func (model Model) Init() (tea.Cmd) {
	cmds := []tea.Cmd{}
	for _, component := range(model.components) {
		cmds = append(cmds, component.Init())
	}
	return tea.Batch(cmds...)
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
    switch msgT := msg.(type) {
	/*case common.AddCriticalToast:
		model.activeMode = common.Critical
		msg = tea.Batch(msg, func() tea.Msg {
			return common.AddInfoToast{
				Item: fmt.Sprintf("Critical failure, press %s to exit", 
					model.keybinds.Quit.Get()[0].Help().Key),
				Timeout: opt.Set(time.Duration(0)),
			}
		})*/
	case tea.WindowSizeMsg:
		model.width, model.height = msgT.Width, msgT.Height
    case tea.KeyPressMsg:
		switch msgT.String() {
		case "q", "ctrl+c":
			return model, tea.Quit
		}
    }

	var compCmd tea.Cmd
	for _, component := range(model.components) {
		compCmd = component.Update(msg);
		if compCmd != nil { cmd = tea.Batch(cmd, compCmd) }
	}
    return model, cmd
}