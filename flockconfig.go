package main

type Preset int

const (
	Custom Preset = iota
	Calm
	Chaotic
	Swarm
	Cluster
)

var presetName = map[Preset]string{
	Custom:  "Custom",
	Calm:    "Calm",
	Chaotic: "Chaotic",
	Swarm:   "Swarm",
	Cluster: "Cluster",
}

func (p Preset) String() string {
	return presetName[p]
}

func (p Preset) GetConfig() FlockConfig {
	return presetConfig[p]
}

var presetConfig = []FlockConfig{
	{Custom, 0.1, 0.1, 0.1},
	{Calm, 0.05, 0.02, 0.04},
	{Chaotic, 0.01, 0.18, 0.01},
	{Swarm, 0.18, 0.05, 0.18},
	{Cluster, 0.15, 0.18, 0.02},
}

type FlockConfig struct {
	Name       Preset
	Cohesion   float64
	Separation float64
	Alignment  float64
}
