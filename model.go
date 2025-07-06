package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	width  = 80
	height = 30
)

var (
	maxMultiplier        = 0.1
	minMultiplier        = 0.001
	cohesionMultiplier   = 0.01
	separationMultiplier = 0.03
	alignmentMultiplier  = 0.01
)

var (
	borderColor = lipgloss.Color("205")
	birdColor   = lipgloss.Color("240")
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor)
	birdsStyle = lipgloss.NewStyle().
			Foreground(birdColor)
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
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
		tickInterval: 100 * time.Millisecond,
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
	help := m.help.View(keys)

	var output string
	for i := range height {
		for j := range width {
			found := false
			for _, bird := range m.flock.birds {
				if bird.position.xInt() == i && bird.position.yInt() == j {
					output += birdsStyle.Render(string(bird.char))
					found = true
					break
				}
			}
			if !found {
				output += " "
			}
		}
		output += "\n"
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		borderStyle.Render(output),
		helpStyle.Render(help),
	)
}
