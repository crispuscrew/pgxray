package toast

import (
	"github.com/crispuscrew/pgxray/src/internal/ui/common"
	"github.com/crispuscrew/pgxray/src/internal/ui/colors"

	"time"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	items		[]entry
	silentMode	bool
	theme		colors.Palette
}

func (model Model) Init(initParams common.InitParams, theme colors.Palette) (common.Component, []tea.Cmd) {
	elements := len(initParams.Warnings)
	cmds := make([]tea.Cmd, 0, elements)
	model = Model{
		items: make([]entry, 0, elements),
		silentMode: true,
		theme: theme,
	}
	var cmd tea.Cmd
	for _, warning := range initParams.Warnings {
		model, cmd = model.Add(Warning{Text: warning})
		cmds = append(cmds, cmd)
	}
	return model, cmds
}

func (model Model) Update(msg tea.Msg) (common.Component, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case AddMsg:
		model, cmd = model.Add(msg.Item)
	case RemoveMsg:
		model, cmd = model.Remove(msg.Item)
	}
	return model, cmd
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

type AddMsg struct { Item toast }
func (model Model) Add(item toast) (Model, tea.Cmd) {
	entry := entry{ID: new(struct{}), Text: item}
	model.items = append(model.items, entry)
	return model, common.After(5 * time.Second, func() tea.Msg { return RemoveMsg{Item: entry} })
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