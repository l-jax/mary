package main

import "testing"

func TestMoveBirds(t *testing.T) {
	birds := [height][width]*bird{}
	birds[0][0] = &bird{char: 'A', dir: east}

	flock := flock{birds: birds}

	flock.moveBirds()

	if flock.birds[0][1] == nil {
		t.Errorf("Bird did not move: %v", flock.birds)
	}

}
