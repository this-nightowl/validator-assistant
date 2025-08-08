package ui

import "github.com/charmbracelet/bubbles/textinput"

func newInput() textinput.Model {
	ti := textinput.New()
	ti.CharLimit = 64
	ti.Width = 30
	return ti
}
