package ui

import (
	"github.com/crispuscrew/pgxray/src/internal/ui/common"

	//"charm.land/lipgloss/v2"
	tea "charm.land/bubbletea/v2"
)

func (model Model) View() tea.View {
	if model.loading {
		return tea.NewView(model.components[common.LoaderID].View())
	}
	return tea.NewView(model.components[common.ToastID].View())
}