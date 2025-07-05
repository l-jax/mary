package main

import "math"

type vector struct {
	x, y float64
}

func (v *vector) xInt() int {
	return int(v.x)
}

func (v *vector) yInt() int {
	return int(v.y)
}

func (v *vector) add(other vector) vector {
	return vector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v *vector) subtract(other vector) vector {
	return vector{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v *vector) divide(i float64) vector {
	return vector{
		x: v.x / i,
		y: v.y / i,
	}
}

func (v *vector) multiply(i float64) vector {
	return vector{
		x: v.x * i,
		y: v.y * i,
	}
}

func (v *vector) distance(other vector) float64 {
	return math.Sqrt(math.Pow(v.x-other.x, 2) + math.Pow(v.y-other.y, 2))
}

func (v *vector) difference(other vector) vector {
	return vector{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}
