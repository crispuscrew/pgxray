package toast

import (
	"github.com/crispuscrew/pgxray/internal/opt"

	"github.com/crispuscrew/pgxray/internal/ui/common"
	"github.com/crispuscrew/pgxray/internal/ui/colors"

	"time"

	tea "charm.land/bubbletea/v2"
)

var _ common.Component = (*Model)(nil)
type Model struct {
	items		[]entry
	silentMode	bool
	Theme		*colors.Palette
}

func (model *Model) Update(msg tea.Msg) tea.Cmd {
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
		return nil
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

func (model *Model) Add(item toast, timeout opt.Opt[time.Duration]) tea.Cmd {
	entry := entry{ID: new(struct{}), Text: item}
	model.items = append(model.items, entry)
	duration := defaultToastTimeout
	if timeout.IsSet() {
		duration = timeout.Get()
		if duration == 0 {
			return nil
		}
	}
	return common.After(duration, RemoveMsg{Item: entry})
}

type RemoveMsg struct { Item entry }
func (model *Model) Remove(toast entry) (tea.Cmd) {
	for i, item := range model.items {
		if item.ID == toast.ID {
			model.items = append(model.items[:i], model.items[i+1:]...)
			break
		}
	}
	return nil
}