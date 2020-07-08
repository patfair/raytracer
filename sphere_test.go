package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSphereIntersection(t *testing.T) {
	// Intersecting from -X
	intersection := Sphere{Center: geometry.Point{2, 0, 0}, Radius: 3}.Intersection(geometry.Ray{geometry.Point{-4.5, 0, 0}, geometry.Vector{1, 0, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 3.5, intersection.Distance)
		assert.Equal(t, geometry.Point{-1, 0, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{-1, 0, 0}, intersection.Normal)
	}

	// Intersecting from +Y
	intersection = Sphere{Center: geometry.Point{0, 2, 0}, Radius: 3}.Intersection(geometry.Ray{geometry.Point{0, 7.5, 0}, geometry.Vector{0, -1, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 2.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 5, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 1, 0}, intersection.Normal)
	}

	// Tangent
	intersection = Sphere{Center: geometry.Point{0, 0, 0}, Radius: 5}.Intersection(geometry.Ray{geometry.Point{-1.5, 0, 5}, geometry.Vector{1, 0, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 0, 5}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 0, 1}, intersection.Normal)
	}

	// Intersecting behind ray
	intersection = Sphere{Center: geometry.Point{2, 0, 0}, Radius: 3}.Intersection(geometry.Ray{geometry.Point{6, 0, 0}, geometry.Vector{1, 0, 0}})
	assert.Nil(t, intersection)

	// Not intersecting
	intersection = Sphere{Center: geometry.Point{0, 0, 0}, Radius: 1}.Intersection(geometry.Ray{geometry.Point{0, 0, 2}, geometry.Vector{1, 0, 0}})
	assert.Nil(t, intersection)
}

func TestSphereToTextureCoordinates(t *testing.T) {
	epsilon := 0.00001
	sphere := Sphere{Center: geometry.Point{0, 0, 0}, Radius: 1, ZenithReference: geometry.Vector{0, 0, 1},
		AzimuthReference: geometry.Vector{1, 0, 0}}

	theta, phi := sphere.toTextureCoordinates(geometry.Point{1, 0, 0})
	assert.Equal(t, 0.0, theta)
	assert.Equal(t, math.Pi/2, phi)

	theta, phi = sphere.toTextureCoordinates(geometry.Point{0, 0, 1})
	assert.Equal(t, 0.0, theta)
	assert.Equal(t, 0.0, phi)

	theta, phi = sphere.toTextureCoordinates(geometry.Point{0, 0, -1})
	assert.Equal(t, 0.0, theta)
	assert.Equal(t, math.Pi, phi)

	theta, phi = sphere.toTextureCoordinates(geometry.Point{0, -1, 1})
	assert.Equal(t, -math.Pi/2, theta)
	assert.InEpsilon(t, math.Pi/4, phi, epsilon)

	theta, phi = sphere.toTextureCoordinates(geometry.Point{-1, 0, -1})
	assert.Equal(t, math.Pi, theta)
	assert.Equal(t, 3*math.Pi/4, phi)
}
