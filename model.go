package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	width  = 80
	height = 30
)

var (
	borderColor = lipgloss.Color("205")
	birdColor   = lipgloss.Color("240")
	borderStyle = lipgloss.
			NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor)
	birdsStyle = lipgloss.
			NewStyle().
			Foreground(birdColor)
)

type tickMsg time.Time

func tick(t time.Duration) tea.Cmd {
	return tea.Every(t, func(t time.Time) tea.Msg { return tickMsg(t) })
}

type model struct {
	flock        flock
	tickInterval time.Duration
	started      bool
}

func newModel() model {
	return model{
		flock:        newFlock(),
		tickInterval: 100 * time.Millisecond,
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
		if msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == " " {
			m.started = !m.started
			if m.started {
				return m, tick(m.tickInterval)
			}
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
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
	return borderStyle.Render(output)
}
