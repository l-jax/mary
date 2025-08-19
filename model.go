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
	mode         ModeConfig
	started      bool
	help         help.Model
}

func newModel() model {
	return model{
		flock:        newFlock(),
		tickInterval: millisBetweenTicks * time.Millisecond,
		started:      false,
		help:         help.New(),
		mode:         modeConfig[0],
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
			m.flock.move(m.mode)
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
			m.setCustomMode(min(m.mode.Cohesion+0.01, 0.2), m.mode.Separation, m.mode.Alignment)
		case key.Matches(msg, keys.CohesionDn):
			m.setCustomMode(max(m.mode.Cohesion-0.01, 0.0), m.mode.Separation, m.mode.Alignment)
		case key.Matches(msg, keys.SeparationUp):
			m.setCustomMode(m.mode.Cohesion, min(m.mode.Separation+0.01, 0.2), m.mode.Alignment)
		case key.Matches(msg, keys.SeparationDn):
			m.setCustomMode(m.mode.Cohesion, max(m.mode.Separation-0.01, 0.0), m.mode.Alignment)
		case key.Matches(msg, keys.AlignmentUp):
			m.setCustomMode(m.mode.Cohesion, m.mode.Separation, min(m.mode.Alignment+0.01, 0.2))
		case key.Matches(msg, keys.AlignmentDn):
			m.setCustomMode(m.mode.Cohesion, m.mode.Separation, max(m.mode.Alignment-0.01, 0.0))
		case key.Matches(msg, keys.Calm):
			m.mode = Calm.GetConfig()
		case key.Matches(msg, keys.Chaotic):
			m.mode = Chaotic.GetConfig()
		case key.Matches(msg, keys.Swarm):
			m.mode = Swarm.GetConfig()
		case key.Matches(msg, keys.Cluster):
			m.mode = Cluster.GetConfig()
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

func (m *model) setCustomMode(cohesion, separation, alignment float64) {
	mode := ModeConfig{
		Name:       Custom,
		Cohesion:   cohesion,
		Separation: separation,
		Alignment:  alignment,
	}

	m.mode = mode
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		renderSliders(m.mode),
		renderBirds(m.flock.birds, m.mode.Name),
		helpStyle.Render(m.help.View(keys)),
	)
}

func renderBirds(birds []*bird, mode Mode) string {
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
				color := defaultBirdColor
				if bird.colorIdx >= 0 {
					color = gradients[mode][bird.colorIdx]
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
	return borderStyle.Render(output)
}

func renderSliders(config ModeConfig) string {
	sliders := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("Mode: %s", config.Name.String()),
		getSlider("Cohesion", config.Cohesion, 0.0, 0.2, 20),
		getSlider("Separation", config.Separation, 0.0, 0.2, 20),
		getSlider("Alignment", config.Alignment, 0.0, 0.2, 20),
	)
	return sliderStyle.Render(sliders)
}

func getSlider(label string, value, min, max float64, width int) string {
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
