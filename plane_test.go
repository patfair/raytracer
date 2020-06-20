package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPlaneIntersection(t *testing.T) {
	plane1 := Plane{Point{0, 0, 0}, Vector{0, 0, 1}}
	plane2 := Plane{Point{0, 0, 0}, Vector{0, 0, -1}}
	ray1 := Ray{Point{0, 0, 1.5}, Vector{0, 0, -1}}
	ray2 := Ray{Point{0, 0, 1.5}, Vector{0, 0, 1}}

	assert.Equal(t, 1.5, plane1.Intersection(ray1))
	assert.Equal(t, -1.5, plane1.Intersection(ray2))
	assert.Equal(t, 1.5, plane2.Intersection(ray1))
	assert.Equal(t, -1.5, plane2.Intersection(ray2))
}

func TestPlaneIntersectionParallel(t *testing.T) {
	plane1 := Plane{Point{1, 2, 3}, Vector{0, 0, 1}}
	plane2 := Plane{Point{-1, 5, 10}, Vector{1, 1, 1}}
	ray1 := Ray{Point{3, 20, -1.5}, Vector{0, 3, 0}}
	ray2 := Ray{Point{9, 0, 13}, Vector{-2, 1, 1}}

	assert.Equal(t, float64(math.MinInt64), plane1.Intersection(ray1))
	assert.NotEqual(t, float64(math.MinInt64), plane1.Intersection(ray2))

	assert.NotEqual(t, float64(math.MinInt64), plane2.Intersection(ray1))
	assert.Equal(t, float64(math.MinInt64), plane2.Intersection(ray2))
}
