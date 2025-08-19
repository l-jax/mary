package main

type Mode int

const (
	Custom Mode = iota
	Calm
	Chaotic
	Swarm
	Cluster
)

var modeName = map[Mode]string{
	Custom:  "Custom",
	Calm:    "Calm",
	Chaotic: "Chaotic",
	Swarm:   "Swarm",
	Cluster: "Cluster",
}

var modeConfig = []ModeConfig{
	{Custom, 0.1, 0.1, 0.1},
	{Calm, 0.05, 0.02, 0.04},
	{Chaotic, 0.01, 0.18, 0.01},
	{Swarm, 0.18, 0.05, 0.18},
	{Cluster, 0.15, 0.18, 0.02},
}

func (m Mode) Index() uint {
	return uint(m)
}

func (m Mode) String() string {
	return modeName[m]
}

func (m Mode) GetConfig() ModeConfig {
	return modeConfig[m]
}

type ModeConfig struct {
	Name       Mode
	Cohesion   float64
	Separation float64
	Alignment  float64
}
