package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	width  = 100
	height = 30
)

var (
	maxMultiplier        = 0.1
	minMultiplier        = 0.001
	cohesionMultiplier   = 0.008
	separationMultiplier = 0.03
	alignmentMultiplier  = 0.01
)

type tickMsg time.Time

func tick(t time.Duration) tea.Cmd {
	return tea.Every(t, func(t time.Time) tea.Msg { return tickMsg(t) })
}

type model struct {
	flock        flock
	tickInterval time.Duration
	started      bool
	help         help.Model
}

func newModel() model {
	return model{
		flock:        newFlock(),
		tickInterval: 150 * time.Millisecond,
		started:      false,
		help:         help.New(),
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
			m.flock.move()
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
		}
		if key.Matches(msg, keys.Quit) {
			return m, tea.Quit
		}
		if key.Matches(msg, keys.CohesionUp) {
			cohesionMultiplier += 0.01
			if cohesionMultiplier > maxMultiplier {
				cohesionMultiplier = maxMultiplier
			}
		}
		if key.Matches(msg, keys.CohesionDown) {
			cohesionMultiplier -= 0.01
			if cohesionMultiplier < minMultiplier {
				cohesionMultiplier = minMultiplier
			}
		}
		if key.Matches(msg, keys.SeparationUp) {
			separationMultiplier += 0.01
			if separationMultiplier > maxMultiplier {
				separationMultiplier = maxMultiplier
			}
		}
		if key.Matches(msg, keys.SeparationDown) {
			separationMultiplier -= 0.01
			if separationMultiplier < minMultiplier {
				separationMultiplier = minMultiplier
			}
		}
		if key.Matches(msg, keys.AlignmentUp) {
			alignmentMultiplier += 0.01
			if alignmentMultiplier > maxMultiplier {
				alignmentMultiplier = maxMultiplier
			}
		}
		if key.Matches(msg, keys.AlignmentDown) {
			alignmentMultiplier -= 0.01
			if alignmentMultiplier < minMultiplier {
				alignmentMultiplier = minMultiplier
			}
		}
	}
	return m, nil
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
			y, x := int(letter.position.x), int(letter.position.y)
			if x < height && y < width {
				style := lipgloss.NewStyle().Foreground(bird.color)
				grid[x][y] = style.Render(string(letter.char))
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
	return lipgloss.JoinVertical(
		lipgloss.Left,
		output,
		helpStyle.Render(m.help.View(keys)),
	)
}
