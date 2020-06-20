package main

import (
	"math"
)

type Plane struct {
	Point
	Vector
}

func (plane Plane) Intersection(ray Ray) float64 {
	denominator := plane.ToUnit().Dot(ray.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return math.MinInt64
	}

	return ray.VectorTo(plane.Point).Dot(plane.ToUnit()) / denominator
}
