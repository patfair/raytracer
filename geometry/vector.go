// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"fmt"
	"math"
)

// Represents a geometric object that has both a magnitude and a direction.
type Vector struct {
	X float64 // X-axis component
	Y float64 // Y-axis component
	Z float64 // Z-axis component
}

// Returns the magnitude of this vector.
func (vector Vector) Norm() float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)
}

// Returns a vector having the same direction as this one but with a magnitude of 1.
func (vector Vector) ToUnit() Vector {
	norm := vector.Norm()
	if norm == 0 || norm == 1 {
		return vector
	}
	return Vector{vector.X / norm, vector.Y / norm, vector.Z / norm}
}

// Returns a vector comprising the sum of this vector and the given other one.
func (vector Vector) Add(other Vector) Vector {
	return Vector{
		vector.X + other.X,
		vector.Y + other.Y,
		vector.Z + other.Z,
	}
}

// Returns a vector having the same direction as this one but whose magnitude is multiplied by the given factor.
func (vector Vector) Multiply(factor float64) Vector {
	return Vector{
		vector.X * factor,
		vector.Y * factor,
		vector.Z * factor,
	}
}

// Returns the scalar product of this vector and the given other one.
func (vector Vector) Dot(other Vector) float64 {
	return vector.X*other.X + vector.Y*other.Y + vector.Z*other.Z
}

// Returns the vector product of this vector and the given other one.
func (vector Vector) Cross(other Vector) Vector {
	return Vector{
		vector.Y*other.Z - vector.Z*other.Y,
		vector.Z*other.X - vector.X*other.Z,
		vector.X*other.Y - vector.Y*other.X,
	}
}

func (vector Vector) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", vector.X, vector.Y, vector.Z)
}
