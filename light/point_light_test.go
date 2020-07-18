// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package light

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewPointLight(t *testing.T) {
	light, err := NewPointLight(geometry.Point{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 254, 2.54)
	assert.Nil(t, err)

	assert.Equal(t, shading.Color{0.1, 0.2, 0.3}, light.Color())
	assert.Equal(t, 254.0/4/4/math.Pi, light.Intensity(geometry.Point{1, 2, 1}))
}

func TestNewPointLightInvalid(t *testing.T) {
	_, err := NewPointLight(geometry.Point{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, -1, 2.54)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be positive")
	}

	_, err = NewPointLight(geometry.Point{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 0, 2.54)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be positive")
	}

	_, err = NewPointLight(geometry.Point{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 1, -0.1)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be non-negative")
	}
}

func TestPointLight_Direction(t *testing.T) {
	light, _ := NewPointLight(geometry.Point{0, 0, 1}, shading.Color{1, 1, 1}, 1, 0)
	assert.Equal(t, geometry.Vector{0, 0, -1}, light.Direction(geometry.Point{0, 0, 0}, 0, 1))
	assert.Equal(t, geometry.Vector{0, 0, -1}, light.Direction(geometry.Point{0, 0, 0}, 1, 1))
	assert.Equal(t, geometry.Vector{0, 0, -1}, light.Direction(geometry.Point{0, 0, 0}, 2, 1))

	light, _ = NewPointLight(geometry.Point{0, 0, 1}, shading.Color{1, 1, 1}, 1, 0.1)
	direction := geometry.Vector{0, 0, -1}

	for i := 0; i < 10; i++ {
		newDirection := light.Direction(geometry.Point{0, 0, 0}, i, 10)
		assert.NotEqual(t, direction, newDirection)
		assert.Greater(t, newDirection.X, -0.1)
		assert.Less(t, newDirection.X, 0.1)
		assert.Greater(t, newDirection.Y, -0.1)
		assert.Less(t, newDirection.Y, 0.1)
		assert.InEpsilon(t, -1, newDirection.Z, 0.01)
	}

	// Check calculation edge cases
	light, _ = NewPointLight(geometry.Point{0, 0, 1}, shading.Color{1, 1, 1}, 1, 0.001)
	geometry.AssertVectorEqual(t, geometry.Vector{1, 0, 0}, light.Direction(geometry.Point{1, 0, 1}, 1, 2))
	geometry.AssertVectorEqual(t, geometry.Vector{0, -1, 0}, light.Direction(geometry.Point{0, -1, 1}, 2, 2))
	geometry.AssertVectorEqual(t, geometry.Vector{0, 0, 0}, light.Direction(geometry.Point{0, 0, 1}, 3, 2))
}

func TestPointLight_IsBlockedByIntersection(t *testing.T) {
	point := geometry.Point{1, -1, 0}
	light, _ := NewPointLight(geometry.Point{1, -1, 5}, shading.Color{1, 1, 1}, 1, 0)
	intersection1 := geometry.Intersection{
		Point:    geometry.Point{1, -1, 3},
		Distance: 3,
		Normal:   geometry.Vector{},
	}
	intersection2 := geometry.Intersection{
		Point:    geometry.Point{1, -1, 6},
		Distance: 6,
		Normal:   geometry.Vector{},
	}

	assert.True(t, light.IsBlockedByIntersection(point, &intersection1))
	assert.False(t, light.IsBlockedByIntersection(point, &intersection2))
}
