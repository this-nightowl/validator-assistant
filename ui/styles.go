package ui

import "github.com/charmbracelet/lipgloss"

var (
        logoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb715")).Bold(true)
        titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb715")).Bold(true)
        headerStyle = lipgloss.NewStyle().Background(lipgloss.Color("#004aad")).Foreground(lipgloss.Color("#F8F8F2")).Padding(1)
        footerStyle = lipgloss.NewStyle().Background(lipgloss.Color("#004aad")).Foreground(lipgloss.Color("#F8F8F2")).Padding(1)

        bodyStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#121212")).Foreground(lipgloss.Color("#F8F8F2"))
        itemStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).PaddingLeft(2)
        cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb715"))

        paneStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#1a1a1a")).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#004aad")).Padding(1)
        previewStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Padding(1)
        promptStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb715")).Bold(true)
        tipStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD")).Italic(true)
)
