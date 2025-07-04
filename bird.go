package main

import (
	"math"
)

const nearbyDistance = 1

type vector struct {
	x, y int
}

type bird struct {
	char     rune
	position vector
	velocity vector
}

func (b *bird) isNear(other *bird) bool {
	distance := math.Sqrt(math.Pow(float64(b.position.x-b.position.y), 2) + math.Pow(float64(b.position.y-other.position.y), 2))
	return distance <= nearbyDistance
}

func (b *bird) move() {
	x := b.position.x + b.velocity.x
	y := b.position.y + b.velocity.y
	b.position = wrap(x, y)
}

func wrap(x, y int) vector {
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
	return vector{x: x, y: y}
}
