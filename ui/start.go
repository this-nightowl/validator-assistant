package ui

import (
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

// Start launches the Bubble Tea program. If possible, it spawns a new
// terminal window so the assistant runs in its own container outside the
// current shell.
func Start() error {
	if os.Getenv("VA_CHILD") != "1" {
		exe, err := os.Executable()
		if err == nil {
			var cmd *exec.Cmd
			switch runtime.GOOS {
			case "darwin":
				cmd = exec.Command("open", "-a", "Terminal", exe)
			case "windows":
				cmd = exec.Command("cmd", "/c", "start", exe)
			default:
				cmd = exec.Command("x-terminal-emulator", "-e", exe)
			}
			cmd.Env = append(os.Environ(), "VA_CHILD=1")
			if err := cmd.Start(); err == nil {
				return nil
			}
		}
	}
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	return p.Start()
}
