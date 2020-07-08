package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
)

func SpheresScene() (*Scene, *Camera, error) {
	tealSphere := newSphere(geometry.Point{0, 20, 1}, shading.Color{0.1, 0.7, 1})
	greenSphere := newSphere(geometry.Point{-2, 15, 1}, shading.Color{0, 0.4, 0})
	redSphere := newSphere(geometry.Point{2.5, 21, 1}, shading.Color{0.8, 0, 0})
	blueSphere := newSphere(geometry.Point{1, 9, 1}, shading.Color{0, 0.3, 0.8})
	yellowSphere := newSphere(geometry.Point{-3, 10, 1}, shading.Color{0.9, 0.7, 0})
	purpleSphere := newSphere(geometry.Point{4, 10.5, 1}, shading.Color{0.75, 0.2, 0.8})
	graySphere := newSphere(geometry.Point{3.5, 16, 1}, shading.Color{0.8, 0.8, 0.8})

	surfaces := []Surface{
		// XY plane
		Plane{
			Corner: geometry.Point{-50, -50, 0},
			Width:  geometry.Vector{100, 0, 0},
			Height: geometry.Vector{0, 100, 0},
			shadingProperties: shading.ShadingProperties{
				DiffuseTexture: shading.CheckerboardTexture{shading.Color{0.9, 0.75, 0.55}, shading.Color{0.2, 0.1, .05}, 1.5, 1.5},
				Opacity:        1,
				Reflectivity:   0.2,
			},
		},

		tealSphere,
		greenSphere,
		redSphere,
		blueSphere,
		yellowSphere,
		purpleSphere,
		graySphere,
	}

	// Glass panel
	glassBaseHeight := 0.05
	glassPaneWidth := 0.1
	glassPane := Plane{
		Corner: geometry.Point{-2.5, 8, glassBaseHeight},
		Width:  geometry.Vector{1.5, 0, 0},
		Height: geometry.Vector{0, 0, 2},
		shadingProperties: shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{shading.Color{1, 1, 1}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           0.1,
			Reflectivity:      0.5,
			RefractiveIndex:   1.1,
		},
	}
	for _, plane := range NewBox(glassPane, glassPaneWidth) {
		surfaces = append(surfaces, plane)
	}

	// Base for glass panel
	margin := 0.03
	glassBase := Plane{
		Corner: geometry.Point{glassPane.Corner.X - margin, glassPane.Corner.Y + margin, 0},
		Width:  geometry.Vector{glassPane.Width.X + 2*margin, 0, 0},
		Height: geometry.Vector{0, 0, glassBaseHeight},
		shadingProperties: shading.ShadingProperties{
			DiffuseTexture: shading.SolidTexture{shading.Color{1, 1, 1}},
			Opacity:        1,
			Reflectivity:   0.05,
		},
	}
	for _, plane := range NewBox(glassBase, glassPaneWidth+2*margin) {
		surfaces = append(surfaces, plane)
	}

	lights := []Light{
		PointLight{
			point:      geometry.Point{10, 0, 30},
			color:      shading.Color{1, 1, 0.8},
			intensity:  30000,
			radius:     2,
			numSamples: 20,
		},
	}

	cameraOrigin := geometry.Point{0, 0, 3}
	focalDistance := cameraOrigin.DistanceTo(blueSphere.Center)
	camera, err := NewCamera(geometry.Ray{cameraOrigin, geometry.Vector{0, 1, -0.2}}, geometry.Vector{0, 0.2, 1}, 3840, 2160, 40, 0.06,
		focalDistance, 25, 2)

	return &Scene{Surfaces: surfaces, Lights: lights, BackgroundColor: shading.Color{0, 0, 0}}, camera, err
}

func newSphere(point geometry.Point, color shading.Color) Sphere {
	return Sphere{
		Center:           point,
		Radius:           1,
		ZenithReference:  geometry.Vector{1, 0, 0},
		AzimuthReference: geometry.Vector{0, 1, 0},
		shadingProperties: shading.ShadingProperties{
			DiffuseTexture:    shading.SolidTexture{color},
			SpecularExponent:  200,
			SpecularIntensity: 2,
			Opacity:           1,
			Reflectivity:      0.2,
		},
	}
}
