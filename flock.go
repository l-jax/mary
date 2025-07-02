package main

import "math/rand"

var text = []string{
	"Chunky and noisy,",
	"but with stars in their black feathers,",
	"they spring from the telephone wire",
	"and instantly",
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
