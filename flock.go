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
