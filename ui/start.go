package ui

import tea "github.com/charmbracelet/bubbletea"

// Start launches the Bubble Tea program.
func Start() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	return p.Start()
}
