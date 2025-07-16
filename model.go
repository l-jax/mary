package main

import (
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
	started      bool
	help         help.Model
}

func newModel() model {
	return model{
		flock:        newFlock(),
		tickInterval: millisBetweenTicks * time.Millisecond,
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
