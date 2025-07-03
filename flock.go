package main

import "math/rand"

/*
Starlings in Winter
Mary Oliver
*/
var text = []string{
	"",
	"dipping and rising;",
	"they float like one stippled star",
	"that opens,",
	"becomes for a moment fragmented,",
	"",
	"then closes again;",
	"and you watch",
	"and you try",
	"but you simply can't imagine",
	"",
	"how they do it",
	"with no articulated instruction, no pause,",
	"only the silent confirmation",
	"that they are this notable thing,",
	"",
	"this wheel of many parts, that can rise and spin",
	"over and over again,",
	"full of gorgeous life.",
}

type flock struct {
	birds []*bird
}

func newFlock() flock {
	birds := make([]*bird, 0, 200)
	for i, line := range text {
		for j, char := range line {
			if char == ' ' || char == '\n' {
				continue
			}

			// randomly assign a direction
			dir := direction(rand.Intn(8))

			// create a new bird
			b := &bird{
				char: char,
				x:    i,
				y:    j,
				dir:  dir,
			}

			birds = append(birds, b)
		}
	}

	return flock{
		birds: birds,
	}
}

func (f *flock) move() {
	for _, bird := range f.birds {
		bird.move()
	}
}

func (f *flock) steer() {
	for i := range f.birds {
		nearbyBirds := make([]*bird, 0, 8)
		for j := range f.birds {
			if i == j {
				continue
			}
			if f.birds[i].isNear(f.birds[j]) {
				nearbyBirds = append(nearbyBirds, f.birds[j])
			}
		}

		if len(nearbyBirds) == 0 {
			continue
		}
		f.birds[i].steer(nearbyBirds)
	}
}
