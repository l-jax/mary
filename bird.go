package main

import (
	"math/rand/v2"

	"github.com/charmbracelet/lipgloss"
)

type letter struct {
	char     rune
	position vector
}

type bird struct {
	word     string
	letters  []letter
	position vector
	velocity vector
	released bool
	color    lipgloss.Color
}

func newBird(word string, position vector) *bird {
	letters := make([]letter, len(word))
	for i, char := range word {
		letters[i] = letter{
			char:     char,
			position: vector{y: position.y, x: position.x + float64(i)},
		}
	}
	return &bird{
		word:     word,
		letters:  letters,
		position: position,
		velocity: vector{y: 0, x: 0},
		color:    softBlue,
	}
}

func (b *bird) release(color lipgloss.Color) {
	if b.released || len(b.letters) == 0 {
		return
	}
	b.color = color
	b.released = true
	b.velocity = vector{
		y: (rand.Float64() - 1) * 2,
		x: (rand.Float64()) * 2,
	}
}

func (b *bird) move() {
	if !b.released || len(b.letters) == 0 {
		return
	}

	b.position.add(b.velocity)

	lead := b.position
	for i := range b.letters {
		dir := lead.difference(b.letters[i].position)
		dir.multiply(0.5)
		b.letters[i].position.add(dir)
		lead = b.letters[i].position
	}
}

func (b *bird) turn(others []*bird) {
	if !b.released || len(others) == 0 {
		return
	}

	b.cohesion(others)
	b.separation(others)
	b.alignment(others)
	b.turnAwayFromEdge()
	b.limitSpeed()
}

func (b *bird) cohesion(others []*bird) {
	var sum vector
	count := 0
	for _, other := range others {
		if other == b {
			continue
		}

		distance := b.position.distance(other.position)

		if distance < near {
			sum.add(other.position)
			count++
		}
	}

	sum.divide(float64(count))
	sum.subtract(b.position)
	sum.multiply(cohesionMultiplier)
	b.velocity.add(sum)
}

func (b *bird) separation(others []*bird) {
	var sum vector
	for _, other := range others {
		if other == b {
			continue
		}

		distance := b.position.distance(other.position)

		if distance < tooClose {
			difference := b.position.difference(other.position)
			sum.add(difference)
		}
	}

	sum.multiply(separationMultiplier)
	b.velocity.add(sum)
}

func (b *bird) alignment(others []*bird) {
	var sum vector
	count := 0

	for _, other := range others {
		if other == b {
			continue
		}

		distance := b.position.distance(other.position)

		if distance < near {
			sum.add(other.velocity)
			count++
		}
	}

	if count == 0 {
		return
	}

	sum.divide(float64(count))
	difference := b.velocity.difference(sum)
	difference.multiply(alignmentMultiplier)
	b.velocity.add(difference)
}

func (b *bird) turnAwayFromEdge() {
	if b.position.y < topMargin {
		b.velocity.y += 1
	}
	if b.position.y > height-topMargin {
		b.velocity.y -= 1
	}
	if b.position.x < sideMargin {
		b.velocity.x += 1
	}
	if b.position.x > width-sideMargin {
		b.velocity.x -= 1
	}
}

func (b *bird) limitSpeed() {
	speed := b.velocity.distance(vector{})
	if speed > maxSpeed {
		b.velocity.divide(speed)
		b.velocity.multiply(maxSpeed)
	}
}
