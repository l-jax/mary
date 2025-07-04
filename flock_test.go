package main

import "testing"

func TestMoveBirds(t *testing.T) {
	birds := []*bird{
		{char: 'A', position: vector{x: 0, y: 0}, velocity: vector{x: 0, y: 1}},
	}

	flock := flock{birds: birds}

	flock.move()

	if flock.birds[0].position.x != 0 || flock.birds[0].position.y != 1 {
		t.Errorf("Expected bird to move to (0, 1), got (%d, %d)", flock.birds[0].position.x, flock.birds[0].position.y)
	}
}
