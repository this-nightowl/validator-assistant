package ui

import "github.com/charmbracelet/lipgloss"

var (
	logoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true)
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")).Bold(true)
	headerStyle = lipgloss.NewStyle().Background(lipgloss.Color("#1E1E1E")).Foreground(lipgloss.Color("#F8F8F2")).Padding(1)
	footerStyle = lipgloss.NewStyle().Background(lipgloss.Color("#1E1E1E")).Foreground(lipgloss.Color("#F8F8F2")).Padding(1)

	itemStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).PaddingLeft(2)
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B"))

	paneStyle    = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#BD93F9")).Padding(1)
	previewStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Padding(1)
	promptStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true)
	tipStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD")).Italic(true)
)
