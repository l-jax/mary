package main

import "github.com/charmbracelet/lipgloss"

var gradients map[Mode][]lipgloss.Color = map[Mode][]lipgloss.Color{
	Calm: []lipgloss.Color{
		lipgloss.Color("#264653"),
		lipgloss.Color("#2978b5"),
		lipgloss.Color("#a8dadc"),
		lipgloss.Color("#f1faee"),
	},
	Chaotic: []lipgloss.Color{
		lipgloss.Color("#e76f51"),
		lipgloss.Color("#f4a261"),
		lipgloss.Color("#ffd6ba"),
		lipgloss.Color("#fff3e2"),
	},
	Swarm: []lipgloss.Color{
		lipgloss.Color("#264653"),
		lipgloss.Color("#2978b5"),
		lipgloss.Color("#88B04B"),
		lipgloss.Color("#a8dadc"),
	},
	Cluster: []lipgloss.Color{
		lipgloss.Color("#264653"),
		lipgloss.Color("#6b8f71"),
		lipgloss.Color("#bfc8ad"),
		lipgloss.Color("#f1faee"),
	},
	Custom: []lipgloss.Color{
		lipgloss.Color("#b0b0b0"),
		lipgloss.Color("#d9cab3"),
		lipgloss.Color("#eaeaea"),
		lipgloss.Color("#f5f3f0"),
	},
}

var (
	defaultBirdColor = lipgloss.Color("#b0b0b0")
	borderColor      = lipgloss.Color("#b0b0b0")
	helpColor        = lipgloss.Color("#b0b0b0")
)

var (
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor)
	helpStyle = lipgloss.NewStyle().
			Foreground(helpColor)
	sliderStyle = lipgloss.NewStyle().
			Foreground(helpColor)
)
