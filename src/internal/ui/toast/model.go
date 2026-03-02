package toast

import (
	"github.com/crispuscrew/pgxray/internal/opt"

	"github.com/crispuscrew/pgxray/internal/ui/common"
	"github.com/crispuscrew/pgxray/internal/ui/colors"

	"time"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	items		[]entry
	silentMode	bool
	theme		colors.Palette
}

func (model Model) Init(initParams common.InitParams, theme colors.Palette) (common.Component, tea.Cmd) {
	elements := len(initParams.Warnings)
	cmds := make([]tea.Cmd, 0, elements)
	model = Model{
		items: make([]entry, 0, elements),
		silentMode: false,
		theme: theme,
	}
	var cmd tea.Cmd
	for _, warning := range initParams.Warnings {
		model, cmd = model.Add(Warning{Text: warning}, opt.Opt[time.Duration]{})
		cmds = append(cmds, cmd)
	}
	return model, tea.Batch(cmds...)
}

func (model Model) Update(msg tea.Msg) (common.Component, tea.Cmd) {
	switch msg := msg.(type) {
	case common.AddCriticalToast:
		return model.Add(Error{Text: msg.Item}	, msg.Timeout)
	case common.AddErrorToast:
		return model.Add(Error{Text: msg.Item}	, msg.Timeout)
	case common.AddWarningToast:
		return model.Add(Warning{Text: msg.Item}, msg.Timeout)
	case common.AddInfoToast:
		return model.Add(Info{Text: msg.Item}	, msg.Timeout)
	case RemoveMsg:
		return model.Remove(msg.Item)
	default:
		return model, nil
	}
}

type entry struct {
	ID 		*struct{}
	Text 	toast
}

type toast interface {
	render(p colors.Palette) string
}

type Warning struct{ Text string }
type Error   struct{ Text string }
type Info    struct{ Text string }

func (model Model) Add(item toast, timeout opt.Opt[time.Duration]) (Model, tea.Cmd) {
	entry := entry{ID: new(struct{}), Text: item}
	model.items = append(model.items, entry)
	duration := defaultToastTimeout
	if timeout.IsSet() {
		duration = timeout.Get()
		if duration == 0 {
			return model, nil
		}
	}
	return model, common.After(duration, RemoveMsg{Item: entry})
}

type RemoveMsg struct { Item entry }
func (model Model) Remove(toast entry) (Model, tea.Cmd) {
	for i, item := range model.items {
		if item.ID == toast.ID {
			model.items = append(model.items[:i], model.items[i+1:]...)
			break
		}
	}
	return model, nil
}