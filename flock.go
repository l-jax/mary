package main

import (
	"math/rand"
	"strings"
)

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
		y := 0
		words := strings.Split(line, " ")
		for j, word := range words {
			b := &bird{
				word: word,
				position: vector{
					x: float64(i) + height/8,
					y: float64(j+y) + width/4,
				},
				velocity: vector{
					x: rand.Float64()*2 - 1,
					y: rand.Float64()*2 - 1,
				},
			}
			y++
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
