package main

type Light interface {
	Direction(point Point, sampleNumber int) Vector
	Color() Color
	Intensity(point Point) float64
	NumSamples() int
	IsBlockedByIntersection(point Point, intersection *Intersection) bool
}
