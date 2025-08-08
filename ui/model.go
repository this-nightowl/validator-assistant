package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/this-nightowl/validator-assistant/config"
)

type state int

const (
	stateMainMenu state = iota
	stateCreateValidator
)

type field struct {
	prompt string
	set    func(*config.ValidatorConfig, string)
}

type model struct {
	state   state
	choices []string
	cursor  int

	cfg    config.ValidatorConfig
	input  textinput.Model
	fields []field
	step   int
}

func initialModel() model {
	ti := newInput()
	ti.Placeholder = ""
	return model{
		state: stateMainMenu,
		choices: []string{
			"Create New Validator",
			"Update Existing Validator",
			"Grafana Dashboard Setup",
			"Take or Restore Snapshot",
			"Troubleshooting & Reset",
			"Exit",
		},
		input: ti,
		fields: []field{
			{prompt: "Moniker", set: func(c *config.ValidatorConfig, v string) { c.Moniker = v }},
			{prompt: "Identity", set: func(c *config.ValidatorConfig, v string) { c.Identity = v }},
			{prompt: "Website", set: func(c *config.ValidatorConfig, v string) { c.Website = v }},
			{prompt: "Details", set: func(c *config.ValidatorConfig, v string) { c.Details = v }},
			{prompt: "Chain ID", set: func(c *config.ValidatorConfig, v string) { c.ChainID = v }},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
