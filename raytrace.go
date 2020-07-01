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

type RaytraceRowsRequest struct {
	Scene       *Scene
	Rays        [][]Ray
	Start       int
	Count       int
	Pixels      [][]Color
	Progress    *pb.ProgressBar
	DoneChannel chan struct{}
}

func (request *RaytraceRowsRequest) Run() {
	for i := 0; i < request.Count; i++ {
		y := request.Start + i
		row := request.Rays[y]
		for x, ray := range row {
			request.Pixels[y][x] = castRay(request.Scene, ray, 0, 1)
			request.Progress.Increment()
		}
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
		kRefraction := 1 - closestSurface.Opacity()
		kReflection := closestSurface.Reflectivity() * closestSurface.Opacity()
		kDiffuse := 1 - kRefraction - kReflection
		var refractedColor, reflectedColor, diffuseColor Color

		if kRefraction > 0 {
			cosIn := -closestIntersection.Normal.Dot(ray.Direction)
			etaIn := refractionIndex
			etaOut := closestSurface.RefractiveIndex()
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
			refractedColor = castRay(scene, refractedRay, depth+1, closestSurface.RefractiveIndex())
		}

		if kReflection > 0 {
			reflectedDirection :=
				ray.Direction.Add(closestIntersection.Normal.Multiply(-2 * closestIntersection.Normal.Dot(ray.Direction)))

			// Bias the intersection point off the surface slightly to avoid immediate self-intersection.
			reflectedPoint := closestIntersection.Point.Translate(closestIntersection.Normal.Multiply(reflectionBias))

			reflectedRay := Ray{reflectedPoint, reflectedDirection.ToUnit()}
			reflectedColor = castRay(scene, reflectedRay, depth+1, refractionIndex)
		}

		if kDiffuse > 0 {
			for _, light := range scene.Lights {
				// Check if there is an object between the intersection point and the light source, in which case
				// it should cast a shadow.
				lightRay := Ray{
					Point:     closestIntersection.Point,
					Direction: light.Direction(closestIntersection.Point).Multiply(-1),
				}
				shadow := false
				for _, surface := range scene.Surfaces {
					if intersection := surface.Intersection(lightRay); intersection != nil {
						// Require a minimum distance to avoid a surface from shadowing itself.
						if intersection.Distance > shadowBias {
							if light.IsBlockedByIntersection(closestIntersection.Point, intersection) {
								shadow = true
								break
							}
						}
					}
				}
				if shadow {
					continue
				}

				incidentDotProduct :=
					light.Direction(closestIntersection.Point).Multiply(-1).Dot(closestIntersection.Normal)
				incidentLight := light.Intensity(closestIntersection.Point) * math.Max(incidentDotProduct, 0)
				diffuseColor.R += closestSurface.AlbedoAt(closestIntersection.Point).R / math.Pi * light.Color().R *
					incidentLight
				diffuseColor.G += closestSurface.AlbedoAt(closestIntersection.Point).G / math.Pi * light.Color().G *
					incidentLight
				diffuseColor.B += closestSurface.AlbedoAt(closestIntersection.Point).B / math.Pi * light.Color().B *
					incidentLight
			}
		}

		pixelColor.R = kRefraction*refractedColor.R + kReflection*reflectedColor.R + kDiffuse*diffuseColor.R
		pixelColor.G = kRefraction*refractedColor.G + kReflection*reflectedColor.G + kDiffuse*diffuseColor.G
		pixelColor.B = kRefraction*refractedColor.B + kReflection*reflectedColor.B + kDiffuse*diffuseColor.B
	}

	return pixelColor
}
