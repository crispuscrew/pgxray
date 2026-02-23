package cfg

import (
	"github.com/crispuscrew/pgxray/src/internal/opt"

	"github.com/charmbracelet/bubbletea/key"
)

type CliConfig struct {
	ProfileOverride 	Profile
	ConfigPath 			opt.Opt[string]
	ProfileName 		opt.Opt[string]
}

type Profile struct {
	Name       	string 	`toml:"name"`

	Host     	opt.Opt[string]	`toml:"host" 		cli:"host,H, override profile host"`
	Port     	opt.Opt[int   ]	`toml:"port" 		cli:"port,P, override profile port"`
	User     	opt.Opt[string]	`toml:"user" 		cli:"user,u, override profile user"`
	Database 	opt.Opt[string]	`toml:"database" 	cli:"database,d, override profile database"`
	SslMode  	opt.Opt[string]	`toml:"sslmode" 	cli:"sslmode,,override profile sslmode (disable|require|verify-ca|verify-full)"`
	PgpassFile 	opt.Opt[string]	`toml:"pgpassfile" 	cli:"pgpassfile,,override profile pgpassfile path"`
}

type Keybinds struct {
	MoveDown 	opt.Opt[key.Binding[]]	`toml:"keybind:move_down"`
	MoveUp   	opt.Opt[key.Binding[]]	`toml:"keybind:move_up"`
	MoveLeft 	opt.Opt[key.Binding[]]	`toml:"keybind:move_left"`
	MoveRight 	opt.Opt[key.Binding[]]	`toml:"keybind:move_right"`

	Select 		opt.Opt[key.Binding[]]	`toml:"keybind:select"`
	Back   		opt.Opt[key.Binding[]]	`toml:"keybind:back"`
	SwitchFocus	opt.Opt[key.Binding[]]	`toml:"keybind:switch_focus"`

	Search 		opt.Opt[key.Binding[]]	`toml:"keybind:search"`
	GoToTop		opt.Opt[key.Binding[]]	`toml:"keybind:go_to_top"`
	GoToBottom 	opt.Opt[key.Binding[]]	`toml:"keybind:go_to_bottom"`

	ViewDDL 	opt.Opt[key.Binding[]]	`toml:"keybind:view_ddl"`
	ViewIndexs 	opt.Opt[key.Binding[]]	`toml:"keybind:view_indexes"`
	OpenPrompt 	opt.Opt[key.Binding[]]	`toml:"keybind:open_prompt"`

	NextPage 	opt.Opt[key.Binding[]]	`toml:"keybind:next_page"`
	PrevPage 	opt.Opt[key.Binding[]]	`toml:"keybind:prev_page"`

	Quit 		opt.Opt[key.Binding[]]	`toml:"keybind:quit"`
	Help 		opt.Opt[key.Binding[]]	`toml:"keybind:help"`
}