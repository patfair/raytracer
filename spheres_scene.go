package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/light"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
)

func SpheresScene() (*Scene, *Camera, error) {
	// Floor plane
	xyPlane, err := surface.NewPlane(
		geometry.Point{-50, -50, 0},
		geometry.Vector{100, 0, 0},
		geometry.Vector{0, 100, 0},
		shading.ShadingProperties{
			DiffuseTexture: shading.CheckerboardTexture{shading.Color{0.9, 0.75, 0.55}, shading.Color{0.2, 0.1, .05}, 1.5, 1.5},
			Opacity:        1,
			Reflectivity:   0.2,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	// Colored spheres
	tealSphere, err := newSphere(geometry.Point{0, 20, 1}, shading.Color{0.1, 0.7, 1})
	if err != nil {
		return nil, nil, err
	}
	greenSphere, err := newSphere(geometry.Point{-2, 15, 1}, shading.Color{0, 0.4, 0})
	if err != nil {
		return nil, nil, err
	}
	redSphere, err := newSphere(geometry.Point{2.5, 21, 1}, shading.Color{0.8, 0, 0})
	if err != nil {
		return nil, nil, err
	}
	blueSphere, err := newSphere(geometry.Point{1, 9, 1}, shading.Color{0, 0.3, 0.8})
	if err != nil {
		return nil, nil, err
	}
	yellowSphere, err := newSphere(geometry.Point{-3, 10, 1}, shading.Color{0.9, 0.7, 0})
	if err != nil {
		return nil, nil, err
	}
	purpleSphere, err := newSphere(geometry.Point{4, 10.5, 1}, shading.Color{0.75, 0.2, 0.8})
	if err != nil {
		return nil, nil, err
	}
	graySphere, err := newSphere(geometry.Point{3.5, 16, 1}, shading.Color{0.8, 0.8, 0.8})
	if err != nil {
		return nil, nil, err
	}

	// Glass panel
	glassBaseHeight := 0.05
	glassPaneWidth := 1.5
	glassPaneDepth := 0.1
	glassPaneCorner := geometry.Point{-2.5, 8, glassBaseHeight}
	glassPanePlanes, err := surface.NewBox(
		glassPaneCorner,
		geometry.Vector{glassPaneWidth, 0, 0},
		geometry.Vector{0, 0, 2},
		glassPaneDepth,
		shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{shading.Color{1, 1, 1}},
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

	// Base for glass panel
	margin := 0.03
	glassBasePlanes, err := surface.NewBox(
		geometry.Point{glassPaneCorner.X - margin, glassPaneCorner.Y + margin, 0},
		geometry.Vector{glassPaneWidth + 2*margin, 0, 0},
		geometry.Vector{0, 0, glassBaseHeight},
		glassPaneDepth+2*margin,
		shading.ShadingProperties{
			DiffuseTexture: shading.SolidTexture{shading.Color{1, 1, 1}},
			Opacity:        1,
			Reflectivity:   0.05,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	surfaces := []surface.Surface{
		xyPlane,
		tealSphere,
		greenSphere,
		redSphere,
		blueSphere,
		yellowSphere,
		purpleSphere,
		graySphere,
	}
	for _, plane := range glassPanePlanes {
		surfaces = append(surfaces, plane)
	}
	for _, plane := range glassBasePlanes {
		surfaces = append(surfaces, plane)
	}

	pointLight, err := light.NewPointLight(
		geometry.Point{10, 0, 30},
		shading.Color{1, 1, 0.8},
		30000,
		2,
		20,
	)
	if err != nil {
		return nil, nil, err
	}
	lights := []light.Light{pointLight}

	cameraOrigin := geometry.Point{0, 0, 3}
	focalDistance := cameraOrigin.DistanceTo(blueSphere.Center())
	camera, err := NewCamera(geometry.Ray{cameraOrigin, geometry.Vector{0, 1, -0.2}}, geometry.Vector{0, 0.2, 1}, 3840,
		2160, 40, 0.06, focalDistance, 25, 2)

	return &Scene{Surfaces: surfaces, Lights: lights, BackgroundColor: shading.Color{0, 0, 0}}, camera, err
}

func newSphere(point geometry.Point, color shading.Color) (surface.Sphere, error) {
	return surface.NewSphere(
		point,
		1,
		geometry.Vector{1, 0, 0},
		geometry.Vector{0, 1, 0},
		shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{color},
			SpecularExponent:  200,
			SpecularIntensity: 2,
			Opacity:           1,
			Reflectivity:      0.2,
		},
	)
}
