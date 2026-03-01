package common

import (
	"github.com/crispuscrew/pgxray/src/internal/ui/colors"

	"time"

	tea "charm.land/bubbletea/v2"
)

type Component interface {
	Init(initParams InitParams, theme colors.Palette) (Component, []tea.Cmd)
	Update(msg tea.Msg) (Component, tea.Cmd)
	View() string
}

type ComponentID int
const (
	ToastID ComponentID = iota
	LoaderID
)

func After(timeout time.Duration, msg tea.Msg) tea.Cmd {
	// tea running gorutine by itself
	return func() tea.Msg {
		time.Sleep(timeout)
		return msg
	}
}