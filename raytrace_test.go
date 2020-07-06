package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArePixelsSimilar(t *testing.T) {
	size := 5
	pixels := make([][]Color, size)
	for i := 0; i < size; i++ {
		pixels[i] = make([]Color, size)
	}

	for i, row := range pixels {
		for j, _ := range row {
			assert.False(t, isSupersamplingRequired(pixels, j, i, 0))
			assert.False(t, isSupersamplingRequired(pixels, j, i, 1))
			assert.False(t, isSupersamplingRequired(pixels, j, i, 2))
			assert.False(t, isSupersamplingRequired(pixels, j, i, 3))
		}
	}

	pixels[2][2].R = 0.1
	assert.False(t, isSupersamplingRequired(pixels, 0, 0, 1))
	assert.True(t, isSupersamplingRequired(pixels, 0, 0, 2))
	assert.True(t, isSupersamplingRequired(pixels, 2, 2, 1))

	for i, row := range pixels {
		for j, _ := range row {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				assert.False(t, isSupersamplingRequired(pixels, j, i, 1))
			} else {
				assert.True(t, isSupersamplingRequired(pixels, j, i, 1))
			}
			assert.False(t, isSupersamplingRequired(pixels, j, i, 0))
			assert.True(t, isSupersamplingRequired(pixels, j, i, 2))
		}
	}
}
