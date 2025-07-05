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

			b := &bird{
				char: char,
				position: vector{
					x: float64(i),
					y: float64(j),
				},
				velocity: vector{
					x: rand.Float64()*2 - 1,
					y: rand.Float64()*2 - 1,
				},
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
		bird.turn(f.birds)
		bird.move()
	}
}
