package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	width  = 80
	height = 30
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
