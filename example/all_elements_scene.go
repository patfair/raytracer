// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package example

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/light"
	"github.com/patfair/raytracer/render"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
	"math"
)

// Creates a scene with pretty much every possible element type in a corner defined by three perpendicular planes.
func AllElementsScene(frame int) (*render.Scene, error) {
	camera, err := render.NewCamera(geometry.Ray{geometry.Point{10, 10, 5}, geometry.Vector{-10, -10, -5}},
		geometry.Vector{-10, -10, 40}, 30, 0, 1, 1, 2)
	if err != nil {
		return nil, err
	}

	scene := render.Scene{Camera: camera, BackgroundColor: shading.Color{0.1, 0.8, 1}}

	yzPlane, err := surface.NewPlane(
		geometry.Point{0, 0, 0},
		geometry.Vector{0, 4, 0},
		geometry.Vector{0, 0, 2},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{shading.Color{0.9, 0.1, 0.1}, shading.Color{0.8, 0.8, 0.8}, 1,
				0.5},
			Opacity: 1,
		},
	)
	if err != nil {
		return nil, err
	}
	scene.AddSurface(yzPlane)

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
		return nil, err
	}
	scene.AddSurface(xzPlane)

	xyPlane, err := surface.NewPlane(
		geometry.Point{0, 0, 0},
		geometry.Vector{4, 0, 0},
		geometry.Vector{0, 10, 0},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{shading.Color{0.9, 0.9, 0.9}, shading.Color{0.2, 0.2, 0.2}, 0.3,
				0.3},
			Opacity: 1,
		},
	)
	if err != nil {
		return nil, err
	}
	scene.AddSurface(xyPlane)

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
		return nil, err
	}
	scene.AddSurface(mirrorSphere)

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
		return nil, err
	}
	scene.AddSurface(checkerboardSphere)

	checkerboardDisc, err := surface.NewDisc(
		geometry.Point{3, 1, 0.5},
		geometry.Vector{0.5, 0, 0},
		geometry.Vector{0, 0.5, 0},
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
		return nil, err
	}
	scene.AddSurface(checkerboardDisc)

	mirrorDisc, err := surface.NewDisc(
		geometry.Point{2, 2, 0.1},
		geometry.Vector{1.5, 0, 0},
		geometry.Vector{0, 1.5, 0},
		shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{shading.Color{0, 0, 0}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           1,
			Reflectivity:      0.7,
		},
	)
	if err != nil {
		return nil, err
	}
	scene.AddSurface(mirrorDisc)

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
		return nil, err
	}
	for _, plane := range goldCubePlanes {
		scene.AddSurface(plane)
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
		return nil, err
	}
	for _, plane := range boxPlanes {
		scene.AddSurface(plane)
	}

	light1, err := light.NewDistantLight(
		geometry.Vector{-10, -10, -20},
		shading.Color{1, 1, 1},
		0.75,
		0,
	)
	if err != nil {
		return nil, err
	}
	scene.AddLight(light1)

	light2, err := light.NewDistantLight(
		geometry.Vector{-10, -10, -25},
		shading.Color{1, 1, 1},
		0.75,
		0,
	)
	if err != nil {
		return nil, err
	}
	scene.AddLight(light2)

	light3, err := light.NewDistantLight(
		geometry.Vector{-11, -9, -20},
		shading.Color{1, 1, 1},
		0.75,
		0,
	)
	if err != nil {
		return nil, err
	}
	scene.AddLight(light3)

	light4, err := light.NewPointLight(
		geometry.Point{5, 1, 10},
		shading.Color{1, 1, 1},
		1000,
		0,
	)
	if err != nil {
		return nil, err
	}
	scene.AddLight(light4)

	return &scene, nil
}
