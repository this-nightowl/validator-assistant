package utils

import (
	"os"
	"path/filepath"
)

// LoadLogo reads the ASCII logo from the assets directory.
func LoadLogo() string {
	p := filepath.Join("assets", "ascii_logo.txt")
	b, err := os.ReadFile(p)
	if err != nil {
		return ""
	}
	return string(b)
}
