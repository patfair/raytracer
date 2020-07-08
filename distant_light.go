package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"math/rand"
)

type DistantLight struct {
	direction          geometry.Vector
	color              shading.Color
	intensity          float64
	directionVariation float64
	numSamples         int
}

func (light DistantLight) Direction(point geometry.Point, sampleNumber, numSamples int) geometry.Vector {
	nominalDirection := light.direction.ToUnit()
	if light.directionVariation == 0 || numSamples <= 1 {
		return nominalDirection
	}

	// Adjust each component of the direction by a random factor.
	direction := nominalDirection
	direction.X += (rand.Float64() - 1) * light.directionVariation * nominalDirection.X
	direction.Y += (rand.Float64() - 1) * light.directionVariation * nominalDirection.Y
	direction.Z += (rand.Float64() - 1) * light.directionVariation * nominalDirection.Z

	return direction.ToUnit()
}

func (light DistantLight) Color() shading.Color {
	return light.color
}

func (light DistantLight) Intensity(point geometry.Point) float64 {
	return light.intensity
}

func (light DistantLight) NumSamples() int {
	return light.numSamples
}

func (light DistantLight) IsBlockedByIntersection(point geometry.Point, intersection *geometry.Intersection) bool {
	// Intersecting distances are always closer than a distant light, which is infinitely far away.
	return true
}
