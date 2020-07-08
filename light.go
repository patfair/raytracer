package main

import "github.com/patfair/raytracer/geometry"

type Light interface {
	Direction(point geometry.Point, sampleNumber, numSamples int) geometry.Vector
	Color() Color
	Intensity(point geometry.Point) float64
	NumSamples() int
	IsBlockedByIntersection(point geometry.Point, intersection *geometry.Intersection) bool
}
