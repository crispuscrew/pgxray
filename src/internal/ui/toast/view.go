package toast

import (
	"github.com/crispuscrew/pgxray/internal/ui/colors"

	"strings"

	"charm.land/lipgloss/v2"
)


func (model Model) View() string {
	if len(model.items) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, entry := range model.items {
		if _, ok := entry.Text.(Info); ok && model.silentMode {
			continue
		}
		sb.WriteString(entry.Text.render(model.theme))
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (toast Warning) render(p colors.Palette) string {
	return lipgloss.NewStyle().Foreground(p.Warning).Render("⚠ " + toast.Text)
}
func (toast Error) render(p colors.Palette) string {
	return lipgloss.NewStyle().Foreground(p.Error).Render("✗ " + toast.Text)
}
func (toast Info) render(p colors.Palette) string {
	return lipgloss.NewStyle().Foreground(p.Info).Render("• " + toast.Text)
}