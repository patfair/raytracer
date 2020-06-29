package main

import "math"

type PointLight struct {
	point     Point
	color     Color
	intensity float64
}

func (light PointLight) Direction(point Point) Vector {
	return light.point.VectorTo(point).ToUnit()
}

func (light PointLight) Color() Color {
	return light.color
}

func (light PointLight) Intensity(point Point) float64 {
	distance := light.point.VectorTo(point).Norm()
	sphereSurfaceArea := 4 * math.Pi * distance * distance
	return light.intensity / sphereSurfaceArea
}

func (light PointLight) IsBlockedByIntersection(point Point, intersection *Intersection) bool {
	// Intersecting distances don't block a point light if they are further away than the light from the ray origin.
	distance := light.point.VectorTo(point).Norm()
	return distance > intersection.Distance
}
