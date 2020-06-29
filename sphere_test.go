package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSphereIntersection(t *testing.T) {
	// Intersecting from -X
	distance, normal := Sphere{Point: Point{2, 0, 0}, Radius: 3}.Intersection(Ray{Point{-4.5, 0, 0}, Vector{1, 0, 0}})
	assert.Equal(t, 3.5, distance)
	assert.Equal(t, Vector{-1, 0, 0}, normal)

	// Intersecting from +Y
	distance, normal = Sphere{Point: Point{0, 2, 0}, Radius: 3}.Intersection(Ray{Point{0, 7.5, 0}, Vector{0, -1, 0}})
	assert.Equal(t, 2.5, distance)
	assert.Equal(t, Vector{0, 1, 0}, normal)

	// Not intersecting
	distance, normal = Sphere{Point: Point{0, 0, 0}, Radius: 5}.Intersection(Ray{Point{-1.5, 0, 5}, Vector{1, 0, 0}})
	assert.Equal(t, 1.5, distance)
	assert.Equal(t, Vector{0, 0, 1}, normal)

	// Intersecting behind ray
	distance, normal = Sphere{Point: Point{2, 0, 0}, Radius: 3}.Intersection(Ray{Point{6, 0, 0}, Vector{1, 0, 0}})
	assert.Equal(t, float64(math.MinInt64), distance)

	// Not intersecting
	distance, normal = Sphere{Point: Point{0, 0, 0}, Radius: 1}.Intersection(Ray{Point{0, 0, 2}, Vector{1, 0, 0}})
	assert.Equal(t, float64(math.MinInt64), distance)
}
