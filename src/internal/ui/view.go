package ui

import (
	"github.com/crispuscrew/pgxray/internal/ui/common"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
		centered := lipgloss.Place(
			model.width, model.height,
			lipgloss.Center, lipgloss.Center,
			model.components[common.LoaderID].View(),
		)
		view = tea.NewView(model.components[common.ToastID].View() + centered)
	case common.Critical:
		view = tea.NewView(model.components[common.ToastID].View())
	default:
		view = tea.NewView(model.components[common.ToastID].View())
	}
	view.AltScreen = true
	return view
}