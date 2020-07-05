package main

import "math"

func AllElementsScene() (*Scene, *Camera, error) {
	surfaces := []Surface{
		// YZ plane
		Plane{
			Corner: Point{0, 0, 0},
			Width:  Vector{0, 4, 0},
			Height: Vector{0, 0, 2},
			shadingProperties: ShadingProperties{
				DiffuseTexture: CheckerboardTexture{Color{0.9, 0.1, 0.1}, Color{0.8, 0.8, 0.8}, 1, 0.5},
				Opacity:        1,
			},
		},

		// XZ plane
		Plane{
			Corner: Point{0, 0, 0},
			Width:  Vector{4, 0, 0},
			Height: Vector{0, 0, 2},
			shadingProperties: ShadingProperties{
				DiffuseTexture: CheckerboardTexture{Color{0.2, 0.5, 1}, Color{0, 0, 0}, 0.1, 0.1},
				Opacity:        1,
			},
		},

		// XY plane
		Plane{
			Corner: Point{0, 0, 0},
			Width:  Vector{4, 0, 0},
			Height: Vector{0, 10, 0},
			shadingProperties: ShadingProperties{
				DiffuseTexture: CheckerboardTexture{Color{0.9, 0.9, 0.9}, Color{0.2, 0.2, 0.2}, 0.3, 0.3},
				Opacity:        1,
			},
		},

		// Mirror sphere
		Sphere{
			Center:           Point{1.5, 1.5, 0.75},
			Radius:           0.5,
			ZenithReference:  Vector{0, 0, 1},
			AzimuthReference: Vector{1, 0, 0},
			shadingProperties: ShadingProperties{
				DiffuseTexture:    SolidTexture{Color{0.5, 0.5, 0.5}},
				SpecularExponent:  100,
				SpecularIntensity: 0.5,
				Opacity:           1,
				Reflectivity:      0.8,
			},
		},

		// Checkerboard sphere
		Sphere{
			Center:           Point{1, 4.4, 1},
			Radius:           0.3,
			ZenithReference:  Vector{0, 1, 0},
			AzimuthReference: Vector{1, 0, 0},
			shadingProperties: ShadingProperties{
				DiffuseTexture:    CheckerboardTexture{Color{1, 1, 1}, Color{0, 0, 1}, math.Pi / 2, math.Pi / 4},
				SpecularExponent:  100,
				SpecularIntensity: 0.5,
				Opacity:           1,
				Reflectivity:      0.3,
			},
		},

		// Checkerboard disc
		Disc{
			plane: Plane{
				Corner: Point{3, 1, 0.5},
				Width:  Vector{0.5, 0, 0},
				Height: Vector{0, 1, 0},
				shadingProperties: ShadingProperties{
					DiffuseTexture: CheckerboardTexture{Color{0.9, 0.8, 0.4}, Color{0.3, 0.3, 0}, 0.125, math.Pi / 2},
					Opacity:        1,
				},
			},
		},

		// Mirror disc
		Disc{
			plane: Plane{
				Corner: Point{2, 2, 0.1},
				Width:  Vector{1.5, 0, 0},
				Height: Vector{0, 1, 0},
				shadingProperties: ShadingProperties{
					DiffuseTexture:    SolidTexture{Color{0, 0, 0}},
					SpecularExponent:  100,
					SpecularIntensity: 0.5,
					Opacity:           1,
					Reflectivity:      0.7,
				},
			},
		},
	}

	// Gold cube
	cubeFront := Plane{
		Corner: Point{1, 3, 0.75},
		Width:  Vector{0, 0.5, 0.5},
		Height: Vector{0, -0.5, 0.5},
		shadingProperties: ShadingProperties{
			DiffuseTexture:    SolidTexture{Color{0.9, 0.6, 0.2}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           1,
			Reflectivity:      0.1,
		},
	}
	for _, plane := range NewBox(cubeFront, 0.5) {
		surfaces = append(surfaces, plane)
	}

	// Glass panel
	boxFront := Plane{
		Corner: Point{2.5, 4.3, 0.1},
		Width:  Vector{-0.8, 0.6, 0},
		Height: Vector{0, 0, 2},
		shadingProperties: ShadingProperties{
			DiffuseTexture:    SolidTexture{Color{0, 1, 0}},
			SpecularExponent:  100,
			SpecularIntensity: 0.5,
			Opacity:           0.1,
			Reflectivity:      0.5,
			RefractiveIndex:   1.1,
		},
	}
	for _, plane := range NewBox(boxFront, 0.05) {
		surfaces = append(surfaces, plane)
	}

	lights := []Light{
		DistantLight{
			direction:  Vector{-10, -10, -20},
			color:      Color{1, 1, 1},
			intensity:  0.75,
			numSamples: 1,
		},
		DistantLight{
			direction:  Vector{-10, -10, -25},
			color:      Color{1, 1, 1},
			intensity:  0.75,
			numSamples: 1,
		},
		DistantLight{
			direction:  Vector{-11, -9, -20},
			color:      Color{1, 1, 1},
			intensity:  0.75,
			numSamples: 1,
		},
		PointLight{
			point:      Point{5, 1, 10},
			color:      Color{1, 1, 1},
			intensity:  1000,
			radius:     0,
			numSamples: 1,
		},
	}

	camera, err := NewCamera(Ray{Point{10, 10, 5}, Vector{-10, -10, -5}}, Vector{-10, -10, 40}, 3840, 2160, 30, 0, 0, 1,
		2)

	return &Scene{Surfaces: surfaces, Lights: lights, BackgroundColor: Color{0.1, 0.8, 1}}, camera, err
}
