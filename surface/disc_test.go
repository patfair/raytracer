// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewDisc(t *testing.T) {
	shadingProperties := shading.ShadingProperties{
		DiffuseTexture: shading.SolidTexture{shading.Color{0, 0.1, 0.2}},
		Reflectivity:   0.5,
		Opacity:        0.9,
	}
	disc, err := NewDisc(geometry.Point{0, 0, 0}, geometry.Vector{0, 0, 1.5}, geometry.Vector{1.5, 0, 0},
		shadingProperties)
	assert.Nil(t, err)

	assert.Equal(t, shadingProperties, disc.ShadingProperties())
	assert.Equal(t, 1.5, disc.radius())
}

func TestNewDiscInvalid(t *testing.T) {
	_, err := NewDisc(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{1, 1, 1},
		shading.ShadingProperties{})
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be perpendicular")
	}

	_, err = NewDisc(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1.5, 0},
		shading.ShadingProperties{})
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must have the same magnitude")
	}
}

func TestDisc_Intersection(t *testing.T) {
	disc1, _ := NewDisc(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shading.ShadingProperties{})
	disc2, _ := NewDisc(geometry.Point{0, 0, 0}, geometry.Vector{0, 1, 0}, geometry.Vector{1, 0, 0},
		shading.ShadingProperties{})
	ray1 := geometry.Ray{geometry.Point{0, 1, 1.5}, geometry.Vector{0, 0, -1}}
	ray2 := geometry.Ray{geometry.Point{1, 0, 1.5}, geometry.Vector{0, 0, 1}}
	ray3 := geometry.Ray{geometry.Point{1.1, 0, 1}, geometry.Vector{0, 0, -1}}

	intersection := disc1.Intersection(ray1)
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 1, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 0, 1}, intersection.Normal)
	}

	// Intersecting behind ray
	intersection = disc1.Intersection(ray2)
	assert.Nil(t, intersection)

	intersection = disc2.Intersection(ray1)
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 1, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 0, 1}, intersection.Normal)
	}

	// Intersecting outside of radius
	intersection = disc2.Intersection(ray3)
	assert.Nil(t, intersection)
}

func TestDisc_IntersectionParallel(t *testing.T) {
	disc, _ := NewDisc(geometry.Point{-50, -50, 0}, geometry.Vector{100, 0, 0}, geometry.Vector{0, 100, 0},
		shading.ShadingProperties{})
	ray := geometry.Ray{geometry.Point{0, 0, 0}, geometry.Vector{0, -3, 0}}

	intersection := disc.Intersection(ray)
	assert.Nil(t, intersection)
}

func TestDisc_ToTextureCoordinates(t *testing.T) {
	disc, _ := NewDisc(geometry.Point{1, 0, 5}, geometry.Vector{2, 0, 0}, geometry.Vector{0, 2, 0},
		shading.ShadingProperties{})

	r, phi := disc.ToTextureCoordinates(geometry.Point{2, 0, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, 0.0, phi)

	r, phi = disc.ToTextureCoordinates(geometry.Point{2, 1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, math.Pi/4, phi)

	r, phi = disc.ToTextureCoordinates(geometry.Point{1, 1, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, math.Pi/2, phi)

	r, phi = disc.ToTextureCoordinates(geometry.Point{0, 1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, 3*math.Pi/4, phi)

	r, phi = disc.ToTextureCoordinates(geometry.Point{0, 0, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, math.Pi, phi)

	r, phi = disc.ToTextureCoordinates(geometry.Point{0, -1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, -3*math.Pi/4, phi)

	r, phi = disc.ToTextureCoordinates(geometry.Point{1, -1, 5})
	assert.Equal(t, 1.0, r)
	assert.Equal(t, -math.Pi/2, phi)

	r, phi = disc.ToTextureCoordinates(geometry.Point{2, -1, 5})
	assert.Equal(t, math.Sqrt(2), r)
	assert.Equal(t, -math.Pi/4, phi)
}
