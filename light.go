package main

type Light interface {
	Direction(point Point) Vector
	Color() Color
	Intensity(point Point) float64
	IsBlockedByIntersection(point Point, intersection *Intersection) bool
}
