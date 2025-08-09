package ui

import (
        "os"
        "strings"

        "github.com/charmbracelet/bubbles/textinput"
        tea "github.com/charmbracelet/bubbletea"
        "github.com/charmbracelet/bubbles/progress"
        "github.com/this-nightowl/validator-assistant/config"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
        var cmds []tea.Cmd
        switch msg := msg.(type) {
        case tea.WindowSizeMsg:
                m.width = msg.Width
                m.height = msg.Height
                m.progress.Width = msg.Width
        }

        var cmd tea.Cmd
        var p tea.Model
        p, cmd = m.progress.Update(msg)
        if pm, ok := p.(progress.Model); ok {
                m.progress = pm
        }
        cmds = append(cmds, cmd)

        switch m.state {
        case stateMainMenu:
                m, cmd = m.updateMainMenu(msg)
        case stateCreateValidator:
                m, cmd = m.updateCreateValidator(msg)
        case stateReviewValidator:
                m, cmd = m.updateReviewValidator(msg)
        }
        if cmd != nil {
                cmds = append(cmds, cmd)
        }
        return m, tea.Batch(cmds...)
}

func (m model) updateMainMenu(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			choice := m.choices[m.cursor]
			if choice == "Exit" {
				return m, tea.Quit
			}
			if choice == "Create New Validator" {
				m.state = stateCreateValidator
				m.function = choice
				m.step = 0
				m.input.SetValue(m.responses[m.step])
				m.input.Placeholder = m.fields[m.step].prompt
				m.input.Focus()
				return m, textinput.Blink
			}
		}
	}
	return m, nil
}

func (m model) updateCreateValidator(msg tea.Msg) (model, tea.Cmd) {
        var cmd tea.Cmd
        m.input, cmd = m.input.Update(msg)
        switch msg := msg.(type) {
        case tea.KeyMsg:
                switch msg.String() {
                case "ctrl+c":
                        return m, tea.Quit
                case "esc":
                        m.state = stateMainMenu
                        m.function = ""
                        m.input.Blur()
                        return m, nil
                case "shift+tab":
                        if m.step > 0 {
                                m.responses[m.step] = m.input.Value()
                                m.fields[m.step].set(&m.cfg, m.input.Value())
                                m.step--
                                m.input.SetValue(m.responses[m.step])
                                m.input.Placeholder = m.fields[m.step].prompt
                                m.input.Focus()
                        }
                        return m, nil
                case "tab", "enter":
                        val := m.input.Value()
                        if m.fields[m.step].required && strings.TrimSpace(val) == "" {
                                return m, nil
                        }
                        m.responses[m.step] = val
                        m.fields[m.step].set(&m.cfg, val)
                        if m.step < len(m.fields)-1 {
                                m.step++
                                m.input.SetValue(m.responses[m.step])
                                m.input.Placeholder = m.fields[m.step].prompt
                                m.input.Focus()
                                return m, nil
                        }
                        m.step = len(m.fields)
                        m.state = stateReviewValidator
                        m.function = "Review Configuration"
                        m.input.Blur()
                        return m, nil
                }
        }
        return m, cmd
}

func (m model) updateReviewValidator(msg tea.Msg) (model, tea.Cmd) {
        switch msg := msg.(type) {
        case tea.KeyMsg:
                switch msg.String() {
                case "ctrl+c":
                        return m, tea.Quit
                case "esc":
                        m.state = stateMainMenu
                        m.function = ""
                        m.step = 0
                        return m, nil
                case "s":
                        saveConfig(m.cfg)
                        m.state = stateMainMenu
                        m.function = ""
                        m.step = 0
                        return m, nil
                case "d":
                        m.state = stateMainMenu
                        m.function = ""
                        m.step = 0
                        return m, nil
                }
        }
        return m, nil
}

func saveConfig(cfg config.ValidatorConfig) {
        _ = os.WriteFile("validator.yaml", []byte(cfg.ToYAML()), 0o644)
}
