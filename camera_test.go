package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCameraXY(t *testing.T) {
	viewDirection := Ray{Point{-3, 2, -1}, Vector{0, 0, -1}}
	upDirection := Vector{0, 1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, err)
	if assert.Equal(t, 2, len(camera.Rays)) && assert.Equal(t, 2, len(camera.Rays[0])) &&
		assert.Equal(t, 1, len(camera.Rays[0][0])) {
		assertRayEqual(t, Ray{viewDirection.Point, Vector{-0.5, 0.5, -1}.ToUnit()}, camera.Rays[0][0][0])
		assertRayEqual(t, Ray{viewDirection.Point, Vector{0.5, 0.5, -1}.ToUnit()}, camera.Rays[0][1][0])
		assertRayEqual(t, Ray{viewDirection.Point, Vector{-0.5, -0.5, -1}.ToUnit()}, camera.Rays[1][0][0])
		assertRayEqual(t, Ray{viewDirection.Point, Vector{0.5, -0.5, -1}.ToUnit()}, camera.Rays[1][1][0])
	}
}

func TestNewCameraRotated(t *testing.T) {
	viewDirection := Ray{Point{-5, -5, -5}, Vector{1, 0, 0}}
	upDirection := Vector{0, -1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, err)
	if assert.Equal(t, 2, len(camera.Rays)) && assert.Equal(t, 2, len(camera.Rays[0])) &&
		assert.Equal(t, 1, len(camera.Rays[0][0])) {
		assertRayEqual(t, Ray{viewDirection.Point, Vector{1, -0.5, 0.5}.ToUnit()}, camera.Rays[0][0][0])
		assertRayEqual(t, Ray{viewDirection.Point, Vector{1, -0.5, -0.5}.ToUnit()}, camera.Rays[0][1][0])
		assertRayEqual(t, Ray{viewDirection.Point, Vector{1, 0.5, 0.5}.ToUnit()}, camera.Rays[1][0][0])
		assertRayEqual(t, Ray{viewDirection.Point, Vector{1, 0.5, -0.5}.ToUnit()}, camera.Rays[1][1][0])
	}
}

func TestNewCameraVectorsNotPerpendicular(t *testing.T) {
	viewDirection := Ray{Point{-5, -5, -5}, Vector{1, 0, 0}}
	upDirection := Vector{1, -1, 0}
	camera, err := NewCamera(viewDirection, upDirection, 2, 2, 90, 0, 1, 1, 1)
	assert.Nil(t, camera)
	assert.NotNil(t, err)
}

func assertRayEqual(t *testing.T, expected, actual Ray) {
	assert.Equal(t, expected.Point, actual.Point)
	assertVectorEqual(t, expected.Direction, actual.Direction)
}

func assertVectorEqual(t *testing.T, expected, actual Vector) {
	epsilon := 0.001
	assert.InEpsilon(t, expected.X, actual.X, epsilon)
	assert.InEpsilon(t, expected.Y, actual.Y, epsilon)
	assert.InEpsilon(t, expected.Z, actual.Z, epsilon)
}
