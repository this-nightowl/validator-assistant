package ui

import (
        "strconv"
        "strings"

        "github.com/charmbracelet/bubbles/progress"
        "github.com/charmbracelet/bubbles/textinput"
        tea "github.com/charmbracelet/bubbletea"
        "github.com/this-nightowl/validator-assistant/config"
)

type state int

const (
        stateMainMenu state = iota
        stateCreateValidator
        stateReviewValidator
)

type field struct {
        prompt   string
        tip      string
        required bool
        set      func(*config.ValidatorConfig, string)
}

type model struct {
	state   state
	choices []string
	cursor  int

	cfg    config.ValidatorConfig
	input  textinput.Model
	fields []field
	step   int
	// responses keeps the values entered for each field so we can
	// navigate backward and forward through the wizard.
	responses []string
	// function is the currently selected high level action, e.g. "Create New Validator".
	function string

        width    int
        height   int
        progress progress.Model
}

func initialModel() model {
        ti := newInput()
        ti.Placeholder = ""
        fields := []field{
                {prompt: "Moniker", tip: "Human-readable name for your validator. Example: MyValidator", required: true, set: func(c *config.ValidatorConfig, v string) { c.Moniker = v }},
                {prompt: "Identity", tip: "Optional identity signature such as a Keybase ID.", required: false, set: func(c *config.ValidatorConfig, v string) { c.Identity = v }},
                {prompt: "Website", tip: "Validator website URL.", required: false, set: func(c *config.ValidatorConfig, v string) { c.Website = v }},
                {prompt: "Details", tip: "Additional details about your validator.", required: false, set: func(c *config.ValidatorConfig, v string) { c.Details = v }},
                {prompt: "Chain ID", tip: "Network chain identifier, e.g. cosmoshub-4.", required: true, set: func(c *config.ValidatorConfig, v string) { c.ChainID = v }},
                {prompt: "External Address", tip: "Publicly reachable node address (host:port).", required: true, set: func(c *config.ValidatorConfig, v string) { c.Network.ExternalAddress = v }},
                {prompt: "Persistent Peers", tip: "Comma-separated node IDs for persistent connections.", required: false, set: func(c *config.ValidatorConfig, v string) { c.Network.PersistentPeers = v }},
                {prompt: "Seed Mode (true/false)", tip: "Run the node in seed mode?", required: false, set: func(c *config.ValidatorConfig, v string) {
                        c.Network.SeedMode = strings.ToLower(v) == "true"
                }},
                {prompt: "Pex (true/false)", tip: "Enable peer exchange?", required: false, set: func(c *config.ValidatorConfig, v string) {
                        c.Network.Pex = strings.ToLower(v) == "true"
                }},
                {prompt: "Addr Book Strict (true/false)", tip: "Strict address book policy.", required: false, set: func(c *config.ValidatorConfig, v string) {
                        c.Network.AddrBookStrict = strings.ToLower(v) == "true"
                }},
                {prompt: "Validator Key", tip: "Path to validator private key file.", required: true, set: func(c *config.ValidatorConfig, v string) { c.Keys.ValidatorKey = v }},
                {prompt: "Keyring Backend", tip: "Keyring backend (os|file|test).", required: true, set: func(c *config.ValidatorConfig, v string) { c.Keys.KeyringBackend = v }},
                {prompt: "Key Name", tip: "Name of the key in keyring.", required: true, set: func(c *config.ValidatorConfig, v string) { c.Keys.KeyName = v }},
                {prompt: "DB Backend", tip: "Database backend (goleveldb|rocksdb|boltdb).", required: false, set: func(c *config.ValidatorConfig, v string) { c.Consensus.DBBackend = v }},
                {prompt: "Fast Sync (true/false)", tip: "Use fast blockchain sync?", required: false, set: func(c *config.ValidatorConfig, v string) {
                        c.Consensus.FastSync = strings.ToLower(v) == "true"
                }},
                {prompt: "Log Level", tip: "Logging verbosity (info|debug|error).", required: false, set: func(c *config.ValidatorConfig, v string) { c.Consensus.LogLevel = v }},
                {prompt: "Commission Rate (%)", tip: "Percentage commission for delegators.", required: false, set: func(c *config.ValidatorConfig, v string) {
                        f, _ := strconv.ParseFloat(v, 64)
                        c.Economics.CommissionRate = f
                }},
                {prompt: "Min Self Delegation", tip: "Minimum self-bond required.", required: false, set: func(c *config.ValidatorConfig, v string) {
                        i, _ := strconv.Atoi(v)
                        c.Economics.MinSelfDelegation = i
                }},
        }
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
                input:     ti,
                fields:    fields,
                responses: make([]string, len(fields)),
                progress:  progress.New(progress.WithGradient("#004aad", "#ffb715"), progress.WithoutPercentage()),
        }
}

func (m model) Init() tea.Cmd {
	return nil
}
