// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import (
	"github.com/stretchr/testify/assert"
	imagecolor "image/color"
	"testing"
)

func TestColor_ToRgba(t *testing.T) {
	assert.Equal(t, imagecolor.RGBA{0, 0, 0, 255}, Color{0, 0, 0}.ToRgba())
	assert.Equal(t, imagecolor.RGBA{255, 255, 255, 255}, Color{1, 1, 1}.ToRgba())
	assert.Equal(t, imagecolor.RGBA{63, 127, 191, 255}, Color{0.25, 0.5, 0.75}.ToRgba())
	assert.Equal(t, imagecolor.RGBA{0, 0, 0, 255}, Color{-0.1, -1, -100}.ToRgba())
	assert.Equal(t, imagecolor.RGBA{255, 255, 255, 255}, Color{1.01, 2, 50}.ToRgba())
}

func TestColor_Dither(t *testing.T) {
	color := Color{0.25, 0.5, 0.75}
	ditherColor := color.Dither(0.1)
	assertColorEqual(t, color, ditherColor, 0.1)
	assert.NotEqual(t, color.R, ditherColor.R)
	assert.NotEqual(t, color.R, ditherColor.G)
	assert.NotEqual(t, color.R, ditherColor.B)
}

// Asserts equality of the two given colors, within a small allowable error.
func assertColorEqual(t *testing.T, expected, actual Color, ditherVariance float64) {
	assert.InDelta(t, expected.R, actual.R, ditherVariance, "X expected: %v, actual: %v", expected.R, actual.R)
	assert.InDelta(t, expected.G, actual.G, ditherVariance, "Y expected: %v, actual: %v", expected.G, actual.G)
	assert.InDelta(t, expected.B, actual.B, ditherVariance, "Z expected: %v, actual: %v", expected.B, actual.B)
}
