package main

func SpheresScene() (*Scene, *Camera, error) {
	tealSphere := newSphere(Point{0, 20, 1}, Color{0.1, 0.7, 1})
	greenSphere := newSphere(Point{-2, 15, 1}, Color{0, 0.4, 0})
	redSphere := newSphere(Point{2.5, 21, 1}, Color{0.8, 0, 0})
	blueSphere := newSphere(Point{1, 9, 1}, Color{0, 0.3, 0.8})
	yellowSphere := newSphere(Point{-3, 10, 1}, Color{0.9, 0.7, 0})
	purpleSphere := newSphere(Point{4, 10.5, 1}, Color{0.75, 0.2, 0.8})
	graySphere := newSphere(Point{3.5, 16, 1}, Color{0.8, 0.8, 0.8})

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
			radius:     0, //2,
			numSamples: 1, //20,
		},
	}

	cameraOrigin := Point{0, 0, 3}
	focalDistance := cameraOrigin.VectorTo(blueSphere.Center).Norm()
	camera, err := NewCamera(Ray{cameraOrigin, Vector{0, 1, -0.2}}, Vector{0, 0.2, 1}, 400, 225, 40, 0.1,
		focalDistance, 20, 2)
	//camera, err := NewCamera(Ray{Point{0, 0, 3}, Vector{0, 1, -0.2}}, Vector{0, 0.2, 1}, 3840, 2160, 40, 2)

	return &Scene{Surfaces: surfaces, Lights: lights, BackgroundColor: Color{0, 0, 0}}, camera, err
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
