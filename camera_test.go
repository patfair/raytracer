package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCameraXY(t *testing.T) {
	viewDirection := geometry.Ray{geometry.Point{-3, 2, -1}, geometry.Vector{0, 0, -1}}
	upDirection := geometry.Vector{0, 1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, err)
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{-0.5, 0.5, -1}.ToUnit()},
		camera.GetRay(0, 0, 0, 1, 0, 0))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{0.5, 0.5, -1}.ToUnit()},
		camera.GetRay(1, 0, 0, 1, 0, 0))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{-0.5, -0.5, -1}.ToUnit()},
		camera.GetRay(0, 1, 0, 1, 0, 0))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{0.5, -0.5, -1}.ToUnit()},
		camera.GetRay(1, 1, 0, 1, 0, 0))
}

func TestNewCameraRotated(t *testing.T) {
	viewDirection := geometry.Ray{geometry.Point{-5, -5, -5}, geometry.Vector{1, 0, 0}}
	upDirection := geometry.Vector{0, -1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, err)
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, -0.5, 0.5}.ToUnit()},
		camera.GetRay(0, 0, 0, 1, 0, 0))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, -0.5, -0.5}.ToUnit()},
		camera.GetRay(1, 0, 0, 1, 0, 0))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, 0.5, 0.5}.ToUnit()},
		camera.GetRay(0, 1, 0, 1, 0, 0))
	geometry.AssertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, 0.5, -0.5}.ToUnit()},
		camera.GetRay(1, 1, 0, 1, 0, 0))
}

func TestNewCameraVectorsNotPerpendicular(t *testing.T) {
	viewDirection := geometry.Ray{geometry.Point{-5, -5, -5}, geometry.Vector{1, 0, 0}}
	upDirection := geometry.Vector{1, -1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, camera)
	assert.NotNil(t, err)
}
