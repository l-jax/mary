package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	width  = 50
	height = 20
)

type model struct {
	flock flock
}

func newModel() model {
	return model{
		flock: newFlock(),
	}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

		if msg.String() == "p" {
			m.flock.move()
			m.flock.steer()
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
				if bird.x == i && bird.y == j {
					output += string(bird.char)
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
	return output
}
