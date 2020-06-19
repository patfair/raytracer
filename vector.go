package main

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (vector *Vector) Normalize() {
	norm := vector.Norm()
	vector.X /= norm
	vector.Y /= norm
	vector.Z /= norm
}

func (vector *Vector) Norm() float64 {
	return math.Sqrt(vector.X * vector.X + vector.Y * vector.Y + vector.Z * vector.Z)
}
