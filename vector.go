package main

import "math"

type vector struct {
	x, y float64
}

func (v *vector) add(other vector) {
	v.y += other.y
	v.x += other.x

}

func (v *vector) subtract(other vector) {
	v.y -= other.y
	v.x -= other.x
}

func (v *vector) divide(i float64) {
	if i == 0 {
		return
	}
	v.y /= i
	v.x /= i
}

func (v *vector) multiply(i float64) {
	if i == 0 {
		return
	}
	v.y *= i
	v.x *= i
}

func (v *vector) distance(other vector) float64 {
	return math.Sqrt(math.Pow(v.y-other.y, 2) + math.Pow(v.x-other.x, 2))
}

func (v *vector) difference(other vector) vector {
	return vector{y: v.y - other.y, x: v.x - other.x}
}
