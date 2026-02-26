// Package ui provides a terminal user interface for pgxray.
package ui

import (
	"github.com/crispuscrew/pgxray/src/internal/cfg"

	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func RunUI(profile cfg.Profile, keybinds cfg.Keybinds, warnings []string) {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatalf("could not start UI: %v", err)
	}
}

type model struct {}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "q" {
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
	return "Hello, World!"
}