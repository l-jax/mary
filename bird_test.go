package main

import "testing"

var moveTests = map[string]struct {
	dir direction
	dx  int
	dy  int
}{
	"up":         {dir: north, dx: -1, dy: 0},
	"up-right":   {dir: northEast, dx: -1, dy: 1},
	"right":      {dir: east, dx: 0, dy: 1},
	"down-right": {dir: southEast, dx: 1, dy: 1},
	"down":       {dir: south, dx: 1, dy: 0},
	"down-left":  {dir: southWest, dx: 1, dy: -1},
	"left":       {dir: west, dx: 0, dy: -1},
	"up-left":    {dir: northWest, dx: -1, dy: -1},
}

func TestMove(t *testing.T) {
	for name, test := range moveTests {
		t.Run(name, func(t *testing.T) {
			bird := &bird{char: 'A', dir: test.dir, x: 1, y: 1}
			bird.move()
			if bird.x != 1+test.dx || bird.y != 1+test.dy {
				t.Errorf("Expected bird to move to (%d, %d), got (%d, %d)", 1+test.dx, 1+test.dy, bird.x, bird.y)
			}
		})
	}
}

var moveWrapTests = map[string]struct {
	dir    direction
	startX int
	startY int
	endX   int
	endY   int
}{
	"up":    {dir: north, startX: 0, startY: 0, endX: height - 1, endY: 0},
	"right": {dir: east, startX: 0, startY: width - 1, endX: 0, endY: 0},
	"down":  {dir: south, startX: height - 1, startY: 0, endX: 0, endY: 0},
	"left":  {dir: west, startX: 0, startY: 0, endX: 0, endY: width - 1},
}

func TestMoveWillWrap(t *testing.T) {
	for name, test := range moveWrapTests {
		t.Run(name, func(t *testing.T) {
			bird := &bird{char: 'A', dir: test.dir, x: test.startX, y: test.startY}
			bird.move()
			if bird.x != test.endX || bird.y != test.endY {
				t.Errorf("Expected bird to wrap to (%d, %d), got (%d, %d)", test.endX, test.endY, bird.x, bird.y)
			}
		})
	}
}

func TestNear(t *testing.T) {
	birdA := &bird{char: 'A', x: 0, y: 0}
	birdB := &bird{char: 'B', x: 0, y: nearbyDistance}
	birdC := &bird{char: 'C', x: 0, y: nearbyDistance + 1}

	if !birdA.isNear(birdB) {
		t.Errorf("Expected A to be near B")
	}

	if birdA.isNear(birdC) {
		t.Errorf("Expected A not to be near C")
	}
}

func TestNearWillWrap(t *testing.T) {
	birdA := &bird{char: 'A', x: 0, y: 0}
	birdB := &bird{char: 'B', x: 0, y: width - nearbyDistance}
	birdC := &bird{char: 'C', x: 0, y: width - nearbyDistance - 1}

	if !birdA.isNear(birdB) {
		t.Errorf("Expected A to be near B")
	}

	if birdA.isNear(birdC) {
		t.Errorf("Expected A not to be near C")
	}
}

func TestSteer(t *testing.T) {
	/*
		- A -
		B - C
	*/

	birdA := &bird{char: 'A', x: 0, y: 1, dir: east}
	birdB := &bird{char: 'B', x: 1, y: 0}
	birdC := &bird{char: 'C', x: 1, y: 2}

	birdA.steer([]*bird{birdB, birdC})

	if birdA.dir != south {
		t.Errorf("Expected A to steer south towards B and C, got %v", birdA.dir)
	}
}
