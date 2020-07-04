package main

import (
	"math"
	"math/rand"
)

type PointLight struct {
	point      Point
	color      Color
	intensity  float64
	radius     float64
	numSamples int
}

func (light PointLight) Direction(point Point, sampleNumber int) Vector {
	nominalDirection := light.point.VectorTo(point).ToUnit()
	sampleNumber %= light.NumSamples()

	// Pick a arbitrary vectors normal to the nominal direction and to each other to represent the plane of the light.
	var uDirection Vector
	if nominalDirection.X != 0 {
		uDirection = Vector{-(nominalDirection.Y + nominalDirection.Z) / nominalDirection.X, 1, 1}.ToUnit()
	} else if nominalDirection.Y != 0 {
		uDirection = Vector{1, -(nominalDirection.X + nominalDirection.Z) / nominalDirection.Y, 1}.ToUnit()
	} else if nominalDirection.Z != 0 {
		uDirection = Vector{1, 1, -(nominalDirection.X + nominalDirection.Y) / nominalDirection.Z}.ToUnit()
	} else {
		// The light and the illuminated point are coincident.
		return nominalDirection
	}
	vDirection := nominalDirection.Cross(uDirection).ToUnit()

	// Randomize the point within the light's radius.
	r := (0.25 + 0.75*rand.Float64()) * light.radius
	phi := (float64(sampleNumber) + rand.Float64()) * 2 * math.Pi / float64(light.NumSamples())
	u := uDirection.Multiply(r * math.Cos(phi))
	v := vDirection.Multiply(r * math.Sin(phi))
	randomPoint := light.point.Translate(u).Translate(v)

	return randomPoint.VectorTo(point).ToUnit()
}

func (light PointLight) Color() Color {
	return light.color
}

func (light PointLight) Intensity(point Point) float64 {
	distance := light.point.VectorTo(point).Norm()
	sphereSurfaceArea := 4 * math.Pi * distance * distance
	return light.intensity / sphereSurfaceArea
}

func (light PointLight) NumSamples() int {
	return int(math.Max(float64(light.numSamples), 1))
}

func (light PointLight) IsBlockedByIntersection(point Point, intersection *Intersection) bool {
	// Intersecting distances don't block a point light if they are further away than the light from the ray origin.
	distance := light.point.VectorTo(point).Norm()
	return distance > intersection.Distance
}
