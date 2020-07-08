// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolidTexture_AlbedoAt(t *testing.T) {
	color := Color{0.3, 0.5, 0.7}
	texture := SolidTexture{color}

	assert.Equal(t, color, texture.AlbedoAt(0, 0))
	assert.Equal(t, color, texture.AlbedoAt(-1, 5))
}
