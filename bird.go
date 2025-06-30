package main

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

func (b *bird) move() {
	switch b.dir {
	case north:
		b.y--
	case northEast:
		b.x++
		b.y--
	case east:
		b.x++
	case southEast:
		b.x++
		b.y++
	case south:
		b.y++
	case southWest:
		b.x--
		b.y++
	case west:
		b.x--
	case northWest:
		b.x--
		b.y--
	}

	b.wrap()
}

func (b *bird) wrap() {
	if b.y < 0 {
		b.y = height - 1
	}
	if b.y >= height {
		b.y = 0
	}
	if b.x < 0 {
		b.x = width - 1
	}
	if b.x >= width {
		b.x = 0
	}
}
