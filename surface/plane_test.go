// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlane(t *testing.T) {
	shadingProperties := shading.ShadingProperties{
		DiffuseTexture:  shading.SolidTexture{shading.Color{0, 0.1, 0.2}},
		Reflectivity:    0.5,
		Opacity:         0.9,
		RefractiveIndex: 1.1,
	}
	plane, err := NewPlane(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shadingProperties)
	assert.Nil(t, err)

	assert.Equal(t, shadingProperties, plane.ShadingProperties())
	assert.Equal(t, geometry.Vector{0, 0, 1}, plane.normal)
}

func TestNewPlaneInvalid(t *testing.T) {
	_, err := NewPlane(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{1, 1, 1},
		shading.ShadingProperties{Opacity: 1})
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "must be perpendicular")
	}

	_, err = NewPlane(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 1},
		shading.ShadingProperties{SpecularExponent: -1})
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "exponent must be non-negative")
	}
}

func TestPlane_Intersection(t *testing.T) {
	plane1, _ := NewPlane(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shading.ShadingProperties{Opacity: 1})
	plane2, _ := NewPlane(geometry.Point{0, 0, 0}, geometry.Vector{0, 1, 0}, geometry.Vector{1, 0, 0},
		shading.ShadingProperties{Opacity: 1})
	ray1 := geometry.Ray{geometry.Point{0, 1, 1.5}, geometry.Vector{0, 0, -1}}
	ray2 := geometry.Ray{geometry.Point{1, 0, 1.5}, geometry.Vector{0, 0, 1}}

	intersection := plane1.Intersection(ray1)
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 1, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 0, 1}, intersection.Normal)
	}

	intersection = plane1.Intersection(ray2)
	assert.Nil(t, intersection)

	intersection = plane2.Intersection(ray1)
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 1, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 0, 1}, intersection.Normal)
	}

	intersection = plane2.Intersection(ray2)
	assert.Nil(t, intersection)
}

func TestPlane_IntersectionParallel(t *testing.T) {
	plane1, _ := NewPlane(geometry.Point{-50, -50, 0}, geometry.Vector{100, 0, 0}, geometry.Vector{0, 100, 0},
		shading.ShadingProperties{Opacity: 1})
	plane2, _ := NewPlane(geometry.Point{50, -50, -50}, geometry.Vector{-100, 100, 0}, geometry.Vector{-100, -100, 100},
		shading.ShadingProperties{Opacity: 1})
	ray1 := geometry.Ray{geometry.Point{0, 0, 0}, geometry.Vector{0, -3, 0}}
	ray2 := geometry.Ray{geometry.Point{0, 0, 0}, geometry.Vector{2, -1, -1}}

	intersection := plane1.Intersection(ray1)
	assert.Nil(t, intersection)

	intersection = plane1.Intersection(ray2)
	assert.NotNil(t, intersection)

	intersection = plane2.Intersection(ray1)
	assert.NotNil(t, intersection)

	intersection = plane2.Intersection(ray2)
	assert.Nil(t, intersection)
}

func TestPlane_ToTextureCoordinates(t *testing.T) {
	plane, _ := NewPlane(geometry.Point{1, -2, 3}, geometry.Vector{5, 0, 0}, geometry.Vector{0, 0, -4},
		shading.ShadingProperties{Opacity: 1})

	u, v := plane.ToTextureCoordinates(geometry.Point{1, -2, 3})
	assert.Equal(t, 0.0, u)
	assert.Equal(t, 0.0, v)

	u, v = plane.ToTextureCoordinates(geometry.Point{4.5, -2, -0.1})
	assert.Equal(t, 3.5, u)
	assert.Equal(t, 3.1, v)
}

func BenchmarkPlane_IntersectionHit(b *testing.B) {
	plane, _ := NewPlane(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shading.ShadingProperties{Opacity: 1})
	ray := geometry.Ray{geometry.Point{0, 1, 1.5}, geometry.Vector{0, 0, -1}}

	for n := 0; n < b.N; n++ {
		plane.Intersection(ray)
	}
}

func BenchmarkPlane_IntersectionMiss(b *testing.B) {
	plane, _ := NewPlane(geometry.Point{0, 0, 0}, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shading.ShadingProperties{Opacity: 1})
	ray := geometry.Ray{geometry.Point{0, -1, 1.5}, geometry.Vector{0, 0, -1}}

	for n := 0; n < b.N; n++ {
		plane.Intersection(ray)
	}
}
