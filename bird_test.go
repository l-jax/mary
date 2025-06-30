package main

import "testing"

var wrapTests = map[string]struct {
	x, y, wantX, wantY int
}{
	"negative x":            {x: -1, y: 0, wantX: width - 1, wantY: 0},
	"negative y":            {x: 0, y: -1, wantX: 0, wantY: height - 1},
	"x greater than width":  {x: width, y: 0, wantX: 0, wantY: 0},
	"y greater than height": {x: 0, y: height, wantX: 0, wantY: 0},
}

func TestWrap(t *testing.T) {
	for name, tc := range wrapTests {
		t.Run(name, func(t *testing.T) {
			bird := bird{char: 'o', x: tc.x, y: tc.y}

			bird.wrap()

			if bird.x != tc.wantX || bird.y != tc.wantY {
				t.Errorf("Expected bird to wrap to (%d, %d), got (%d, %d)", tc.wantX, tc.wantY, bird.x, bird.y)
			}
		})
	}
}

func TestCohesion(t *testing.T) {

}
