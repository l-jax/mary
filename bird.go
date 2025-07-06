package main

const (
	edge                 = 10
	near                 = 30
	tooClose             = 2.5
	cohesionMultiplier   = 0.01
	separationMultiplier = 0.03
	alignmentMultiplier  = 0.01
	maxSpeed             = 2.0
)

type bird struct {
	char     rune
	position vector
	velocity vector
}

func (b *bird) move() {
	b.position.add(b.velocity)
}

func (b *bird) turn(others []*bird) {
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
	if b.position.x < edge {
		b.velocity.x += 1
	}
	if b.position.x > height-edge {
		b.velocity.x -= 1
	}
	if b.position.y < edge {
		b.velocity.y += 1
	}
	if b.position.y > width-edge {
		b.velocity.y -= 1
	}
}

func (b *bird) limitSpeed() {
	speed := b.velocity.distance(vector{})
	if speed > maxSpeed {
		b.velocity.divide(speed)
		b.velocity.multiply(maxSpeed)
	}
}
