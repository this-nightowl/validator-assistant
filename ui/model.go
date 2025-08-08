package ui

import (
	"strconv"
	"strings"

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
			{prompt: "External Address", set: func(c *config.ValidatorConfig, v string) { c.Network.ExternalAddress = v }},
			{prompt: "Persistent Peers", set: func(c *config.ValidatorConfig, v string) { c.Network.PersistentPeers = v }},
			{prompt: "Seed Mode (true/false)", set: func(c *config.ValidatorConfig, v string) {
				c.Network.SeedMode = strings.ToLower(v) == "true"
			}},
			{prompt: "Pex (true/false)", set: func(c *config.ValidatorConfig, v string) {
				c.Network.Pex = strings.ToLower(v) == "true"
			}},
			{prompt: "Addr Book Strict (true/false)", set: func(c *config.ValidatorConfig, v string) {
				c.Network.AddrBookStrict = strings.ToLower(v) == "true"
			}},
			{prompt: "Validator Key", set: func(c *config.ValidatorConfig, v string) { c.Keys.ValidatorKey = v }},
			{prompt: "Keyring Backend", set: func(c *config.ValidatorConfig, v string) { c.Keys.KeyringBackend = v }},
			{prompt: "Key Name", set: func(c *config.ValidatorConfig, v string) { c.Keys.KeyName = v }},
			{prompt: "DB Backend", set: func(c *config.ValidatorConfig, v string) { c.Consensus.DBBackend = v }},
			{prompt: "Fast Sync (true/false)", set: func(c *config.ValidatorConfig, v string) {
				c.Consensus.FastSync = strings.ToLower(v) == "true"
			}},
			{prompt: "Log Level", set: func(c *config.ValidatorConfig, v string) { c.Consensus.LogLevel = v }},
			{prompt: "Commission Rate (%)", set: func(c *config.ValidatorConfig, v string) {
				f, _ := strconv.ParseFloat(v, 64)
				c.Economics.CommissionRate = f
			}},
			{prompt: "Min Self Delegation", set: func(c *config.ValidatorConfig, v string) {
				i, _ := strconv.Atoi(v)
				c.Economics.MinSelfDelegation = i
			}},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
