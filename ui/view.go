package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/this-nightowl/validator-assistant/utils"
)

func (m model) View() string {
	body := ""
	switch m.state {
	case stateMainMenu:
		body = m.mainMenuView()
	case stateCreateValidator:
		body = m.createValidatorView()
	}
	return lipgloss.JoinVertical(lipgloss.Left, m.headerView(), body, m.footerView())
}

func (m model) headerView() string {
	ascii := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, logoStyle.Render(utils.LoadLogo()))
	functions := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, strings.Join(m.choices, " | "))
	title := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, titleStyle.Render(m.currentTitle()))
	content := ascii + "\n" + functions + "\n" + title
	return headerStyle.Width(m.width).Render(content)
}

func (m model) footerView() string {
	help := ""
	switch m.state {
	case stateMainMenu:
		help = "[↑/↓] move  [Enter] select  [q] quit"
	case stateCreateValidator:
		help = "[Esc] back  [Enter] next  [Ctrl+C] quit"
	}
	return footerStyle.Width(m.width).Render(help)
}

func (m model) mainMenuView() string {
	var b strings.Builder
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}
		b.WriteString(fmt.Sprintf("%s %s\n", cursor, itemStyle.Render(choice)))
	}
	return lipgloss.PlaceHorizontal(m.width, lipgloss.Left, b.String())
}

func (m model) createValidatorView() string {
	prompt := promptStyle.Render(m.fields[m.step].prompt)
	tip := tipStyle.Render(m.fields[m.step].tip)
	left := fmt.Sprintf("%s\n%s\n\n%s", prompt, tip, m.input.View())
	preview := m.cfg.ToYAML()
	if preview == "" {
		preview = "# configuration preview"
	}
	leftPane := paneStyle.Width(m.width / 2).Render(left)
	rightPane := paneStyle.Width(m.width / 2).Render(previewStyle.Render(preview))
	return lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)
}

func (m model) currentTitle() string {
	switch m.state {
	case stateCreateValidator:
		return "Validator Creation Assistant"
	default:
		return "Main Menu"
	}
}
