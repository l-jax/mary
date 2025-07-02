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
	birds [height][width]*bird
}

func newFlock() flock {
	birds := [height][width]*bird{}
	for i := range text {
		for j := range text[i] {
			if text[i][j] == ' ' {
				continue
			}
			birds[i][j] = &bird{
				char: rune(text[i][j]),
				dir:  direction(rand.Intn(8)),
			}
		}
	}
	return flock{
		birds: birds,
	}
}

func (f *flock) moveBirds() {
	birds := [height][width]*bird{}
	for i := range f.birds {
		for j := range f.birds[i] {
			bird := f.birds[i][j]

			if bird == nil {
				continue
			}

			x, y := findNewPosition(i, j, bird)
			birds[x][y] = bird
		}
	}
	f.birds = birds
}

func findNewPosition(i int, j int, bird *bird) (int, int) {
	x := i
	y := j
	switch bird.dir {
	case north:
		x--
	case northEast:
		x--
		y++
	case east:
		y++
	case southEast:
		x++
		y++
	case south:
		x++
	case southWest:
		x++
		y--
	case west:
		y--
	case northWest:
		x--
		y--
	}

	wrap(&x, &y)

	return x, y
}

func wrap(x *int, y *int) {
	if *x < 0 {
		*x = height - 1
	} else if *x >= height {
		*x = 0
	}

	if *y < 0 {
		*y = width - 1
	} else if *y >= width {
		*y = 0
	}
}
