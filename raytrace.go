package main

import (
	"github.com/cheggaaa/pb/v3"
	"math"
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
			var closestIntersection *Intersection
			var closestSurface Surface
			for _, surface := range request.Scene.Surfaces {
				if intersection := surface.Intersection(ray); intersection != nil {
					if closestIntersection == nil || intersection.Distance < closestIntersection.Distance {
						closestIntersection = intersection
						closestSurface = surface
					}
				}
			}

			pixelColor := request.Scene.BackgroundColor
			if closestIntersection != nil {
				var color Color
				for _, light := range request.Scene.Lights {
					// Check if there is an object between the intersection point and the light source, in which case
					// it should cast a shadow.
					lightRay := Ray{
						Point:     closestIntersection.Point,
						Direction: light.Direction(closestIntersection.Point).Multiply(-1),
					}
					shadow := false
					for _, surface := range request.Scene.Surfaces {
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
					color.R += closestSurface.AlbedoAt(closestIntersection.Point).R / math.Pi * light.Color().R *
						incidentLight
					color.G += closestSurface.AlbedoAt(closestIntersection.Point).G / math.Pi * light.Color().G *
						incidentLight
					color.B += closestSurface.AlbedoAt(closestIntersection.Point).B / math.Pi * light.Color().B *
						incidentLight
				}
				pixelColor = color
			}
			request.Pixels[y][x] = pixelColor
			request.Progress.Increment()
		}
	}
	request.DoneChannel <- struct{}{}
}
