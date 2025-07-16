package main

import (
	"math/rand"
	"strings"
)

type flock struct {
	birds                 []*bird
	next                  int
	ticksSinceLastRelease int
}

func newFlock() flock {
	birds := []*bird{}
	lines := strings.Split(text, "\n")
	for y, line := range lines {
		words := strings.Fields(line)
		offset := 0
		for _, word := range words {
			pos := vector{
				y: float64(y + topMargin),
				x: float64(offset + sideMargin),
			}
			b := newBird(word, pos)
			birds = append(birds, b)
			offset += len(word) + 1
		}
	}
	return flock{
		birds:                 birds,
		next:                  0,
		ticksSinceLastRelease: 0,
	}
}

func (f *flock) move() {
	for _, bird := range f.birds {
		bird.turn(f.birds)
		bird.move()
	}
}

func (f *flock) release() {
	if f.next >= len(f.birds) {
		return
	}

	f.ticksSinceLastRelease++
	if f.ticksSinceLastRelease < ticksBetweenReleases {
		return
	}

	color := birdColors[rand.Intn(len(birdColors))]
	f.birds[f.next].release(color)
	f.next++
	f.ticksSinceLastRelease = 0
}
