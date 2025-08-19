package main

import "github.com/charmbracelet/lipgloss"

/*
	COLORS
*/

// palette
var (
	navy     = lipgloss.Color("#03396c")
	deepBlue = lipgloss.Color("#005b96")
	blue     = lipgloss.Color("#6497b1")
	softBlue = lipgloss.Color("#b3cde0")
	lavender = lipgloss.Color("#cdb4db")
	pink     = lipgloss.Color("#ffc8dd")
	rose     = lipgloss.Color("#ffafcc")
	sky      = lipgloss.Color("#a2d2ff")
	mint     = lipgloss.Color("#b5ead7")
	blush    = lipgloss.Color("#f7cac9")
)

var birdColors = []lipgloss.Color{
	softBlue,
	blue,
	deepBlue,
	navy,
	lavender,
	pink,
	rose,
	sky,
	mint,
	blush,
}

var (
	borderColor = deepBlue
	helpColor   = softBlue
)

/*
	STYLES
*/

var (
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor)
	helpStyle = lipgloss.NewStyle().
			Foreground(helpColor)
	sliderStyle = lipgloss.NewStyle().
			Foreground(helpColor)
)
