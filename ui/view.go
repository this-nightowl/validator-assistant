package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/this-nightowl/validator-assistant/utils"
)

func (m model) View() string {
	logo := logoStyle.Render(utils.LoadLogo())
	header := logo + "\n" + titleStyle.Render("Nectar Validator Assistant")

	switch m.state {
	case stateMainMenu:
		var b strings.Builder
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = cursorStyle.Render(">")
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, itemStyle.Render(choice)))
		}
		b.WriteString("\nPress q to quit")
		return header + "\n\n" + b.String()
	case stateCreateValidator:
		left := fmt.Sprintf("%s\n\n%s", m.fields[m.step].prompt, m.input.View())
		preview := m.cfg.ToYAML()
		if preview == "" {
			preview = "# configuration preview"
		}
		content := lipgloss.JoinHorizontal(lipgloss.Top,
			paneStyle.Render(left),
			paneStyle.Render(previewStyle.Render(preview)),
		)
		footer := "\n[Esc] back [Enter] next"
		return header + "\n" + content + footer
	}
	return header
}
