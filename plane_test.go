package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaneIntersection(t *testing.T) {
	plane1 := Plane{Point: Point{0, 0, 0}, Normal: Vector{0, 0, 1}}
	plane2 := Plane{Point: Point{0, 0, 0}, Normal: Vector{0, 0, -1}}
	ray1 := Ray{Point{0, 1, 1.5}, Vector{0, 0, -1}}
	ray2 := Ray{Point{1, 0, 1.5}, Vector{0, 0, 1}}

	intersection := plane1.Intersection(ray1)
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, Point{0, 1, 0}, intersection.Point)
		assert.Equal(t, Vector{0, 0, 1}, intersection.Normal)
	}

	intersection = plane1.Intersection(ray2)
	assert.Nil(t, intersection)

	intersection = plane2.Intersection(ray1)
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, Point{0, 1, 0}, intersection.Point)
		assert.Equal(t, Vector{0, 0, 1}, intersection.Normal)
	}

	intersection = plane2.Intersection(ray2)
	assert.Nil(t, intersection)
}

func TestPlaneIntersectionParallel(t *testing.T) {
	plane1 := Plane{Point: Point{1, 2, 3}, Normal: Vector{0, 0, 1}}
	plane2 := Plane{Point: Point{-1, 5, 10}, Normal: Vector{1, 1, 1}}
	ray1 := Ray{Point{3, 20, -1.5}, Vector{0, -3, 0}}
	ray2 := Ray{Point{9, 0, 13}, Vector{2, -1, -1}}

	intersection := plane1.Intersection(ray1)
	assert.Nil(t, intersection)

	intersection = plane1.Intersection(ray2)
	assert.NotNil(t, intersection)

	intersection = plane2.Intersection(ray1)
	assert.NotNil(t, intersection)

	intersection = plane2.Intersection(ray2)
	assert.Nil(t, intersection)
}
