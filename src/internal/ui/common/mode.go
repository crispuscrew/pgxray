package common

type UIMode int
const (
	Init 		UIMode = iota
	Critical

	Tree
	Data
	Prompt
)