package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	switch m.state {
	case stateMainMenu:
		return m.updateMainMenu(msg)
	case stateCreateValidator:
		return m.updateCreateValidator(msg)
	}
	return m, nil
}

func (m model) updateMainMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				return m, textinput.Blink
			}
		}
	}
	return m, nil
}

func (m model) updateCreateValidator(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, nil
		case "shift+tab":
			if m.step > 0 {
				m.responses[m.step] = m.input.Value()
				m.fields[m.step].set(&m.cfg, m.input.Value())
				m.step--
				m.input.SetValue(m.responses[m.step])
				m.input.Placeholder = m.fields[m.step].prompt
			}
			return m, nil
		case "tab", "enter":
			m.responses[m.step] = m.input.Value()
			m.fields[m.step].set(&m.cfg, m.input.Value())
			if m.step < len(m.fields)-1 {
				m.step++
				m.input.SetValue(m.responses[m.step])
				m.input.Placeholder = m.fields[m.step].prompt
				return m, nil
			}
			m.state = stateMainMenu
			m.function = ""
			return m, nil
		}
	}
	return m, cmd
}
