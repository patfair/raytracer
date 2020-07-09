// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBox(t *testing.T) {
	shadingProperties := shading.ShadingProperties{
		DiffuseTexture: shading.SolidTexture{shading.Color{0, 0.1, 0.2}},
		Reflectivity:   0.5,
		Opacity:        0.9,
	}
	planes, err := NewBox(geometry.Point{1, 2, 3}, geometry.Vector{2, 0, 0}, geometry.Vector{0, 1, 0}, -5,
		shadingProperties)
	assert.Nil(t, err)

	assert.Equal(t, geometry.Point{1, 2, 3}, planes[0].bottomLeftCorner)
	assert.Equal(t, geometry.Vector{2, 0, 0}, planes[0].width)
	assert.Equal(t, geometry.Vector{0, 1, 0}, planes[0].height)
	assert.Equal(t, shadingProperties, planes[0].shadingProperties)

	assert.Equal(t, geometry.Point{1, 2, 3}, planes[1].bottomLeftCorner)
	assert.Equal(t, geometry.Vector{0, 0, -5}, planes[1].width)
	assert.Equal(t, geometry.Vector{2, 0, 0}, planes[1].height)
	assert.Equal(t, shadingProperties, planes[1].shadingProperties)

	assert.Equal(t, geometry.Point{1, 2, 3}, planes[2].bottomLeftCorner)
	assert.Equal(t, geometry.Vector{0, 0, -5}, planes[2].width)
	assert.Equal(t, geometry.Vector{0, 1, 0}, planes[2].height)
	assert.Equal(t, shadingProperties, planes[2].shadingProperties)

	assert.Equal(t, geometry.Point{3, 3, -2}, planes[3].bottomLeftCorner)
	assert.Equal(t, geometry.Vector{-2, 0, 0}, planes[3].width)
	assert.Equal(t, geometry.Vector{0, -1, 0}, planes[3].height)
	assert.Equal(t, shadingProperties, planes[3].shadingProperties)

	assert.Equal(t, geometry.Point{3, 3, -2}, planes[4].bottomLeftCorner)
	assert.Equal(t, geometry.Vector{0, 0, 5}, planes[4].width)
	assert.Equal(t, geometry.Vector{-2, 0, 0}, planes[4].height)
	assert.Equal(t, shadingProperties, planes[4].shadingProperties)

	assert.Equal(t, geometry.Point{3, 3, -2}, planes[5].bottomLeftCorner)
	assert.Equal(t, geometry.Vector{0, 0, 5}, planes[5].width)
	assert.Equal(t, geometry.Vector{0, -1, 0}, planes[5].height)
	assert.Equal(t, shadingProperties, planes[5].shadingProperties)
}

func TestNewBoxInvalid(t *testing.T) {
	_, err := NewBox(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0}, 0,
		shading.ShadingProperties{})
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be non-zero")
	}

	_, err = NewBox(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{-1, 0, 0}, 1,
		shading.ShadingProperties{})
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be perpendicular")
	}
}
