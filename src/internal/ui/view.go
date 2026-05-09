package ui

import (
	"github.com/crispuscrew/pgxray/internal/ui/common"

	//"charm.land/lipgloss/v2" // for future styling
	tea "charm.land/bubbletea/v2"
)

/*
Mode	|  	Toasts
-------------------
Tree 	| 	Data
		|  	Prompt (opt)
*/
func (model Model) View() tea.View {
	var view tea.View
	switch model.activeMode {
	case common.Init:
		view = tea.NewView(
			model.components[common.ToastID].View() +
			model.components[common.LoaderID].View(),
		)
	case common.Critical:
		view = tea.NewView(model.components[common.ToastID].View())
	default:
		view = tea.NewView(model.components[common.ToastID].View())
	}
	view.AltScreen = true
	return view
}