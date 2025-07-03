package main

import "math"

const nearbyDistance = 1

type direction uint

const (
	north direction = iota
	northEast
	east
	southEast
	south
	southWest
	west
	northWest
)

type bird struct {
	char rune
	x, y int
	dir  direction
}

func (b *bird) isNear(other *bird) bool {
	dx := math.Abs(float64(b.x - other.x))
	dy := math.Abs(float64(b.y - other.y))

	return (dx <= nearbyDistance || dx >= height-nearbyDistance) &&
		(dy <= nearbyDistance || dy >= width-nearbyDistance)
}

func (b *bird) move() {
	x := b.x
	y := b.y
	switch b.dir {
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

	b.x, b.y = wrap(x, y)
}

func (b *bird) steer(neighbors []*bird) {
	count := len(neighbors)

	var x, y int
	for _, neighbor := range neighbors {
		x += neighbor.x
		y += neighbor.y
	}
	x /= count
	y /= count

	if x == 0 && y == 0 {
		return
	}

	if x < 0 && y < 0 {
		b.dir = northWest
	} else if x < 0 && y > 0 {
		b.dir = northEast
	} else if x > 0 && y < 0 {
		b.dir = southWest
	} else if x > 0 && y > 0 {
		b.dir = southEast
	} else if x < 0 {
		b.dir = north
	} else if x > 0 {
		b.dir = south
	} else if y < 0 {
		b.dir = west
	} else if y > 0 {
		b.dir = east
	}
}

func wrap(x, y int) (int, int) {
	if x < 0 {
		x = height - 1
	}
	if x >= height {
		x = 0
	}
	if y < 0 {
		y = width - 1
	}
	if y >= width {
		y = 0
	}
	return x, y
}
