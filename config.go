package main

/*
VIEWPORT
*/

const (
	width      = 90
	height     = 35
	topMargin  = 3
	sideMargin = 15
)

/*
BIRDS
*/
const (
	near          = 25
	tooClose      = 4.0
	maxSpeed      = 2.0
	maxMultiplier = 0.1
	minMultiplier = 0.001
	step          = 0.01
)

var (
	cohesionMultiplier   = 0.008
	separationMultiplier = 0.03
	alignmentMultiplier  = 0.01
)

/*
FLOCK
*/
const (
	ticksBetweenReleases = 1
)

/*
Starlings in Winter
Mary Oliver
*/
const text = `
Chunky and noisy,
but with stars in their black feathers,
they spring from the telephone wire
and instantly

they are acrobats
in the freezing wind.
And now, in the theater of air,
they swing over buildings,

dipping and rising;
they float like one stippled star
that opens,
becomes for a moment fragmented,

then closes again;
and you watch
and you try
but you simply can’t imagine

how they do it
with no articulated instruction, no pause,
only the silent confirmation
that they are this notable thing,

this wheel of many parts, that can rise and spin
over and over again,
full of gorgeous life.`
