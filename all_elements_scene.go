package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
	"math"
)

func AllElementsScene() (*Scene, *Camera, error) {
	yzPlane, err := surface.NewPlane(
		geometry.Point{0, 0, 0},
		geometry.Vector{0, 4, 0},
		geometry.Vector{0, 0, 2},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{shading.Color{0.9, 0.1, 0.1}, shading.Color{0.8, 0.8, 0.8}, 1, 0.5},
			Opacity:        1,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	xzPlane, err := surface.NewPlane(
		geometry.Point{0, 0, 0},
		geometry.Vector{4, 0, 0},
		geometry.Vector{0, 0, 2},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{shading.Color{0.2, 0.5, 1}, shading.Color{0, 0, 0}, 0.1, 0.1},
			Opacity:        1,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	xyPlane, err := surface.NewPlane(
		geometry.Point{0, 0, 0},
		geometry.Vector{4, 0, 0},
		geometry.Vector{0, 10, 0},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{shading.Color{0.9, 0.9, 0.9}, shading.Color{0.2, 0.2, 0.2}, 0.3, 0.3},
			Opacity:        1,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	mirrorSphere, err := surface.NewSphere(
		geometry.Point{1.5, 1.5, 0.75},
		0.5,
		geometry.Vector{0, 0, 1},
		geometry.Vector{1, 0, 0},
		shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{shading.Color{0.5, 0.5, 0.5}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           1,
			Reflectivity:      0.8,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	checkerboardSphere, err := surface.NewSphere(
		geometry.Point{1, 4.4, 1},
		0.3,
		geometry.Vector{0, 1, 0},
		geometry.Vector{1, 0, 0},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{
				Color1: shading.Color{1, 1, 1},
				Color2: shading.Color{0, 0, 1},
				UPitch: math.Pi / 2,
				VPitch: math.Pi / 4,
			},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           1,
			Reflectivity:      0.3,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	checkerboardDisc, err := surface.NewDisc(
		geometry.Point{3, 1, 0.5},
		geometry.Vector{0.5, 0, 0},
		geometry.Vector{0, 1, 0},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{
				Color1: shading.Color{0.9, 0.8, 0.4},
				Color2: shading.Color{0.3, 0.3, 0},
				UPitch: 0.125,
				VPitch: math.Pi / 2,
			},
			Opacity: 1,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	mirrorDisc, err := surface.NewDisc(
		geometry.Point{2, 2, 0.1},
		geometry.Vector{1.5, 0, 0},
		geometry.Vector{0, 1, 0},
		shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{shading.Color{0, 0, 0}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           1,
			Reflectivity:      0.7,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	goldCubePlanes, err := surface.NewBox(
		geometry.Point{1, 3, 0.75},
		geometry.Vector{0, 0.5, 0.5},
		geometry.Vector{0, -0.5, 0.5},
		0.5,
		shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{shading.Color{0.9, 0.6, 0.2}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           1,
			Reflectivity:      0.1,
		},
	)
	if err != nil {
		return nil, nil, err
	}
	boxPlanes, err := surface.NewBox(
		geometry.Point{2.5, 4.3, 0.1},
		geometry.Vector{-0.8, 0.6, 0},
		geometry.Vector{0, 0, 2},
		0.05,
		shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{shading.Color{0, 1, 0}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           0.1,
			Reflectivity:      0.5,
			RefractiveIndex:   1.1,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	surfaces := []surface.Surface{
		yzPlane,
		xzPlane,
		xyPlane,
		mirrorSphere,
		checkerboardSphere,
		checkerboardDisc,
		mirrorDisc,
	}
	for _, plane := range goldCubePlanes {
		surfaces = append(surfaces, plane)
	}
	for _, plane := range boxPlanes {
		surfaces = append(surfaces, plane)
	}

	lights := []Light{
		DistantLight{
			direction:  geometry.Vector{-10, -10, -20},
			color:      shading.Color{1, 1, 1},
			intensity:  0.75,
			numSamples: 1,
		},
		DistantLight{
			direction:  geometry.Vector{-10, -10, -25},
			color:      shading.Color{1, 1, 1},
			intensity:  0.75,
			numSamples: 1,
		},
		DistantLight{
			direction:  geometry.Vector{-11, -9, -20},
			color:      shading.Color{1, 1, 1},
			intensity:  0.75,
			numSamples: 1,
		},
		PointLight{
			point:      geometry.Point{5, 1, 10},
			color:      shading.Color{1, 1, 1},
			intensity:  1000,
			radius:     0,
			numSamples: 1,
		},
	}

	camera, err := NewCamera(geometry.Ray{geometry.Point{10, 10, 5}, geometry.Vector{-10, -10, -5}},
		geometry.Vector{-10, -10, 40}, 3840, 2160, 30, 0, 0, 1, 2)

	return &Scene{Surfaces: surfaces, Lights: lights, BackgroundColor: shading.Color{0.1, 0.8, 1}}, camera, err
}
