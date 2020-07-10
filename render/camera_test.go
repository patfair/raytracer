// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package render

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCameraXY(t *testing.T) {
	viewDirection := geometry.Ray{geometry.Point{-3, 2, -1}, geometry.Vector{0, 0, -1}}
	upDirection := geometry.Vector{0, 1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 90, 0, 1, 1, 1)
	assert.Nil(t, err)
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{-0.5, 0.5, -1}.ToUnit()},
		camera.GetRay(2, 2, 0, 0, 0, 1, 0, 0, 1))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{0.5, 0.5, -1}.ToUnit()},
		camera.GetRay(2, 2, 1, 0, 0, 1, 0, 0, 1))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{-0.5, -0.5, -1}.ToUnit()},
		camera.GetRay(2, 2, 0, 1, 0, 1, 0, 0, 1))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{0.5, -0.5, -1}.ToUnit()},
		camera.GetRay(2, 2, 1, 1, 0, 1, 0, 0, 1))
}

func TestNewCameraRotated(t *testing.T) {
	viewDirection := geometry.Ray{geometry.Point{-5, -5, -5}, geometry.Vector{1, 0, 0}}
	upDirection := geometry.Vector{0, -1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 90, 0, 1, 1, 1)
	assert.Nil(t, err)
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, -0.5, 0.5}.ToUnit()},
		camera.GetRay(2, 2, 0, 0, 0, 1, 0, 0, 1))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, -0.5, -0.5}.ToUnit()},
		camera.GetRay(2, 2, 1, 0, 0, 1, 0, 0, 1))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, 0.5, 0.5}.ToUnit()},
		camera.GetRay(2, 2, 0, 1, 0, 1, 0, 0, 1))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, 0.5, -0.5}.ToUnit()},
		camera.GetRay(2, 2, 1, 1, 0, 1, 0, 0, 1))
}

func TestNewCameraInvalid(t *testing.T) {
	camera, err := NewCamera(geometry.Ray{geometry.Point{-5, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{1, -1, 0}, 90, 0, 1, 1, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "vectors must be perpendicular")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, -1, 0, 1, 1, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "view must be positive")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 0, 0, 1, 1, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "view must be positive")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 90, -0.1, 1, 1, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "radius must be non-negative")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 90, 1, 0, 1, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "distance must be positive")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 90, 1, -0.1, 1, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "distance must be positive")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 90, 1, 1, -1, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "field samples must be at least 1")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 90, 1, 1, 0, 1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "field samples must be at least 1")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 90, 1, 1, 1, -1)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "antialias samples must be at least 1")
	}

	camera, err = NewCamera(geometry.Ray{geometry.Point{1, -5, -5}, geometry.Vector{1, 0, 0}},
		geometry.Vector{0, -1, 0}, 90, 1, 1, 1, 0)
	assert.Nil(t, camera)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "antialias samples must be at least 1")
	}
}
