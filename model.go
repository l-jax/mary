package main

import (
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	width  = 40
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
			for i := range m.flock.birds {
				m.flock.birds[i].dir = direction((rand.Intn(8)))
				m.flock.birds[i].move()
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var output string
	for i := range height {
		for j := range width {
			birdFound := false
			for _, b := range m.flock.birds {
				if b.x == j && b.y == i {
					output += string(b.char)
					birdFound = true
					break
				}
			}
			if !birdFound {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}
