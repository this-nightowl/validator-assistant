package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// mainMenuModel holds state for the main menu.
type mainMenuModel struct {
	choices []string
	cursor  int
}

func initialModel() mainMenuModel {
	return mainMenuModel{
		choices: []string{
			"Create New Validator",
			"Update Existing Validator",
			"Grafana Dashboard Setup",
			"Take or Restore Snapshot",
			"Troubleshooting & Reset",
			"Exit",
		},
	}
}

func (m mainMenuModel) Init() tea.Cmd {
	return nil
}

func (m mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.choices[m.cursor] == "Exit" {
				return m, tea.Quit
			}
			return m, tea.Quit // Placeholder for full flows
		}
	}
	return m, nil
}

var (
	logoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	titleStyle  = lipgloss.NewStyle().MarginBottom(1).Bold(true)
	itemStyle   = lipgloss.NewStyle().PaddingLeft(2)
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
)

func (m mainMenuModel) View() string {
	logo := logoStyle.Render(`
 _   _            _             
| \ | | ___  __ _| | _____ _ __ 
|  \| |/ _ \/ _` + "`" + ` | |/ / _ \ '__|
| |\  |  __/ (_| |   <  __/ |   
|_| \_|\___|\__,_|_|\_\___|_|   
`)
	s := logo + "\n" + titleStyle.Render("Nectar Validator Assistant") + "\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}
		s += fmt.Sprintf("%s %s\n", cursor, itemStyle.Render(choice))
	}
	s += "\nPress q to quit"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
