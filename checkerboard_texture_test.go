package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckerboardAlbedoAt(t *testing.T) {
	color1 := Color{1, 0, 0}
	color2 := Color{0, 0, 1}
	uPitch := 1.0
	vPitch := 2.0
	texture := CheckerboardTexture{color1, color2, uPitch, vPitch}

	assert.Equal(t, color1, texture.AlbedoAt(0.1, 0.1))
	assert.Equal(t, color1, texture.AlbedoAt(0.4, 0.1))
	assert.Equal(t, color2, texture.AlbedoAt(0.6, 0.1))
	assert.Equal(t, color2, texture.AlbedoAt(0.9, 0.1))
	assert.Equal(t, color1, texture.AlbedoAt(1.1, 0.1))

	assert.Equal(t, color2, texture.AlbedoAt(-0.1, 0.1))
	assert.Equal(t, color2, texture.AlbedoAt(-0.4, 0.1))
	assert.Equal(t, color1, texture.AlbedoAt(-0.6, 0.1))
	assert.Equal(t, color1, texture.AlbedoAt(-0.9, 0.1))
	assert.Equal(t, color2, texture.AlbedoAt(-1.1, 0.1))

	assert.Equal(t, color1, texture.AlbedoAt(0.1, 0.1))
	assert.Equal(t, color1, texture.AlbedoAt(0.1, 0.9))
	assert.Equal(t, color2, texture.AlbedoAt(0.1, 1.1))
	assert.Equal(t, color2, texture.AlbedoAt(0.1, 1.9))
	assert.Equal(t, color1, texture.AlbedoAt(0.1, 2.1))
}
