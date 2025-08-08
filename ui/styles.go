package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Vibrant color palette for a more playful UI.
	logoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#F4B400")).Bold(true)
	titleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8C00")).Bold(true).MarginBottom(1)
	itemStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#E0E0E0")).PaddingLeft(2)
	cursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#F4B400"))
	paneStyle    = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#F4B400")).Padding(1)
	previewStyle = lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("#888888"))
)
