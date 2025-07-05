package main

const (
	alignmentDistance    = 4
	separationDistance   = 2
	maxSpeed             = 1
	cohesionMultiplier   = 0.03
	separationMultiplier = 0.03
	alignmentMultiplier  = 0.03
)

type bird struct {
	char     rune
	position vector
	velocity vector
}

func (b *bird) move() {
	position := b.position.add(b.velocity)

	if position.x < 0 {
		position.x = height - 1
	} else if position.x >= height {
		position.x = 0
	}

	if position.y < 0 {
		position.y = width - 1
	} else if position.y >= width {
		position.y = 0
	}

	b.position = position
}

func (b *bird) newThing(others []*bird) {
	var avgVelocity vector
	var avgPosition vector
	var avgSeparation vector
	nearbyCount := 0

	for _, other := range others {
		if other == b {
			continue
		}

		distance := b.position.distance(other.position)

		if distance < alignmentDistance {
			avgVelocity.add(other.velocity)
			avgPosition.add(other.position)

			if distance < separationDistance {
				diff := b.position.difference(other.position)
				avgSeparation.add(diff)
			}
			nearbyCount++
		}
	}

	if nearbyCount == 0 {
		return
	}

	avgVelocity.divide(float64(nearbyCount))
	avgPosition.divide(float64(nearbyCount))
	avgSeparation.divide(float64(nearbyCount))

	avgVelocity.multiply(alignmentMultiplier)
	avgPosition.multiply(cohesionMultiplier)
	avgSeparation.multiply(separationMultiplier)

	b.velocity.add(avgVelocity)
	b.velocity.add(avgPosition)
	b.velocity.subtract(avgSeparation)

	if b.velocity.x > maxSpeed {
		b.velocity.x = maxSpeed
	}
	if b.velocity.y > maxSpeed {
		b.velocity.y = maxSpeed
	}
	if b.velocity.x < -maxSpeed {
		b.velocity.x = -maxSpeed
	}
	if b.velocity.y < -maxSpeed {
		b.velocity.y = -maxSpeed
	}
}
