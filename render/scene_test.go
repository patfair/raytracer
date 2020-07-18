// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package render

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/light"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func TestScene(t *testing.T) {
	camera, err := NewCamera(geometry.Ray{geometry.Point{0, 0, 3}, geometry.Vector{0, 0, -1}}, geometry.Vector{0, 1, 0},
		90, 0.1, 5, 2, 2)
	assert.Nil(t, err)
	backgroundColor := shading.Color{0, 1, 0}

	scene := Scene{Camera: camera, BackgroundColor: backgroundColor}

	distantLight, err := light.NewDistantLight(geometry.Vector{0, 0, -1}, shading.Color{1, 1, 0}, 1, 0.1)
	assert.Nil(t, err)
	scene.AddLight(distantLight)

	shadingProperties := shading.ShadingProperties{
		DiffuseTexture:    shading.SolidTexture{Color: shading.Color{1, 1, 1}},
		SpecularExponent:  1,
		SpecularIntensity: 1,
		Opacity:           0.5,
		Reflectivity:      0.5,
		RefractiveIndex:   1.5,
	}
	plane, err := surface.NewPlane(geometry.Point{-1, -1, 0}, geometry.Vector{2, 0, 0}, geometry.Vector{0, 2, 0},
		shadingProperties)
	assert.Nil(t, err)
	scene.AddSurface(plane)

	_, err = scene.Render(RenderDraftPass, -1, 3)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be positive")
	}
	_, err = scene.Render(RenderDraftPass, 5, 0)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be positive")
	}

	image, err := scene.Render(RenderFinishPass, 16, 9)
	assert.Nil(t, err)
	assert.Equal(t, color.RGBA{0, 255, 0, 255}, image.RGBAAt(0, 0))
	assert.Equal(t, color.RGBA{0, 255, 0, 255}, image.RGBAAt(15, 0))
	assert.Equal(t, color.RGBA{0, 255, 0, 255}, image.RGBAAt(0, 8))
	assert.Equal(t, color.RGBA{0, 255, 0, 255}, image.RGBAAt(15, 8))
	assert.Equal(t, color.RGBA{255, 255, 0, 255}, image.RGBAAt(7, 4))
}
