package main

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
	"math"
)

const (
	maxReflectionDepth = 20
	reflectionBias     = 0.001
	shadowBias         = 0.001
	colorThreshold     = 0.01
	adjacentPixels     = 4
)

type RaytraceRowRequest struct {
	Scene       *Scene
	Camera      *Camera
	RowIndex    int
	IsDraft     bool
	DraftPixels [][]shading.Color
	Pixels      [][]shading.Color
	Progress    *pb.ProgressBar
	DoneChannel chan struct{}
}

func (request *RaytraceRowRequest) Run() {
	camera := request.Camera
	for j := 0; j < camera.Width; j++ {
		supersampleFactor := camera.SupersampleFactor
		supersamplingRequired := !request.IsDraft && isSupersamplingRequired(request.DraftPixels, j, request.RowIndex,
			adjacentPixels)
		if !supersamplingRequired {
			supersampleFactor = 1
		}

		var averagePixel shading.Color
		for n := 0; n < camera.DepthOfFieldSamples; n++ {
			for a := 0; a < supersampleFactor; a++ {
				for b := 0; b < supersampleFactor; b++ {
					ray := camera.GetRay(j, request.RowIndex, n, supersampleFactor, a, b)
					pixel := request.castRay(request.Scene, ray, 0, 1, supersamplingRequired)
					averagePixel.R += pixel.R
					averagePixel.G += pixel.G
					averagePixel.B += pixel.B
				}
			}
		}

		averagePixel.R /= float64(camera.DepthOfFieldSamples * supersampleFactor * supersampleFactor)
		averagePixel.G /= float64(camera.DepthOfFieldSamples * supersampleFactor * supersampleFactor)
		averagePixel.B /= float64(camera.DepthOfFieldSamples * supersampleFactor * supersampleFactor)
		request.Pixels[request.RowIndex][j] = averagePixel
		request.Progress.Increment()
	}

	request.DoneChannel <- struct{}{}
}

