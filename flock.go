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
				x: float64(offset + sideMargin),
				y: float64(y + topMargin),
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

func (f *flock) move(config FlockConfig) {
	for _, bird := range f.birds {
		bird.turn(f.birds, config)
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

	colorIdx := rand.Intn(4)
	f.birds[f.next].release(colorIdx)
	f.next++
	f.ticksSinceLastRelease = 0
}
