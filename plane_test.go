package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaneIntersection(t *testing.T) {
	plane1 := Plane{Corner: Point{0, 0, 0}, Width: Vector{1, 0, 0}, Height: Vector{0, 1, 0}}
	plane2 := Plane{Corner: Point{0, 0, 0}, Width: Vector{0, 1, 0}, Height: Vector{1, 0, 0}}
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
	plane1 := Plane{Corner: Point{-50, -50, 0}, Width: Vector{100, 0, 0}, Height: Vector{0, 100, 0}}
	plane2 := Plane{Corner: Point{50, -50, -50}, Width: Vector{-100, 100, 0}, Height: Vector{-100, -100, 100}}
	ray1 := Ray{Point{0, 0, 0}, Vector{0, -3, 0}}
	ray2 := Ray{Point{0, 0, 0}, Vector{2, -1, -1}}

	intersection := plane1.Intersection(ray1)
	assert.Nil(t, intersection)

	intersection = plane1.Intersection(ray2)
	assert.NotNil(t, intersection)

	intersection = plane2.Intersection(ray1)
	assert.NotNil(t, intersection)

	intersection = plane2.Intersection(ray2)
	assert.Nil(t, intersection)
}
