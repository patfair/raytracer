package main

import (
	"github.com/cheggaaa/pb/v3"
	"math"
)

const (
	maxReflectionDepth = 20
	reflectionBias     = 0.001
	shadowBias         = 0.001
)

type RaytraceRowRequest struct {
	Scene       *Scene
	Row         [][]Ray
	RowIndex    int
	Pixels      [][]Color
	Progress    *pb.ProgressBar
	DoneChannel chan struct{}
}

func (request *RaytraceRowRequest) Run() {
	for x, pixelRays := range request.Row {
		var averagePixel Color
		for _, ray := range pixelRays {
			pixel := castRay(request.Scene, ray, 0, 1)
			averagePixel.R += pixel.R
			averagePixel.G += pixel.G
			averagePixel.B += pixel.B
		}
		averagePixel.R /= float64(len(pixelRays))
		averagePixel.G /= float64(len(pixelRays))
		averagePixel.B /= float64(len(pixelRays))
		request.Pixels[request.RowIndex][x] = averagePixel
		request.Progress.Increment()
	}

	request.DoneChannel <- struct{}{}
}

func castRay(scene *Scene, ray Ray, depth int, refractionIndex float64) Color {
	pixelColor := scene.BackgroundColor

	// Limit recursion caused by reflecting rays off multiple surfaces.
	if depth == maxReflectionDepth {
		return pixelColor
	}

	var closestIntersection *Intersection
	var closestSurface Surface
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
		var refractedColor, reflectedColor, diffuseColor, specularColor Color

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

			refractedRay := Ray{refractionPoint, refractionDirection.ToUnit()}
			refractedColor = castRay(scene, refractedRay, depth+1, shadingProperties.RefractiveIndex)
		}

		reflectedDirection := ray.Direction.Add(
			closestIntersection.Normal.Multiply(-2 * closestIntersection.Normal.Dot(ray.Direction))).ToUnit()
		if kReflection > 0 {
			// Bias the intersection point off the surface slightly to avoid immediate self-intersection.
			reflectedPoint := closestIntersection.Point.Translate(closestIntersection.Normal.Multiply(reflectionBias))

			reflectedRay := Ray{reflectedPoint, reflectedDirection.ToUnit()}
			reflectedColor = castRay(scene, reflectedRay, depth+1, refractionIndex)
		}

		if kDiffuse > 0 || kSpecular > 0 {
			for _, light := range scene.Lights {
				for i := 0; i < light.NumSamples(); i++ {
					// Check if there is an object between the intersection point and the light source, in which case
					// it should cast a shadow.
					lightRay := Ray{
						Point:     closestIntersection.Point,
						Direction: light.Direction(closestIntersection.Point, i).Multiply(-1).ToUnit(),
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

					incidentDotProduct :=
						light.Direction(closestIntersection.Point, i).Multiply(-1).Dot(closestIntersection.Normal)
					incidentLight := light.Intensity(closestIntersection.Point) * math.Max(incidentDotProduct, 0) *
						transparency / float64(light.NumSamples())
					diffuseColor.R += closestSurface.AlbedoAt(closestIntersection.Point).R / math.Pi * light.Color().R *
						incidentLight
					diffuseColor.G += closestSurface.AlbedoAt(closestIntersection.Point).G / math.Pi * light.Color().G *
						incidentLight
					diffuseColor.B += closestSurface.AlbedoAt(closestIntersection.Point).B / math.Pi * light.Color().B *
						incidentLight

					// Calculate specular reflection.
					specularIntensity := math.Pow(math.Max(reflectedDirection.Dot(lightRay.Direction), 0),
						shadingProperties.SpecularExponent) / float64(light.NumSamples())
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
