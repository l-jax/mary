package main

import "testing"

var moveTests = map[string]struct {
	start    vector
	end      vector
	velocity vector
}{
	"up":         {start: vector{x: 1, y: 1}, end: vector{x: 0, y: 1}, velocity: vector{x: -1, y: 0}},
	"down":       {start: vector{x: 1, y: 1}, end: vector{x: 2, y: 1}, velocity: vector{x: 1, y: 0}},
	"left":       {start: vector{x: 1, y: 1}, end: vector{x: 1, y: 0}, velocity: vector{x: 0, y: -1}},
	"right":      {start: vector{x: 1, y: 1}, end: vector{x: 1, y: 2}, velocity: vector{x: 0, y: 1}},
	"up-left":    {start: vector{x: 1, y: 1}, end: vector{x: 0, y: 0}, velocity: vector{x: -1, y: -1}},
	"up-right":   {start: vector{x: 1, y: 1}, end: vector{x: 0, y: 2}, velocity: vector{x: -1, y: 1}},
	"down-left":  {start: vector{x: 1, y: 1}, end: vector{x: 2, y: 0}, velocity: vector{x: 1, y: -1}},
	"down-right": {start: vector{x: 1, y: 1}, end: vector{x: 2, y: 2}, velocity: vector{x: 1, y: 1}},
}

func TestMove(t *testing.T) {
	for name, test := range moveTests {
		t.Run(name, func(t *testing.T) {
			bird := &bird{char: 'A', position: test.start, velocity: test.velocity}
			bird.move()
			if bird.position != test.end {
				t.Errorf("Expected bird to move to (%v), got (%v)", test.end, bird.position)
			}
		})
	}
}

var moveWrapTests = map[string]struct {
	position vector
	end      vector
	start    vector
}{
	"up":    {position: vector{0, 0}, start: vector{-1, 0}, end: vector{1, 0}},
	"down":  {position: vector{height - 1, 0}, start: vector{1, 0}, end: vector{-1, 0}},
	"left":  {position: vector{0, 0}, start: vector{0, -1}, end: vector{0, 1}},
	"right": {position: vector{0, width - 1}, start: vector{0, 1}, end: vector{0, -1}},
}

func TestMoveWillReverseAtEdges(t *testing.T) {
	for name, test := range moveWrapTests {
		t.Run(name, func(t *testing.T) {
			bird := &bird{char: 'A', position: test.position, velocity: test.start}
			bird.move()
			if bird.velocity != test.end {
				t.Errorf("Expected bird to reverse (%v), got (%v)", test.end, bird.velocity)
			}
		})
	}
}

func TestNear(t *testing.T) {
	birdA := &bird{char: 'A', position: vector{x: 0, y: 0}}
	birdB := &bird{char: 'B', position: vector{x: 0, y: nearbyDistance}}
	birdC := &bird{char: 'C', position: vector{x: 0, y: nearbyDistance + 1}}

	if !birdA.isNear(birdB) {
		t.Errorf("Expected A to be near B")
	}

	if birdA.isNear(birdC) {
		t.Errorf("Expected A not to be near C")
	}
}
