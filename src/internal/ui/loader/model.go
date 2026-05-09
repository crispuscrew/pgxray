package loader

import (
	"github.com/crispuscrew/pgxray/internal/ui/common"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/progress"
)

var _ common.Component = &Model{}
type Model struct {
	Desc		string
	bar     	progress.Model
	realPct  	float64
}

func (model *Model) Init() *Model {
	model.bar = progress.New()
	return model
}

func (model *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case progress.FrameMsg:
		updated, cmd := model.bar.Update(msg)
		model.bar = updated
		if cmd == nil && model.realPct >= 1.0 {
			return func() tea.Msg { return common.CompleteLoading{} }
		}
		return cmd
	case SetProgress:
		model.realPct = msg.Percent
		return model.bar.SetPercent(msg.Percent)
	default:
		return nil
	}
}

type SetProgress struct {Percent float64}