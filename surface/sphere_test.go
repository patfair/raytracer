// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	shadingProperties := shading.ShadingProperties{
		DiffuseTexture:  shading.SolidTexture{shading.Color{0, 0.1, 0.2}},
		Reflectivity:    0.5,
		Opacity:         0.9,
		RefractiveIndex: 1.1,
	}
	sphere, err := NewSphere(geometry.Point{0, 0, 0}, 1, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shadingProperties)
	assert.Nil(t, err)

	assert.Equal(t, shadingProperties, sphere.ShadingProperties())
}

func TestNewSphereInvalid(t *testing.T) {
	_, err := NewSphere(geometry.Point{0, 0, 0}, -1, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shading.ShadingProperties{Opacity: 1})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "must be positive")

	_, err = NewSphere(geometry.Point{0, 0, 0}, 0, geometry.Vector{1, 0, 0}, geometry.Vector{0, 1, 0},
		shading.ShadingProperties{Opacity: 1})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "must be positive")

	_, err = NewSphere(geometry.Point{0, 0, 0}, 1, geometry.Vector{1, 0, 0}, geometry.Vector{1, 1, 0},
		shading.ShadingProperties{Opacity: 1})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "must be perpendicular")

	_, err = NewSphere(geometry.Point{0, 0, 0}, 1, geometry.Vector{1, 0, 0}, geometry.Vector{1, 1, 0},
		shading.ShadingProperties{Opacity: 1, Reflectivity: 2})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "reflectivity must be in [0, 1]")
}

func TestSphere_Intersection(t *testing.T) {
	// Intersecting from -X
	intersection := newTestSphere(geometry.Point{2, 0, 0}, 3).Intersection(geometry.Ray{geometry.Point{-4.5, 0, 0},
		geometry.Vector{1, 0, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 3.5, intersection.Distance)
		assert.Equal(t, geometry.Point{-1, 0, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{-1, 0, 0}, intersection.Normal)
	}

	// Intersecting from +Y
	intersection = newTestSphere(geometry.Point{0, 2, 0}, 3).Intersection(geometry.Ray{geometry.Point{0, 7.5, 0},
		geometry.Vector{0, -1, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 2.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 5, 0}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 1, 0}, intersection.Normal)
	}

	// Tangent
	intersection = newTestSphere(geometry.Point{0, 0, 0}, 5).Intersection(geometry.Ray{geometry.Point{-1.5, 0, 5},
		geometry.Vector{1, 0, 0}})
	if assert.NotNil(t, intersection) {
		assert.Equal(t, 1.5, intersection.Distance)
		assert.Equal(t, geometry.Point{0, 0, 5}, intersection.Point)
		assert.Equal(t, geometry.Vector{0, 0, 1}, intersection.Normal)
	}

	// Intersecting behind ray
	intersection = newTestSphere(geometry.Point{2, 0, 0}, 3).Intersection(geometry.Ray{geometry.Point{6, 0, 0},
		geometry.Vector{1, 0, 0}})
	assert.Nil(t, intersection)

	// Not intersecting
	intersection = newTestSphere(geometry.Point{0, 0, 0}, 1).Intersection(geometry.Ray{geometry.Point{0, 0, 2},
		geometry.Vector{1, 0, 0}})
	assert.Nil(t, intersection)
}

func TestSphere_ToTextureCoordinates(t *testing.T) {
	epsilon := 0.00001
	sphere := newTestSphere(geometry.Point{0, 0, 0}, 1)

	theta, phi := sphere.ToTextureCoordinates(geometry.Point{1, 0, 0})
	assert.Equal(t, 0.0, theta)
	assert.Equal(t, math.Pi/2, phi)

	theta, phi = sphere.ToTextureCoordinates(geometry.Point{0, 0, 1})
	assert.Equal(t, 0.0, theta)
	assert.Equal(t, 0.0, phi)

	theta, phi = sphere.ToTextureCoordinates(geometry.Point{0, 0, -1})
	assert.Equal(t, 0.0, theta)
	assert.Equal(t, math.Pi, phi)

	theta, phi = sphere.ToTextureCoordinates(geometry.Point{0, -1, 1})
	assert.Equal(t, -math.Pi/2, theta)
	assert.InEpsilon(t, math.Pi/4, phi, epsilon)

	theta, phi = sphere.ToTextureCoordinates(geometry.Point{-1, 0, -1})
	assert.Equal(t, math.Pi, theta)
	assert.Equal(t, 3*math.Pi/4, phi)
}

func BenchmarkSphere_IntersectionHit(b *testing.B) {
	sphere := newTestSphere(geometry.Point{2, 0, 0}, 3)
	ray := geometry.Ray{geometry.Point{-4.5, 0, 0}, geometry.Vector{1, 0, 0}}

	for n := 0; n < b.N; n++ {
		sphere.Intersection(ray)
	}
}

func BenchmarkSphere_IntersectionMiss(b *testing.B) {
	sphere := newTestSphere(geometry.Point{2, 0, 0}, 3)
	ray := geometry.Ray{geometry.Point{-4.5, 50, 100}, geometry.Vector{1, 0, 0}}

	for n := 0; n < b.N; n++ {
		sphere.Intersection(ray)
	}
}

func newTestSphere(point geometry.Point, radius float64) Sphere {
	sphere, _ := NewSphere(point, radius, geometry.Vector{0, 0, 1}, geometry.Vector{1, 0, 0},
		shading.ShadingProperties{Opacity: 1})
	return sphere
}
