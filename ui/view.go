package ui

import (
        "fmt"
        "strings"

        "github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
        header := m.headerView()
        footer := m.footerView()
        progress := ""
        if m.state == stateCreateValidator || m.state == stateReviewValidator {
                progress = m.progressView()
        }

        body := ""
        switch m.state {
        case stateMainMenu:
                body = m.mainMenuView()
        case stateCreateValidator:
                body = m.createValidatorView()
        case stateReviewValidator:
                body = m.reviewValidatorView()
        }

        usedHeight := lipgloss.Height(header) + lipgloss.Height(footer) + lipgloss.Height(progress)
        bodyHeight := m.height - usedHeight
        if bodyHeight < 0 {
                bodyHeight = 0
        }
        body = bodyStyle.Width(m.width).Height(bodyHeight).Render(body)
        return lipgloss.JoinVertical(lipgloss.Left, header, body, progress, footer)
}

func (m model) headerView() string {
        var functionLine string
        if m.state != stateMainMenu && m.function != "" {
                functionLine = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, m.function)
        }
        title := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, titleStyle.Render(m.currentTitle()))
        content := title
        if functionLine != "" {
                content += "\n" + functionLine
        }
        return headerStyle.Width(m.width).Render(content)
}

func (m model) footerView() string {
        help := ""
        switch m.state {
        case stateMainMenu:
                help = "[↑/↓] move  [Enter] select  [q] quit"
        case stateCreateValidator:
                help = "[Esc] menu  [Shift+Tab] prev  [Tab/Enter] next  [Ctrl+C] quit"
        case stateReviewValidator:
                help = "[s] save  [d] deploy  [Esc] menu  [Ctrl+C] quit"
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
        pr := m.fields[m.step]
        promptText := pr.prompt
        if pr.required {
                promptText += " *"
        }
        prompt := promptStyle.Render(promptText)
        tipText := pr.tip
        if pr.required {
                tipText += " (required)"
        } else {
                tipText += " (optional)"
        }
        tip := tipStyle.Render(tipText)
        left := fmt.Sprintf("%s\n%s\n\n%s", prompt, tip, m.input.View())
        preview := m.cfg.ToYAML()
        if preview == "" {
                preview = "# configuration preview"
        }
        leftPane := paneStyle.Width(m.width / 2).Render(left)
        rightPane := paneStyle.Width(m.width / 2).Render(previewStyle.Copy().MaxHeight(10).Render(preview))
        return lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)
}

func (m model) reviewValidatorView() string {
        preview := m.cfg.ToYAML()
        if preview == "" {
                preview = "# configuration preview"
        }
        pane := paneStyle.Width(m.width).Render(previewStyle.Copy().MaxHeight(m.height-6).Render(preview))
        hint := tipStyle.Render("s: save  d: deploy  esc: menu")
        return lipgloss.JoinVertical(lipgloss.Left, pane, hint)
}

func (m model) currentTitle() string {
        switch m.state {
        case stateCreateValidator:
                return "Validator Creation Assistant"
        case stateReviewValidator:
                return "Review Configuration"
        default:
                return "Main Menu"
        }
}

func (m model) progressView() string {
        percent := 0.0
        if len(m.fields) > 0 {
                percent = float64(m.step) / float64(len(m.fields))
        }
        if percent > 1 {
                percent = 1
        }
        return m.progress.ViewAs(percent)
}
