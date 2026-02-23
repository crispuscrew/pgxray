package cfg

var defaultProfile = Profile{
	Name: "default",

	Host:     opt.Set("localhost"),
	Port:     opt.Set(5432),
	User:     opt.Set("postgres"),
	Database: opt.Set("postgres"),
	SslMode:  opt.Set("disable"),
}

var defaultKeybinds = Keybinds{
	MoveDown: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j/down", "move down"),
		),
	}),
	MoveUp: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k/up", "move up"),
		),
	}),
	MoveLeft: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("h", "left"),
			key.WithHelp("h/left", "move left"),
		),
	}),
	MoveRight: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("l", "right"),
			key.WithHelp("l/right", "move right"),
		),
	}),

	Select: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
	}),
	Back: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("esc", "backspace"),
			key.WithHelp("esc/backspace", "go back"),
		),
	}),
	SwitchFocus: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch focus"),
		),
	}),

	Search: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "search"),
		),
	}),
	GoToTop: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("g"),
			key.WithHelp("g", "go to top"),
		),
	}),
	GoToBottom: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("G"),
			key.WithHelp("G", "go to bottom"),
		),
	}),

	ViewDDL: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "view DDL"),
		),
	}),
	ViewIndexs: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "view indexes"),
		),
	}),
	OpenPrompt: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys(":"),
			key.WithHelp(":", "open prompt"),
		),
	}),

	NextPage: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("f", "space", "pagedown"),
			key.WithHelp("f/space/pagedown", "next page"),
		),
	}),
	PrevPage: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("b", "pageup"),
			key.WithHelp("b/pageup", "previous page"),
		),
	}),

	Quit: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit"),
		),
	}),
	Help: opt.Set([]key.Binding{
		key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
	}),
}