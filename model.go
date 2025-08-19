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

type model struct {
	flock        flock
	tickInterval time.Duration
	config       FlockConfig
	started      bool
	help         help.Model
}

func newModel() model {
	return model{
		flock:        newFlock(),
		tickInterval: millisBetweenTicks * time.Millisecond,
		started:      false,
		help:         help.New(),
		config:       presetConfig[0],
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
			m.flock.move(m.config)
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
			m.setCustomConfig(min(m.config.Cohesion+0.01, 0.2), m.config.Separation, m.config.Alignment)
		case key.Matches(msg, keys.CohesionDn):
			m.setCustomConfig(max(m.config.Cohesion-0.01, 0.0), m.config.Separation, m.config.Alignment)
		case key.Matches(msg, keys.SeparationUp):
			m.setCustomConfig(m.config.Cohesion, min(m.config.Separation+0.01, 0.2), m.config.Alignment)
		case key.Matches(msg, keys.SeparationDn):
			m.setCustomConfig(m.config.Cohesion, max(m.config.Separation-0.01, 0.0), m.config.Alignment)
		case key.Matches(msg, keys.AlignmentUp):
			m.setCustomConfig(m.config.Cohesion, m.config.Separation, min(m.config.Alignment+0.01, 0.2))
		case key.Matches(msg, keys.AlignmentDn):
			m.setCustomConfig(m.config.Cohesion, m.config.Separation, max(m.config.Alignment-0.01, 0.0))
		case key.Matches(msg, keys.Calm):
			m.applyPreset(Calm)
		case key.Matches(msg, keys.Chaotic):
			m.applyPreset(Chaotic)
		case key.Matches(msg, keys.Swarm):
			m.applyPreset(Swarm)
		case key.Matches(msg, keys.Cluster):
			m.applyPreset(Cluster)
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
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

func (m *model) setCustomConfig(cohesion, separation, alignment float64) {
	config := FlockConfig{
		Name:       Custom,
		Cohesion:   cohesion,
		Separation: separation,
		Alignment:  alignment,
	}

	m.config = config
}

func (m *model) applyPreset(preset Preset) {
	m.config = preset.GetConfig()
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		sliderStyle.Render(getSliders(m.config)),
		borderStyle.Render(getBirds(m.flock.birds, m.config.Name)),
		helpStyle.Render(m.help.View(keys)),
	)
}

func getBirds(birds []*bird, preset Preset) string {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	for _, bird := range birds {
		for _, letter := range bird.letters {
			x, y := int(letter.position.x), int(letter.position.y)
			if x >= 0 && x < width && y >= 0 && y < height {
				var color lipgloss.Color
				if bird.colorIdx < 0 {
					color = defaultBirdColor
				} else {
					color = getBirdGradient(preset)[bird.colorIdx]
				}
				style := lipgloss.NewStyle().Foreground(color)
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
	return output
}

func getSliders(config FlockConfig) string {
	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("Mode: %s", config.Name.String()),
		slider("Cohesion", config.Cohesion, 0.0, 0.2, 20),
		slider("Separation", config.Separation, 0.0, 0.2, 20),
		slider("Alignment", config.Alignment, 0.0, 0.2, 20),
	)
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
	return fmt.Sprintf("%-11s %s %.3f", label+":", bar, value)
}
