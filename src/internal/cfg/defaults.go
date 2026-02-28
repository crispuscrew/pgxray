package cfg

import (
	"github.com/crispuscrew/pgxray/src/internal/opt"

	"github.com/charmbracelet/bubbles/key"
)

const configPathEnvVar 		= "PGXRAY_CONFIG_PATH"
const keybindsPathEnvVar 	= "PGXRAY_KEYBINDS_PATH"

const defaultConfigPath 	= "~/.config/pgxray/config.toml"
const defaultKeybindsPath 	= "~/.config/pgxray/keybinds.toml"

var defaultProfile = Profile{
	Name: "default",

	Host:     opt.Set("localhost"),
	Port:     opt.Set(5432),
	User:     opt.Set("postgres"),
	Database: opt.Set("postgres"),
	SslMode:  opt.Set("disable"),
}

var defaultKeybinds = Keybinds{
	MoveDown: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j/down", "move down"),
		),
	}),
	MoveUp: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k/up", "move up"),
		),
	}),
	MoveLeft: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("h", "left"),
			key.WithHelp("h/left", "move left"),
		),
	}),
	MoveRight: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("l", "right"),
			key.WithHelp("l/right", "move right"),
		),
	}),

	Select: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
	}),
	Back: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("esc", "backspace"),
			key.WithHelp("esc/backspace", "go back"),
		),
	}),
	SwitchFocus: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch focus"),
		),
	}),

	Search: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "search"),
		),
	}),
	GoToTop: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("g"),
			key.WithHelp("g", "go to top"),
		),
	}),
	GoToBottom: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("G"),
			key.WithHelp("G", "go to bottom"),
		),
	}),

	ViewDDL: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "view DDL"),
		),
	}),
	ViewIndexes: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "view indexes"),
		),
	}),
	OpenPrompt: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys(":"),
			key.WithHelp(":", "open prompt"),
		),
	}),

	NextPage: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("f", "space", "pagedown"),
			key.WithHelp("f/space/pagedown", "next page"),
		),
	}),
	PrevPage: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("b", "pageup"),
			key.WithHelp("b/pageup", "previous page"),
		),
	}),

	Quit: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit"),
		),
	}),
	Help: opt.Set(Keybind{
		key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
	}),
}