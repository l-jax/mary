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
	char  rune
	x, y  int
	speed int
	dir   direction
}
