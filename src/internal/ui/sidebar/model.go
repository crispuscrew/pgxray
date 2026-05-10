package sidebar

import (
	"github.com/crispuscrew/pgxray/internal/opt"
	"github.com/crispuscrew/pgxray/internal/ui/common"

	tea "charm.land/bubbletea/v2"
)

var _ common.Component = (*Model)(nil)
type Model struct {
	DisplayNodes	[]DisplayNode
	TabSize 		int

	CursorPos 		int
	SelectedID 		opt.Opt[int]
}

func (model *Model) Init() tea.Cmd {
	return nil
}

func (model *Model) Update(msg tea.Msg) tea.Cmd {
	return nil
}

type DisplayNode struct {
	ID			int
	Name 		string		// For Node grouping
	DataNodes 	[]DataNode
	Depth 		int
	Expanded 	bool
	Loading 	bool
}

type DataNode struct {
	Name 		string
	Type 		NodeType
	Children 	[]DataNode
}

type NodeType int
const (
	SchemaNode NodeType = iota
	TableNode
	ViewNode
	SequenceNode
)
