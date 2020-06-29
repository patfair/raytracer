package main

import (
	"math"
)

type Plane struct {
	Point
	Vector
	Color
}

func (plane Plane) Albedo() Color {
	return plane.Color
}

func (plane Plane) Intersection(ray Ray) (float64, Vector) {
	denominator := plane.ToUnit().Dot(ray.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return math.MinInt64, Vector{}
	}

	distance := ray.VectorTo(plane.Point).Dot(plane.Vector.ToUnit()) / denominator
	normal := plane.Vector
	if normal.Dot(ray.Vector) > 0 {
		normal = normal.Multiply(-1)
	}

	return distance, normal.ToUnit()
}
