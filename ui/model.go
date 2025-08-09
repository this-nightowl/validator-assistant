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
	tip    string
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

	width  int
	height int
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
			{prompt: "Moniker", tip: "Human-readable name for your validator. Example: MyValidator", set: func(c *config.ValidatorConfig, v string) { c.Moniker = v }},
			{prompt: "Identity", tip: "Optional identity signature such as a Keybase ID.", set: func(c *config.ValidatorConfig, v string) { c.Identity = v }},
			{prompt: "Website", tip: "Validator website URL.", set: func(c *config.ValidatorConfig, v string) { c.Website = v }},
			{prompt: "Details", tip: "Additional details about your validator.", set: func(c *config.ValidatorConfig, v string) { c.Details = v }},
			{prompt: "Chain ID", tip: "Network chain identifier, e.g. cosmoshub-4.", set: func(c *config.ValidatorConfig, v string) { c.ChainID = v }},
			{prompt: "External Address", tip: "Publicly reachable node address (host:port).", set: func(c *config.ValidatorConfig, v string) { c.Network.ExternalAddress = v }},
			{prompt: "Persistent Peers", tip: "Comma-separated node IDs for persistent connections.", set: func(c *config.ValidatorConfig, v string) { c.Network.PersistentPeers = v }},
			{prompt: "Seed Mode (true/false)", tip: "Run the node in seed mode?", set: func(c *config.ValidatorConfig, v string) {
				c.Network.SeedMode = strings.ToLower(v) == "true"
			}},
			{prompt: "Pex (true/false)", tip: "Enable peer exchange?", set: func(c *config.ValidatorConfig, v string) {
				c.Network.Pex = strings.ToLower(v) == "true"
			}},
			{prompt: "Addr Book Strict (true/false)", tip: "Strict address book policy.", set: func(c *config.ValidatorConfig, v string) {
				c.Network.AddrBookStrict = strings.ToLower(v) == "true"
			}},
			{prompt: "Validator Key", tip: "Path to validator private key file.", set: func(c *config.ValidatorConfig, v string) { c.Keys.ValidatorKey = v }},
			{prompt: "Keyring Backend", tip: "Keyring backend (os|file|test).", set: func(c *config.ValidatorConfig, v string) { c.Keys.KeyringBackend = v }},
			{prompt: "Key Name", tip: "Name of the key in keyring.", set: func(c *config.ValidatorConfig, v string) { c.Keys.KeyName = v }},
			{prompt: "DB Backend", tip: "Database backend (goleveldb|rocksdb|boltdb).", set: func(c *config.ValidatorConfig, v string) { c.Consensus.DBBackend = v }},
			{prompt: "Fast Sync (true/false)", tip: "Use fast blockchain sync?", set: func(c *config.ValidatorConfig, v string) {
				c.Consensus.FastSync = strings.ToLower(v) == "true"
			}},
			{prompt: "Log Level", tip: "Logging verbosity (info|debug|error).", set: func(c *config.ValidatorConfig, v string) { c.Consensus.LogLevel = v }},
			{prompt: "Commission Rate (%)", tip: "Percentage commission for delegators.", set: func(c *config.ValidatorConfig, v string) {
				f, _ := strconv.ParseFloat(v, 64)
				c.Economics.CommissionRate = f
			}},
			{prompt: "Min Self Delegation", tip: "Minimum self-bond required.", set: func(c *config.ValidatorConfig, v string) {
				i, _ := strconv.Atoi(v)
				c.Economics.MinSelfDelegation = i
			}},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
