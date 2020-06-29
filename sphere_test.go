package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSphereIntersection(t *testing.T) {
	// Intersecting from -X
	intersection := Sphere{Center: Point{2, 0, 0}, Radius: 3}.Intersection(Ray{Point{-4.5, 0, 0}, Vector{1, 0, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 3.5, intersection.Distance)
		assert.Equal(t, Point{-1, 0, 0}, intersection.Point)
		assert.Equal(t, Vector{-1, 0, 0}, intersection.Normal)
	}

	// Intersecting from +Y
	intersection = Sphere{Center: Point{0, 2, 0}, Radius: 3}.Intersection(Ray{Point{0, 7.5, 0}, Vector{0, -1, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 2.5, intersection.Distance)
		assert.Equal(t, Point{0, 5, 0}, intersection.Point)
		assert.Equal(t, Vector{0, 1, 0}, intersection.Normal)
	}

	// Tangent
	intersection = Sphere{Center: Point{0, 0, 0}, Radius: 5}.Intersection(Ray{Point{-1.5, 0, 5}, Vector{1, 0, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, Point{0, 0, 5}, intersection.Point)
		assert.Equal(t, Vector{0, 0, 1}, intersection.Normal)
	}

	// Intersecting behind ray
	intersection = Sphere{Center: Point{2, 0, 0}, Radius: 3}.Intersection(Ray{Point{6, 0, 0}, Vector{1, 0, 0}})
	assert.Nil(t, intersection)

	// Not intersecting
	intersection = Sphere{Center: Point{0, 0, 0}, Radius: 1}.Intersection(Ray{Point{0, 0, 2}, Vector{1, 0, 0}})
	assert.Nil(t, intersection)
}
