package db

import (
	"github.com/crispuscrew/pgxray/internal/cfg"
	"github.com/crispuscrew/pgxray/internal/db"

	"github.com/crispuscrew/pgxray/internal/ui/colors"
	"github.com/crispuscrew/pgxray/internal/ui/common"
	"github.com/crispuscrew/pgxray/internal/ui/loader"

	"time"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	conn *db.Conn
}

func (model Model) Init(initParams common.InitParams, theme colors.Palette) (common.Component, tea.Cmd) {
	return model, connectCmd(initParams.Profile)
}

func (model Model) Update(msg tea.Msg) (common.Component, tea.Cmd) {
	switch msg := msg.(type) {
	case ConnectedMsg:
		model.conn = msg.Conn
		return model, tea.Batch(
			func() tea.Msg { return loader.ProgressMsg{Percent: 0.35} },
			common.After(3 * time.Second, SchemasMsg{}),
		)
	case SchemasMsg:
		return model, tea.Batch(
			func() tea.Msg { return loader.ProgressMsg{Percent: 0.7} },
			common.After(3 * time.Second, DatabaseReadyMsg{}),
		)
	case DatabaseReadyMsg:
		return model, func() tea.Msg { return loader.ProgressMsg{Percent: 1.0} }
	}
	return model, nil
}

type ConnectedMsg 		struct { Conn *db.Conn }
type SchemasMsg 		struct {} //Schemas []string }
type DatabaseReadyMsg 	struct {} //DB *db.Database }

func connectCmd(profile cfg.Profile) tea.Cmd {
	return func() tea.Msg {
		conn, err := db.Connect(profile)
		if err != nil {
			return common.AddCriticalToast{Item : "failed to connect to database: " + err.Error()}
		}
		return ConnectedMsg{Conn: conn}
	}
}