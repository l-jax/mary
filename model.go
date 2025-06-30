package main

import (
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	grid [][]bird
}

func newModel() model {
	text := []string{
		"Chunky and noisy,",
		"but with stars in their black feathers,",
		"they spring from the telephone wire",
		"and instantly",
	}

	grid := make([][]bird, len(text))
	for i := range grid {
		grid[i] = make([]bird, len(text[i]))
		for j := range grid[i] {
			grid[i][j] = bird{
				char:  rune(text[i][j]),
				x:     j,
				y:     i,
				speed: 1,
				dir:   direction(rand.Intn(7)),
			}
		}
	}

	return model{
		grid: grid,
	}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	var output string
	for _, row := range m.grid {
		for _, b := range row {
			if b.char == ' ' {
				output += " "
			} else {
				output += string(b.char)
			}
		}
		output += "\n"
	}
	return output
}
