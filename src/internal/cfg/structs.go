package cfg

import (
	"github.com/crispuscrew/pgxray/src/internal/opt"

	"github.com/charmbracelet/bubbles/key"
)

type CliConfig struct {
	ProfileOverride 	Profile
	ConfigPath 			opt.Opt[string]
	ProfileName 		opt.Opt[string]
	KeybindsPath 		opt.Opt[string]
}

type Profile struct {
	Name       	string 	`toml:"name"`

	Host     	opt.Opt[string]	`toml:"host" cli:"host,H, override profile host"`
	Port     	opt.Opt[int   ]	`toml:"port" cli:"port,P, override profile port"`
	User     	opt.Opt[string]	`toml:"user" cli:"user,u, override profile user"`
	Database 	opt.Opt[string]	`toml:"database" cli:"database,d, override profile database"`
	SslMode  	opt.Opt[string]	`toml:"sslmode" cli:"sslmode,,override profile sslmode (disable|require|verify-ca|verify-full)"`
	PgpassFile 	opt.Opt[string]	`toml:"pgpassfile" cli:"pgpassfile,,override profile pgpassfile path"`
}

type Keybind []key.Binding

type Keybinds struct {
	MoveDown 	opt.Opt[Keybind]	`toml:"move_down"`
	MoveUp   	opt.Opt[Keybind]	`toml:"move_up"`
	MoveLeft 	opt.Opt[Keybind]	`toml:"move_left"`
	MoveRight 	opt.Opt[Keybind]	`toml:"move_right"`

	Select 		opt.Opt[Keybind]	`toml:"select"`
	Back   		opt.Opt[Keybind]	`toml:"back"`
	SwitchFocus	opt.Opt[Keybind]	`toml:"switch_focus"`

	Search 		opt.Opt[Keybind]	`toml:"search"`
	GoToTop		opt.Opt[Keybind]	`toml:"go_to_top"`
	GoToBottom 	opt.Opt[Keybind]	`toml:"go_to_bottom"`

	ViewDDL 	opt.Opt[Keybind]	`toml:"view_ddl"`
	ViewIndexes	opt.Opt[Keybind]	`toml:"view_indexes"`
	OpenPrompt 	opt.Opt[Keybind]	`toml:"open_prompt"`

	NextPage 	opt.Opt[Keybind]	`toml:"next_page"`
	PrevPage 	opt.Opt[Keybind]	`toml:"prev_page"`

	Quit 		opt.Opt[Keybind]	`toml:"quit"`
	Help 		opt.Opt[Keybind]	`toml:"help"`
}