// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckerboardTexture_AlbedoAt(t *testing.T) {
	color1 := Color{1, 0, 0}
	color2 := Color{0, 0, 1}
	uPitch := 1.0
	vPitch := 2.0
	texture := CheckerboardTexture{color1, color2, uPitch, vPitch}
	ditherVariation := 0.02
	assert.True(t, texture.NeedsTextureCoordinates())

	assertColorEqual(t, color1, texture.AlbedoAt(0.1, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color1, texture.AlbedoAt(0.4, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color2, texture.AlbedoAt(0.6, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color2, texture.AlbedoAt(0.9, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color1, texture.AlbedoAt(1.1, 0.1, ditherVariation), ditherVariation)

	assertColorEqual(t, color2, texture.AlbedoAt(-0.1, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color2, texture.AlbedoAt(-0.4, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color1, texture.AlbedoAt(-0.6, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color1, texture.AlbedoAt(-0.9, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color2, texture.AlbedoAt(-1.1, 0.1, ditherVariation), ditherVariation)

	assertColorEqual(t, color1, texture.AlbedoAt(0.1, 0.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color1, texture.AlbedoAt(0.1, 0.9, ditherVariation), ditherVariation)
	assertColorEqual(t, color2, texture.AlbedoAt(0.1, 1.1, ditherVariation), ditherVariation)
	assertColorEqual(t, color2, texture.AlbedoAt(0.1, 1.9, ditherVariation), ditherVariation)
	assertColorEqual(t, color1, texture.AlbedoAt(0.1, 2.1, ditherVariation), ditherVariation)
}
