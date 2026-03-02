// Package colors provides the color palette
package colors

import (
	"image/color"
	"os"

	"charm.land/lipgloss/v2"
)

type Palette struct {
	// Text
	TextPrimary 	color.Color
	TextMuted   	color.Color
	TextSubtle  	color.Color

	// Status
	Warning 		color.Color
	Error   		color.Color
	Success 		color.Color
	Info    		color.Color

	// UI chrome
	BorderActive   	color.Color
	BorderInactive 	color.Color
	Selection      	color.Color

	// Data value types
	Null    		color.Color
	Number  		color.Color
	String  		color.Color
	Boolean 		color.Color
	Date    		color.Color

	// Color for other purposes
	Accent 			color.Color
	Harmonic		color.Color
}

// Creates a palette adapted to the terminal's light/dark background.
func Default() Palette {
	hasDark := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
	ld := lipgloss.LightDark(hasDark)

	return Palette{
		TextPrimary: 	ld(lipgloss.Color("#1a1a1a"), lipgloss.Color("#f0f0f0")),
		TextMuted:   	ld(lipgloss.Color("#555555"), lipgloss.Color("#888888")),
		TextSubtle:  	ld(lipgloss.Color("#aaaaaa"), lipgloss.Color("#555555")),

		Warning: 		ld(lipgloss.Color("#cc6600"), lipgloss.Color("#ffb86c")),
		Error:   		ld(lipgloss.Color("#cc0000"), lipgloss.Color("#ff5555")),
		Success: 		ld(lipgloss.Color("#006600"), lipgloss.Color("#50fa7b")),
		Info:    		ld(lipgloss.Color("#0066cc"), lipgloss.Color("#8be9fd")),

		BorderActive:   ld(lipgloss.Color("#0066cc"), lipgloss.Color("#bd93f9")),
		BorderInactive: ld(lipgloss.Color("#cccccc"), lipgloss.Color("#44475a")),
		Selection:      ld(lipgloss.Color("#cce5ff"), lipgloss.Color("#44475a")),

		Null:    		ld(lipgloss.Color("#aaaaaa"), lipgloss.Color("#6272a4")),
		Number:  		ld(lipgloss.Color("#0055cc"), lipgloss.Color("#bd93f9")),
		String:  		ld(lipgloss.Color("#006600"), lipgloss.Color("#f1fa8c")),
		Boolean: 		ld(lipgloss.Color("#cc00cc"), lipgloss.Color("#ff79c6")),
		Date:    		ld(lipgloss.Color("#cc6600"), lipgloss.Color("#ffb86c")),

		Accent: 		ld(lipgloss.Color("#00FFA3"), lipgloss.Color("#B14FFF")),
		Harmonic:		ld(lipgloss.Color("#B14FFF"), lipgloss.Color("#00FFA3")),
	}
}
