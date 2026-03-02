package loader

import (
	"github.com/crispuscrew/pgxray/internal/ui/colors"
	"github.com/crispuscrew/pgxray/internal/ui/common"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/progress"
)

type Model struct {
	bar     	progress.Model
	RealPct  	float64
}

func (model Model) Init(initParams common.InitParams, theme colors.Palette) (common.Component, tea.Cmd) {
	model.bar = progress.New(progress.WithColors(theme.Accent, theme.Garmonic))
	return model, nil
}

func (model Model) Update(msg tea.Msg) (common.Component, tea.Cmd) {
	switch msg := msg.(type) {
	case progress.FrameMsg:
		updated, cmd := model.bar.Update(msg)
		model.bar = updated
		if cmd == nil && model.RealPct >= 1.0 {
			return model, func() tea.Msg { return common.CompleteLoadingMsg{} }
		}
		return model, cmd
	case ProgressMsg:
		model.RealPct = msg.Percent
		return model, model.bar.SetPercent(msg.Percent)
	default:
		return model, nil
	}
}

type ProgressMsg struct {Percent float64}