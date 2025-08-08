package ui

import "github.com/charmbracelet/lipgloss"

var (
	logoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	titleStyle   = lipgloss.NewStyle().Bold(true).MarginBottom(1)
	itemStyle    = lipgloss.NewStyle().PaddingLeft(2)
	cursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	paneStyle    = lipgloss.NewStyle().Padding(1)
	previewStyle = lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("240"))
)
