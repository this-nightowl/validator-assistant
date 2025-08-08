# Nectar Validator Assistant

An interactive terminal user interface for configuring and managing validators on the **Nectar** network. Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and styled using [Lip Gloss](https://github.com/charmbracelet/lipgloss).

## Features
- 💡 Vibrant full-screen TUI with a stylized ASCII **Nectar** logo
- 🧭 Menu-driven wizard for creating a new validator
- 🧾 Live YAML preview of configuration as you type
- 🔧 Hooks for updating validators, Grafana setup, snapshots and reset tooling
- 🧱 Extensible architecture for future enhancements like automated deployment and dashboards

## Requirements
- Go 1.20+
- Git

## Running
```bash
go run .
```

## Navigation
- `↑/↓` move
- `Enter` select or confirm
- `Esc` go back
- `q` or `Ctrl+C` quit

The assistant guides you through filling out validator details. Configuration is previewed live and can later be used to automate deployment.

## Project Layout
```
validator-assistant/
├── main.go                 # Entry point
├── assets/                 # ASCII logos and other static assets
├── ui/                     # Bubble Tea models, views and styles
├── config/                 # Validator configuration structs and helpers
├── deploy/                 # Deployment logic (future)
└── utils/                  # Shared helpers
```

Contributions and feature requests are welcome!
