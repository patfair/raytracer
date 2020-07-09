// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package light

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDistantLight(t *testing.T) {
	light, err := NewDistantLight(geometry.Vector{-1, -1, -1}, shading.Color{0.3, 0.2, 0.1}, 25.4, 0.1, 3)
	assert.Nil(t, err)

	assert.Equal(t, shading.Color{0.3, 0.2, 0.1}, light.Color())
	assert.Equal(t, 25.4, light.Intensity(geometry.Point{1, 2, 1}))
	assert.Equal(t, 3, light.NumSamples())
}

func TestNewDistantLightInvalid(t *testing.T) {
	_, err := NewDistantLight(geometry.Vector{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, -1, 2.54, 3)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be positive")
	}

	_, err = NewDistantLight(geometry.Vector{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 0, 2.54, 3)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be positive")
	}

	_, err = NewDistantLight(geometry.Vector{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 1, -0.1, 3)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be in [0, 1]")
	}

	_, err = NewDistantLight(geometry.Vector{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 1, 1.1, 3)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be in [0, 1]")
	}

	_, err = NewDistantLight(geometry.Vector{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 1, 0, 0)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be at least 1")
	}

	_, err = NewDistantLight(geometry.Vector{1, 2, 3}, shading.Color{0.1, 0.2, 0.3}, 1, 0, 3)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be non-zero if")
	}
}

func TestDistantLight_Direction(t *testing.T) {
	light, _ := NewDistantLight(geometry.Vector{0, 0, -1}, shading.Color{1, 1, 1}, 1, 0, 1)
	assert.Equal(t, geometry.Vector{0, 0, -1}, light.Direction(geometry.Point{0, 0, 0}, 0, 1))
	assert.Equal(t, geometry.Vector{0, 0, -1}, light.Direction(geometry.Point{0, 0, 0}, 1, 1))
	assert.Equal(t, geometry.Vector{0, 0, -1}, light.Direction(geometry.Point{0, 0, 0}, 2, 1))

	light, _ = NewDistantLight(geometry.Vector{0, 0, -1}, shading.Color{1, 1, 1}, 1, 0.1, 10)
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
}

func TestDistantLight_IsBlockedByIntersection(t *testing.T) {
	light, _ := NewDistantLight(geometry.Vector{-1, 2, 3}, shading.Color{1, 1, 1}, 1, 0, 1)
	intersection := geometry.Intersection{
		Point:    geometry.Point{1, 6, -13},
		Distance: 3,
		Normal:   geometry.Vector{},
	}

	assert.True(t, light.IsBlockedByIntersection(geometry.Point{1, 2, 3}, &intersection))
}
