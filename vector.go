package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (vector Vector) ToUnit() Vector {
	norm := vector.Norm()
	return Vector{vector.X / norm, vector.Y / norm, vector.Z / norm}
}

func (vector Vector) Norm() float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)
}

func (vector Vector) Add(other Vector) Vector {
	return Vector{
		vector.X + other.X,
		vector.Y + other.Y,
		vector.Z + other.Z,
	}
}

func (vector Vector) Multiply(factor float64) Vector {
	return Vector{
		vector.X * factor,
		vector.Y * factor,
		vector.Z * factor,
	}
}

func (vector Vector) Dot(other Vector) float64 {
	return vector.X*other.X + vector.Y*other.Y + vector.Z*other.Z
}

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
