package loader

import (
	"github.com/crispuscrew/pgxray/src/internal/ui/colors"
	"github.com/crispuscrew/pgxray/src/internal/ui/common"

	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/progress"
)

type Model struct {
	bar     progress.Model
	percent float64
}

func (model Model) Init(initParams common.InitParams, theme colors.Palette) (common.Component, []tea.Cmd) {
	model.bar = progress.New(progress.WithColors(theme.Accent, theme.Garmonic))
	return model, []tea.Cmd{tick()}
}

func (model Model) Update(msg tea.Msg) (common.Component, tea.Cmd) {
	switch msg := msg.(type) {
	case progress.FrameMsg:
		updated, cmd := model.bar.Update(msg)
		model.bar = updated
		return model, cmd
	case TickMsg:
		model.percent += 0.02
		if model.percent >= 1.0 {
			model.percent = 1.0
			return model, func() tea.Msg { return common.CompleteLoadingMsg{} }
		}
		return model, tick()
	default:
		return model, nil
	}
}

type TickMsg struct{}

func tick() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(50 * time.Millisecond)
		return TickMsg{}
	}
}