package sidebar

import (
	"github.com/crispuscrew/pgxray/src/internal/ui/common"
	"github.com/crispuscrew/pgxray/src/internal/opt"
)

type Model struct {
	ViewNodes 	[]ViewNode
	TabSize 	int

	CursorPos 	int
	SelectedID 	opt.Option[int]
}

func (model Model) Init(initParams common.InitParams, theme colors.Palette) (common.Component, []tea.Cmd) {
	return model, nil
}

func (model Model) Update(msg tea.Msg) (common.Component, tea.Cmd) {
	return model, nil
}

type ViewNode struct {
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
