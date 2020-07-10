// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"fmt"
)

// Represents a point in 3D Cartesian space.
type Point struct {
	X float64 // X-axis coordinate
	Y float64 // Y-axis coordinate
	Z float64 // Z-axis coordinate
}

// Returns the absolute distance between this point and the given other point.
func (point Point) DistanceTo(other Point) float64 {
	return point.VectorTo(other).Norm()
}

// Returns the point reached by translating this point by the given vector.
func (point Point) Translate(vector Vector) Point {
	return Point{
		point.X + vector.X,
		point.Y + vector.Y,
		point.Z + vector.Z,
	}
}

// Returns a vector that represents the translation from this point to the given other point.
func (point Point) VectorTo(other Point) Vector {
	return Vector{
		X: other.X - point.X,
		Y: other.Y - point.Y,
		Z: other.Z - point.Z,
	}
}

func (point Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", point.X, point.Y, point.Z)
}
