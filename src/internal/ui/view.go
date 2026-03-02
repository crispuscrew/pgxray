package ui

import (
	"github.com/crispuscrew/pgxray/internal/ui/common"

	//"charm.land/lipgloss/v2" // for future styling
	tea "charm.land/bubbletea/v2"
)

func (model Model) View() tea.View {
	var view tea.View
	if model.critical {
		view = tea.NewView(model.components[common.ToastID].View())
	} else if model.loading {
		view = tea.NewView(
			model.components[common.ToastID].View() +
			model.components[common.LoaderID].View(),
		)
	} else {
		view = tea.NewView(model.components[common.ToastID].View())
	}
	view.AltScreen = true
	return view
}