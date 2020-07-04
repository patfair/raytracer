package main

func SpheresScene() Scene {
	surfaces := []Surface{
		// XY plane
		Plane{
			Corner: Point{-50, -50, 0},
			Width:  Vector{100, 0, 0},
			Height: Vector{0, 100, 0},
			shadingProperties: ShadingProperties{
				DiffuseTexture: CheckerboardTexture{Color{0.9, 0.75, 0.55}, Color{0.2, 0.1, .05}, 1.5, 1.5},
				Opacity:        1,
				Reflectivity:   0.2,
			},
		},

		newSphere(Point{0, 20, 1}, Color{0.1, 0.7, 1}),      // Teal sphere
		newSphere(Point{-2, 15, 1}, Color{0, 0.4, 0}),       // Dark green sphere
		newSphere(Point{2.5, 21, 1}, Color{0.8, 0, 0}),      // Red sphere
		newSphere(Point{1, 9, 1}, Color{0, 0.3, 0.8}),       // Blue sphere
		newSphere(Point{-3, 10, 1}, Color{0.9, 0.7, 0}),     // Yellow sphere
		newSphere(Point{4, 10.5, 1}, Color{0.75, 0.2, 0.8}), // Purple sphere
		newSphere(Point{3.5, 16, 1}, Color{0.8, 0.8, 0.8}),  // Gray sphere
	}

	// Glass panel
	glassBaseHeight := 0.05
	glassPaneWidth := 0.1
	glassPane := Plane{
		Corner: Point{-2.5, 8, glassBaseHeight},
		Width:  Vector{1.5, 0, 0},
		Height: Vector{0, 0, 2},
		shadingProperties: ShadingProperties{
			DiffuseTexture:    SolidTexture{Color{1, 1, 1}},
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
		Corner: Point{glassPane.Corner.X - margin, glassPane.Corner.Y + margin, 0},
		Width:  Vector{glassPane.Width.X + 2*margin, 0, 0},
		Height: Vector{0, 0, glassBaseHeight},
		shadingProperties: ShadingProperties{
			DiffuseTexture: SolidTexture{Color{1, 1, 1}},
			Opacity:        1,
			Reflectivity:   0.05,
		},
	}
	for _, plane := range NewBox(glassBase, glassPaneWidth+2*margin) {
		surfaces = append(surfaces, plane)
	}

	lights := []Light{
		PointLight{
			point:      Point{10, 0, 30},
			color:      Color{1, 1, 0.8},
			intensity:  30000,
			radius:     2,
			numSamples: 20,
		},
	}

	return Scene{Surfaces: surfaces, Lights: lights, BackgroundColor: Color{0, 0, 0}}
}

func newSphere(point Point, color Color) Sphere {
	return Sphere{
		Center:           point,
		Radius:           1,
		ZenithReference:  Vector{1, 0, 0},
		AzimuthReference: Vector{0, 1, 0},
		shadingProperties: ShadingProperties{
			DiffuseTexture:    SolidTexture{color},
			SpecularExponent:  200,
			SpecularIntensity: 2,
			Opacity:           1,
			Reflectivity:      0.2,
		},
	}
}
