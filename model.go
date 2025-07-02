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
	}

	return m, nil
}

func (m model) View() string {
	var output string
	for i := range height {
		for j := range width {
			birdMaybe := m.flock.birds[i][j]
			if birdMaybe == nil {
				output += " "
				continue
			}
			output += string(birdMaybe.char)
		}
		output += "\n"
	}
	return output
}