func (request *RaytraceRowRequest) castRay(scene *Scene, ray geometry.Ray, depth int, refractionIndex float64,
	supersamplingRequired bool) shading.Color {
	pixelColor := scene.BackgroundColor

	// Limit recursion caused by reflecting rays off multiple surfaces.
	if depth == maxReflectionDepth {
		return pixelColor
	}

	var closestIntersection *geometry.Intersection
	var closestSurface surface.Surface
	for _, surface := range scene.Surfaces {
		if intersection := surface.Intersection(ray); intersection != nil {
			if closestIntersection == nil || intersection.Distance < closestIntersection.Distance {
				closestIntersection = intersection
				closestSurface = surface
			}
		}
	}

	if closestIntersection != nil {
		shadingProperties := closestSurface.ShadingProperties()
		kRefraction := 1 - shadingProperties.Opacity
		kReflection := shadingProperties.Reflectivity * shadingProperties.Opacity
		kDiffuse := 1 - kRefraction - kReflection
		kSpecular := shadingProperties.SpecularIntensity
		var refractedColor, reflectedColor, diffuseColor, specularColor shading.Color

		if kRefraction > 0 {
			cosIn := -closestIntersection.Normal.Dot(ray.Direction)
			etaIn := refractionIndex
			etaOut := shadingProperties.RefractiveIndex
			if refractionIndex > 1 {
				// If the previous refraction index isn't 1, the ray is exiting the material instead of entering.
				etaIn, etaOut = etaOut, etaIn
			}

			// Determine reflection and refraction components of the refracted light.
			sinOut := etaIn / etaOut * math.Sqrt(math.Max(1-cosIn*cosIn, 0))
			if sinOut < 1 {
				cosOut := math.Abs(math.Sqrt(math.Max(1-sinOut*sinOut, 0)))
				rParallel := ((etaOut * cosIn) - (etaIn * cosOut)) / ((etaOut * cosIn) + (etaIn * cosOut))
				rPerpendicular := ((etaIn * cosIn) - (etaOut * cosOut)) / ((etaIn * cosIn) + (etaOut * cosOut))
				reflectedComponent := (rParallel*rParallel + rPerpendicular*rPerpendicular) / 2

				// Adjust the coefficients for the reflected light coming from refraction.
				delta := reflectedComponent * kRefraction
				kRefraction -= delta
				kReflection += delta
			}

			eta := etaIn / etaOut
			k := 1 - eta*eta*(1-cosIn*cosIn)
			refractionDirection :=
				ray.Direction.Multiply(eta).Add(closestIntersection.Normal.Multiply(eta*cosIn - math.Sqrt(k)))

			// Bias the intersection point off the surface slightly to avoid immediate self-intersection.
			refractionPoint :=
				closestIntersection.Point.Translate(closestIntersection.Normal.Multiply(-reflectionBias))

			refractedRay := geometry.Ray{refractionPoint, refractionDirection.ToUnit()}
			refractedColor = request.castRay(scene, refractedRay, depth+1, shadingProperties.RefractiveIndex,
				supersamplingRequired)
		}

		reflectedDirection := ray.Direction.Add(
			closestIntersection.Normal.Multiply(-2 * closestIntersection.Normal.Dot(ray.Direction))).ToUnit()
		if kReflection > 0 {
			// Bias the intersection point off the surface slightly to avoid immediate self-intersection.
			reflectedPoint := closestIntersection.Point.Translate(closestIntersection.Normal.Multiply(reflectionBias))

			reflectedRay := geometry.Ray{reflectedPoint, reflectedDirection.ToUnit()}
			reflectedColor = request.castRay(scene, reflectedRay, depth+1, refractionIndex, supersamplingRequired)
		}

		if kDiffuse > 0 || kSpecular > 0 {
			for _, light := range scene.Lights {
				numSamples := light.NumSamples()
				if !supersamplingRequired && !request.IsDraft {
					numSamples = 1
				}
				for i := 0; i < numSamples; i++ {
					// Check if there is an object between the intersection point and the light source, in which case
					// it should cast a shadow.
					lightRay := geometry.Ray{
						Origin:    closestIntersection.Point,
						Direction: light.Direction(closestIntersection.Point, i, numSamples).Multiply(-1).ToUnit(),
					}
					transparency := 1.0
					for _, surface := range scene.Surfaces {
						if intersection := surface.Intersection(lightRay); intersection != nil {
							// Require a minimum distance to avoid a surface from shadowing itself.
							if intersection.Distance > shadowBias {
								if light.IsBlockedByIntersection(closestIntersection.Point, intersection) {
									transparency *= 1 - surface.ShadingProperties().Opacity
								}
							}
						}
					}
					if transparency == 0 {
						continue
					}

					incidentDotProduct := light.Direction(closestIntersection.Point, i,
						numSamples).Multiply(-1).Dot(closestIntersection.Normal)
					incidentLight := light.Intensity(closestIntersection.Point) * math.Max(incidentDotProduct, 0) *
						transparency / float64(numSamples)
					u, v := closestSurface.ToTextureCoordinates(closestIntersection.Point)
					diffuseColor.R += closestSurface.ShadingProperties().DiffuseTexture.AlbedoAt(u, v).R / math.Pi *
						light.Color().R * incidentLight
					diffuseColor.G += closestSurface.ShadingProperties().DiffuseTexture.AlbedoAt(u, v).G / math.Pi *
						light.Color().G * incidentLight
					diffuseColor.B += closestSurface.ShadingProperties().DiffuseTexture.AlbedoAt(u, v).B / math.Pi *
						light.Color().B * incidentLight

					// Calculate specular reflection.
					specularIntensity := math.Pow(math.Max(reflectedDirection.Dot(lightRay.Direction), 0),
						shadingProperties.SpecularExponent) / float64(numSamples)
					specularColor.R += light.Color().R * specularIntensity
					specularColor.G += light.Color().G * specularIntensity
					specularColor.B += light.Color().B * specularIntensity
				}
			}
		}

		pixelColor.R = kRefraction*refractedColor.R + kReflection*reflectedColor.R + kDiffuse*diffuseColor.R +
			kSpecular*specularColor.R
		pixelColor.G = kRefraction*refractedColor.G + kReflection*reflectedColor.G + kDiffuse*diffuseColor.G +
			kSpecular*specularColor.G
		pixelColor.B = kRefraction*refractedColor.B + kReflection*reflectedColor.B + kDiffuse*diffuseColor.B +
			kSpecular*specularColor.B
	}

	return pixelColor
}

func isSupersamplingRequired(draftPixels [][]shading.Color, x, y, numAdjacent int) bool {
	for i := y - numAdjacent; i <= y+numAdjacent; i++ {
		for j := x - numAdjacent; j <= x+numAdjacent; j++ {
			if !arePixelsSimilar(draftPixels, x, y, j, i) {
				return true
			}
		}
	}
	return false
}

func arePixelsSimilar(draftPixels [][]shading.Color, xA, yA, xB, yB int) bool {
	height := len(draftPixels)
	if height == 0 {
		return true
	}
	width := len(draftPixels[0])
	if xA < 0 || xA >= width || xB < 0 || xB >= width || yA < 0 || yA >= height || yB < 0 || yB >= height {
		return true
	}

	pixelA := draftPixels[yA][xA]
	pixelB := draftPixels[yB][xB]
	return math.Abs(pixelA.R-pixelB.R) <= colorThreshold && math.Abs(pixelA.G-pixelB.G) <= colorThreshold &&
		math.Abs(pixelA.B-pixelB.B) <= colorThreshold
}
