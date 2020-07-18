// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolidTexture_AlbedoAt(t *testing.T) {
	color := Color{0.3, 0.5, 0.7}
	texture := SolidTexture{color}
	assert.False(t, texture.NeedsTextureCoordinates())

	assertColorEqual(t, color, texture.AlbedoAt(0, 0), checkerboardDither)
	assertColorEqual(t, color, texture.AlbedoAt(-1, 5), checkerboardDither)
}
