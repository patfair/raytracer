package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointLightDirection(t *testing.T) {
	light := PointLight{point: Point{0, 0, 1}, radius: 0}
	assert.Equal(t, Vector{0, 0, -1}, light.Direction(Point{0, 0, 0}, 0))
	assert.Equal(t, Vector{0, 0, -1}, light.Direction(Point{0, 0, 0}, 1))
	assert.Equal(t, Vector{0, 0, -1}, light.Direction(Point{0, 0, 0}, 2))

	light = PointLight{point: Point{0, 0, 1}, radius: 0.1}
	direction := Vector{0, 0, -1}

	for i := 0; i < 10; i++ {
		newDirection := light.Direction(Point{0, 0, 0}, i)
		assert.NotEqual(t, direction, newDirection)
		assert.Greater(t, newDirection.X, -0.1)
		assert.Less(t, newDirection.X, 0.1)
		assert.Greater(t, newDirection.Y, -0.1)
		assert.Less(t, newDirection.Y, 0.1)
		assert.InEpsilon(t, -1, newDirection.Z, 0.01)
	}
}
