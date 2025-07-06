package main

import "github.com/charmbracelet/bubbles/key"

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Start, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Start, k.CohesionUp, k.AlignmentUp, k.SeparationUp},
		{k.Quit, k.CohesionDown, k.AlignmentDown, k.SeparationDown},
	}
}

type keyMap struct {
	Start          key.Binding
	CohesionUp     key.Binding
	AlignmentUp    key.Binding
	SeparationUp   key.Binding
	CohesionDown   key.Binding
	AlignmentDown  key.Binding
	SeparationDown key.Binding
	Quit           key.Binding
}

var keys = keyMap{
	Start: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "start/stop"),
	),
	CohesionUp: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "increase cohesion"),
	),
	AlignmentUp: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "increase alignment"),
	),
	SeparationUp: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "increase separation"),
	),
	CohesionDown: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "decrease cohesion"),
	),
	AlignmentDown: key.NewBinding(
		key.WithKeys("z"),
		key.WithHelp("z", "decrease alignment"),
	),
	SeparationDown: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "decrease separation"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
