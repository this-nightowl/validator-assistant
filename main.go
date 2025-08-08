package main

import (
	"fmt"
	"os"

	"github.com/this-nightowl/validator-assistant/ui"
)

func main() {
	if err := ui.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
