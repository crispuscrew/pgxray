package loader

import (
	"github.com/crispuscrew/pgxray/internal/ui/common"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
)

var _ common.Component = (*Model)(nil)
type Model struct {
	Desc	string
	bar		progress.Model
	pct		float64
}

func (model *Model) Init() tea.Cmd {
	model.bar = progress.New()
	return nil
}

func (model *Model) Update(msg tea.Msg) tea.Cmd {
	switch msgT := msg.(type) {
	case progress.FrameMsg:
		updated, cmd := model.bar.Update(msgT)
		model.bar = updated
		if cmd == nil && model.pct >= 1.0 {
			return func() tea.Msg { return common.CompleteLoading{} }
		}
		return cmd
	case SetProgress:
		model.pct = msgT.Pct
		return model.bar.SetPercent(msgT.Pct)
	case tea.WindowSizeMsg:
		model.bar = progress.New(progress.WithWidth(msgT.Width / 2))
		return model.bar.SetPercent(model.pct)
	default:
		return nil
	}
}

type SetProgress struct {Pct float64}