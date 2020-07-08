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
