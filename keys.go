package main

import "github.com/charmbracelet/bubbles/key"

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Start, k.Calm, k.Chaotic, k.Swarm, k.Cluster, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Start, k.Calm, k.Chaotic, k.Swarm, k.Cluster, k.Quit},
	}
}

type keyMap struct {
	Start        key.Binding
	Quit         key.Binding
	CohesionUp   key.Binding
	CohesionDn   key.Binding
	SeparationUp key.Binding
	SeparationDn key.Binding
	AlignmentUp  key.Binding
	AlignmentDn  key.Binding
	Calm         key.Binding
	Chaotic      key.Binding
	Swarm        key.Binding
	Cluster      key.Binding
}

var keys = keyMap{
	Start: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "start/stop"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	CohesionUp: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "cohesion+"),
	),
	CohesionDn: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "cohesion-"),
	),
	SeparationUp: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "separation+"),
	),
	SeparationDn: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "separation-"),
	),
	AlignmentUp: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "alignment+"),
	),
	AlignmentDn: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "alignment-"),
	),
	Calm: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "calm"),
	),
	Chaotic: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "chaotic"),
	),
	Swarm: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "swarm"),
	),
	Cluster: key.NewBinding(
		key.WithKeys("4"),
		key.WithHelp("4", "cluster"),
	),
}
