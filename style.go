package main

import "github.com/charmbracelet/lipgloss"

/*
   COLORS
*/

// Gradients for each mode (arrays from dark to light)
var (
	gradientCalm = []lipgloss.Color{
		lipgloss.Color("#264653"), // deep blue
		lipgloss.Color("#2978b5"),
		lipgloss.Color("#a8dadc"),
		lipgloss.Color("#f1faee"), // very light blue
	}
	gradientChaotic = []lipgloss.Color{
		lipgloss.Color("#e76f51"), // deep orange/red
		lipgloss.Color("#f4a261"),
		lipgloss.Color("#ffd6ba"),
		lipgloss.Color("#fff3e2"), // pale orange
	}
	gradientSwarm = []lipgloss.Color{
		lipgloss.Color("#264653"),
		lipgloss.Color("#2978b5"),
		lipgloss.Color("#88B04B"), // green accent for energy
		lipgloss.Color("#a8dadc"),
	}
	gradientCluster = []lipgloss.Color{
		lipgloss.Color("#264653"),
		lipgloss.Color("#6b8f71"), // muted green
		lipgloss.Color("#bfc8ad"), // stone grey-green
		lipgloss.Color("#f1faee"),
	}
	gradientCustom = []lipgloss.Color{
		lipgloss.Color("#b0b0b0"),
		lipgloss.Color("#d9cab3"),
		lipgloss.Color("#eaeaea"),
		lipgloss.Color("#f5f3f0"),
	}
)

func getBirdGradient(preset Preset) []lipgloss.Color {
	switch preset {
	case Calm:
		return gradientCalm
	case Chaotic:
		return gradientChaotic
	case Swarm:
		return gradientSwarm
	case Cluster:
		return gradientCluster
	default:
		return gradientCustom
	}
}

var (
	defaultBirdColor = lipgloss.Color("#b0b0b0")
	borderColor      = lipgloss.Color("#b0b0b0")
	helpColor        = lipgloss.Color("#b0b0b0")
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
