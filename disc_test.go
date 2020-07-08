package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDiscToTextureCoordinates(t *testing.T) {
	disc := Disc{
		plane: Plane{
			Corner: geometry.Point{1, 0, 5},
			Width:  geometry.Vector{2, 0, 0},
			Height: geometry.Vector{0, 1, 0},
			shadingProperties: shading.ShadingProperties{
				DiffuseTexture: shading.SolidTexture{shading.Color{0, 0, 0}},
				Opacity:        1,
			},
		},
	}

	r, phi := disc.toTextureCoordinates(geometry.Point{2, 0, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, 0.0, phi)

	r, phi = disc.toTextureCoordinates(geometry.Point{2, 1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, math.Pi/4, phi)

	r, phi = disc.toTextureCoordinates(geometry.Point{1, 1, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, math.Pi/2, phi)

	r, phi = disc.toTextureCoordinates(geometry.Point{0, 1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, 3*math.Pi/4, phi)

	r, phi = disc.toTextureCoordinates(geometry.Point{0, 0, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, math.Pi, phi)

	r, phi = disc.toTextureCoordinates(geometry.Point{0, -1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, -3*math.Pi/4, phi)

	r, phi = disc.toTextureCoordinates(geometry.Point{1, -1, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, -math.Pi/2, phi)

	r, phi = disc.toTextureCoordinates(geometry.Point{2, -1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, -math.Pi/4, phi)
}
