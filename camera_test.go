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
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{-0.5, 0.5, -1}.ToUnit()}, camera.GetRay(0, 0, 0, 1, 0, 0))
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{0.5, 0.5, -1}.ToUnit()}, camera.GetRay(1, 0, 0, 1, 0, 0))
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{-0.5, -0.5, -1}.ToUnit()}, camera.GetRay(0, 1, 0, 1, 0, 0))
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{0.5, -0.5, -1}.ToUnit()}, camera.GetRay(1, 1, 0, 1, 0, 0))
}

func TestNewCameraRotated(t *testing.T) {
	viewDirection := geometry.Ray{geometry.Point{-5, -5, -5}, geometry.Vector{1, 0, 0}}
	upDirection := geometry.Vector{0, -1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, err)
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, -0.5, 0.5}.ToUnit()}, camera.GetRay(0, 0, 0, 1, 0, 0))
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, -0.5, -0.5}.ToUnit()}, camera.GetRay(1, 0, 0, 1, 0, 0))
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, 0.5, 0.5}.ToUnit()}, camera.GetRay(0, 1, 0, 1, 0, 0))
	assertRayEqual(t, geometry.Ray{viewDirection.Origin, geometry.Vector{1, 0.5, -0.5}.ToUnit()}, camera.GetRay(1, 1, 0, 1, 0, 0))
}

func TestNewCameraVectorsNotPerpendicular(t *testing.T) {
	viewDirection := geometry.Ray{geometry.Point{-5, -5, -5}, geometry.Vector{1, 0, 0}}
	upDirection := geometry.Vector{1, -1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, camera)
	assert.NotNil(t, err)
}

func assertRayEqual(t *testing.T, expected, actual geometry.Ray) {
	assert.Equal(t, expected.Origin, actual.Origin)
	assertVectorEqual(t, expected.Direction, actual.Direction)
}

func assertVectorEqual(t *testing.T, expected, actual geometry.Vector) {
	epsilon := 0.001
	assert.InEpsilon(t, expected.X, actual.X, epsilon, "X expected: %v, actual: %v", expected.X, actual.X)
	assert.InEpsilon(t, expected.Y, actual.Y, epsilon, "Y expected: %v, actual: %v", expected.Y, actual.Y)
	assert.InEpsilon(t, expected.Z, actual.Z, epsilon, "Z expected: %v, actual: %v", expected.Z, actual.Z)
}
