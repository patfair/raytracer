package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDiscToTextureCoordinates(t *testing.T) {
	disc := Disc{
		plane: Plane{
			Corner: Point{1, 0, 5},
			Width:  Vector{2, 0, 0},
			Height: Vector{0, 1, 0},
			shadingProperties: ShadingProperties{
				DiffuseTexture: SolidTexture{Color{0, 0, 0}},
				Opacity:        1,
			},
		},
	}

	r, phi := disc.toTextureCoordinates(Point{2, 0, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, 0.0, phi)

	r, phi = disc.toTextureCoordinates(Point{2, 1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, math.Pi/4, phi)

	r, phi = disc.toTextureCoordinates(Point{1, 1, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, math.Pi/2, phi)

	r, phi = disc.toTextureCoordinates(Point{0, 1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, 3*math.Pi/4, phi)

	r, phi = disc.toTextureCoordinates(Point{0, 0, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, math.Pi, phi)

	r, phi = disc.toTextureCoordinates(Point{0, -1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, -3*math.Pi/4, phi)

	r, phi = disc.toTextureCoordinates(Point{1, -1, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, -math.Pi/2, phi)

	r, phi = disc.toTextureCoordinates(Point{2, -1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, -math.Pi/4, phi)
}
