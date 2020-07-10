// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package render

import (
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArePixelsSimilar(t *testing.T) {
	var operation RaytraceRowOperation
	size := 5
	operation.Width = size
	operation.Height = size
	operation.RoughPassPixels = make([][]shading.Color, size)
	for i := 0; i < size; i++ {
		operation.RoughPassPixels[i] = make([]shading.Color, size)
	}

	for i := 0; i < size; i++ {
		operation.RowIndex = i
		for j := 0; j < size; j++ {
			assert.False(t, operation.isSupersamplingRequired(j, 0))
			assert.False(t, operation.isSupersamplingRequired(j, 1))
			assert.False(t, operation.isSupersamplingRequired(j, 2))
			assert.False(t, operation.isSupersamplingRequired(j, 3))
		}
	}

	operation.RoughPassPixels[2][2].R = 0.1
	operation.RowIndex = 0
	assert.False(t, operation.isSupersamplingRequired(0, 1))
	assert.True(t, operation.isSupersamplingRequired(0, 2))
	operation.RowIndex = 2
	assert.True(t, operation.isSupersamplingRequired(2, 1))

	for i := 0; i < size; i++ {
		operation.RowIndex = i
		for j := 0; j < size; j++ {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				assert.False(t, operation.isSupersamplingRequired(j, 1))
			} else {
				assert.True(t, operation.isSupersamplingRequired(j, 1))
			}
			assert.False(t, operation.isSupersamplingRequired(j, 0))
			assert.True(t, operation.isSupersamplingRequired(j, 2))
		}
	}
}
