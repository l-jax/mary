package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

func tick(t time.Duration) tea.Cmd {
	return tea.Every(t, func(t time.Time) tea.Msg { return tickMsg(t) })
}

type FlockPreset struct {
	Name       string
	Cohesion   float64
	Separation float64
	Alignment  float64
}

var presets = []FlockPreset{
	{"Calm", 0.05, 0.02, 0.04},
	{"Chaotic", 0.01, 0.18, 0.01},
	{"Swarm", 0.18, 0.05, 0.18},
	{"Cluster", 0.15, 0.18, 0.02},
}

type model struct {
	flock        flock
	tickInterval time.Duration
	started      bool
	help         help.Model
	cohesion     float64
	separation   float64
	alignment    float64
	presetIdx    int
}

func newModel() model {
	return model{
		flock:        newFlock(),
		tickInterval: millisBetweenTicks * time.Millisecond,
		started:      false,
		help:         help.New(),
		cohesion:     0.02,
		separation:   0.03,
		alignment:    0.01,
		presetIdx:    0,
	}

}

func (m model) Init() tea.Cmd {
	return tick(m.tickInterval)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		if m.started {
			m.flock.release()
			m.flock.move(m.cohesion, m.separation, m.alignment)
			return m, tick(m.tickInterval)
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Start):
			m.started = !m.started
			if !m.started {
				return m, nil
			}
			return m, tick(m.tickInterval)
		case key.Matches(msg, keys.CohesionUp):
			m.cohesion = min(m.cohesion+0.01, 0.2)
		case key.Matches(msg, keys.CohesionDn):
			m.cohesion = max(m.cohesion-0.01, 0.0)
		case key.Matches(msg, keys.SeparationUp):
			m.separation = min(m.separation+0.01, 0.2)
		case key.Matches(msg, keys.SeparationDn):
			m.separation = max(m.separation-0.01, 0.0)
		case key.Matches(msg, keys.AlignmentUp):
			m.alignment = min(m.alignment+0.01, 0.2)
		case key.Matches(msg, keys.AlignmentDn):
			m.alignment = max(m.alignment-0.01, 0.0)
		case msg.String() == "1":
			m.applyPreset(0)
		case msg.String() == "2":
			m.applyPreset(1)
		case msg.String() == "3":
			m.applyPreset(2)
		case msg.String() == "4":
			m.applyPreset(3)
		}
		if key.Matches(msg, keys.Quit) {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *model) applyPreset(idx int) {
	if idx >= 0 && idx < len(presets) {
		m.cohesion = presets[idx].Cohesion
		m.separation = presets[idx].Separation
		m.alignment = presets[idx].Alignment
		m.presetIdx = idx
	}
}

func (m model) View() string {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	for _, bird := range m.flock.birds {
		for _, letter := range bird.letters {
			x, y := int(letter.position.x), int(letter.position.y)
			if x < width && y < height {
				style := lipgloss.NewStyle().Foreground(bird.color)
				grid[y][x] = style.Render(string(letter.char))
			}
		}
	}

	var output string
	for _, row := range grid {
		for _, cell := range row {
			output += cell
		}
		output += "\n"
	}
	output = borderStyle.Render(output)

	sliders := lipgloss.JoinVertical(lipgloss.Left,
		slider("Cohesion", m.cohesion, 0.0, 0.2, 20),
		slider("Separation", m.separation, 0.0, 0.2, 20),
		slider("Alignment", m.alignment, 0.0, 0.2, 20),
	)

	presetLabel := fmt.Sprintf("Preset: %s", presets[m.presetIdx].Name)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		presetLabel,
		sliders,
		output,
		helpStyle.Render(m.help.View(keys)),
	)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func slider(label string, value, min, max float64, width int) string {
	pos := int((value - min) / (max - min) * float64(width-2))
	if pos < 0 {
		pos = 0
	}
	if pos > width-2 {
		pos = width - 2
	}
	bar := "[" + strings.Repeat("=", pos) + ">" + strings.Repeat(" ", width-2-pos) + "]"
	return lipgloss.NewStyle().
		Render(
			fmt.Sprintf("%-11s %s %.3f", label+":", bar, value),
		)
}
