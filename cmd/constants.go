package cmd

import "github.com/charmbracelet/lipgloss"

var ERROR_MESSAGE_TEMPLATE = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
var SUCCESS_MESSAGE_TEMPLATE = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))

const VERSION = "1.0.0"
