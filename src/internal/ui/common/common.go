package common

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type Component interface {
	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	View() string
}

func After(timeout time.Duration, msg tea.Msg) tea.Cmd {
	// tea running goroutine by itself
	return func() tea.Msg {
		time.Sleep(timeout)
		return msg
	}
}

func Cmd(msg tea.Msg) tea.Cmd {
	return func ()tea.Msg { return msg }
}