package config

import "gopkg.in/yaml.v3"

// ValidatorConfig defines the configuration for a Nectar validator.
type ValidatorConfig struct {
	Moniker     string          `yaml:"moniker"`
	Identity    string          `yaml:"identity,omitempty"`
	Website     string          `yaml:"website,omitempty"`
	Details     string          `yaml:"details,omitempty"`
	ChainID     string          `yaml:"chain_id"`
	NodeKey     string          `yaml:"node_key,omitempty"`
	GenesisFile string          `yaml:"genesis_file,omitempty"`
	Network     NetworkConfig   `yaml:"network"`
	Keys        KeyConfig       `yaml:"keys"`
	Consensus   ConsensusConfig `yaml:"consensus"`
	Economics   EconomicsConfig `yaml:"economics"`
}

// NetworkConfig describes network and p2p settings.
type NetworkConfig struct {
	ExternalAddress string `yaml:"external_address,omitempty"`
	PersistentPeers string `yaml:"persistent_peers,omitempty"`
	SeedMode        bool   `yaml:"seed_mode"`
	Pex             bool   `yaml:"pex"`
	AddrBookStrict  bool   `yaml:"addr_book_strict"`
}

// KeyConfig contains key and keyring information.
type KeyConfig struct {
	ValidatorKey   string `yaml:"validator_key"`
	KeyringBackend string `yaml:"keyring_backend"`
	KeyName        string `yaml:"key_name"`
	Mnemonic       string `yaml:"mnemonic,omitempty"`
}

// ConsensusConfig includes consensus and database settings.
type ConsensusConfig struct {
	DBBackend         string `yaml:"db_backend"`
	FastSync          bool   `yaml:"fast_sync"`
	LogLevel          string `yaml:"log_level"`
	TimeoutCommit     string `yaml:"timeout_commit"`
	SkipTimeoutCommit bool   `yaml:"skip_timeout_commit"`
}

// EconomicsConfig describes economics and staking parameters.
type EconomicsConfig struct {
	CommissionRate    float64 `yaml:"commission_rate"`
	MinSelfDelegation int     `yaml:"min_self_delegation"`
	MaxChangeRate     float64 `yaml:"max_change_rate"`
	DelegatorAddress  string  `yaml:"delegator_address,omitempty"`
}

// ToYAML renders the configuration as a YAML string.
func (c ValidatorConfig) ToYAML() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(b)
}
