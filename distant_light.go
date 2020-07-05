package main

import "math/rand"

type DistantLight struct {
	direction          Vector
	color              Color
	intensity          float64
	directionVariation float64
	numSamples         int
}

func (light DistantLight) Direction(point Point, sampleNumber int) Vector {
	nominalDirection := light.direction.ToUnit()
	if light.directionVariation == 0 {
		return nominalDirection
	}

	// Adjust each component of the direction by a random factor.
	direction := nominalDirection
	direction.X += (rand.Float64() - 1) * light.directionVariation * nominalDirection.X
	direction.Y += (rand.Float64() - 1) * light.directionVariation * nominalDirection.Y
	direction.Z += (rand.Float64() - 1) * light.directionVariation * nominalDirection.Z

	return direction.ToUnit()
}

func (light DistantLight) Color() Color {
	return light.color
}

func (light DistantLight) Intensity(point Point) float64 {
	return light.intensity
}

func (light DistantLight) NumSamples() int {
	return light.numSamples
}

func (light DistantLight) IsBlockedByIntersection(point Point, intersection *Intersection) bool {
	// Intersecting distances are always closer than a distant light, which is infinitely far away.
	return true
}
