package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

// Start launches the Bubble Tea program. If possible, it spawns a new
// terminal window so the assistant runs in its own container outside the
// current shell. If launching a new window isn't supported (for example, in a
// headless environment), it falls back to running in the current terminal.
func Start() error {
	if os.Getenv("VA_CHILD") != "1" {
		exe, err := os.Executable()
		if err == nil {
			var cmd *exec.Cmd
			switch runtime.GOOS {
			case "darwin":
				script := fmt.Sprintf(`tell application "Terminal" to do script "%s"`, exe)
				cmd = exec.Command("osascript", "-e", script)
			case "windows":
				cmd = exec.Command("cmd", "/c", "start", "", exe)
			default:
				if os.Getenv("DISPLAY") != "" {
					if _, err := exec.LookPath("x-terminal-emulator"); err == nil {
						cmd = exec.Command("x-terminal-emulator", "-e", exe)
					}
				}
			}
			if cmd != nil {
				cmd.Env = append(os.Environ(), "VA_CHILD=1")
				if err := cmd.Start(); err == nil {
					return nil
				}
			}
		}
	}
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	return p.Start()
}
