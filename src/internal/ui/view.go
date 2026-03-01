package ui

import (
	"github.com/crispuscrew/pgxray/src/internal/ui/common"
	
	//"charm.land/lipgloss/v2"
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	return tea.NewView(m.components[common.ToastID].View())
}