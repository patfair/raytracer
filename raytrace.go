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
			request.Pixels[y][x] = castRay(request.Scene, ray, 0)
			request.Progress.Increment()
		}
	}
	request.DoneChannel <- struct{}{}
}

func castRay(scene *Scene, ray Ray, depth int) Color {
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
		reflectionCoefficient := math.Min(closestSurface.Reflection(), 1)
		var diffuseColor Color
		if reflectionCoefficient < 1 {
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

		var reflectedColor Color
		if reflectionCoefficient > 0 {
			reflectedDirection :=
				ray.Direction.Add(closestIntersection.Normal.Multiply(-2 * closestIntersection.Normal.Dot(ray.Direction)))

			// Bias the intersection point off the surface slightly to avoid immediate self-intersection.
			reflectedPoint := closestIntersection.Point.Translate(closestIntersection.Normal.Multiply(reflectionBias))

			reflectedRay := Ray{reflectedPoint, reflectedDirection.ToUnit()}
			reflectedColor = castRay(scene, reflectedRay, depth+1)
		}

		pixelColor.R = (1-reflectionCoefficient)*diffuseColor.R + reflectionCoefficient*reflectedColor.R
		pixelColor.G = (1-reflectionCoefficient)*diffuseColor.G + reflectionCoefficient*reflectedColor.G
		pixelColor.B = (1-reflectionCoefficient)*diffuseColor.B + reflectionCoefficient*reflectedColor.B
	}

	return pixelColor
}
