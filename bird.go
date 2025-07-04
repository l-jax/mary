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

	if x < 0 {
		x = 0
		b.velocity.x = -b.velocity.x
	} else if x >= height {
		x = height - 1
		b.velocity.x = -b.velocity.x
	}

	if y < 0 {
		y = 0
		b.velocity.y = -b.velocity.y
	} else if y >= width {
		y = width - 1
		b.velocity.y = -b.velocity.y
	}

	b.position = vector{x: x, y: y}
}
