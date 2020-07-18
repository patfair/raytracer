// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package light

import (
	"errors"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"math/rand"
)

// Represents a light source for which the incident direction and intensity are the same regardless of the location of
// the illuminated point (such as the sun, for all practical purposes).
type DistantLight struct {
	direction          geometry.Vector // Fixed direction of the light rays leaving the light source
	color              shading.Color
	intensity          float64 // Fixed intensity of the light rays leaving the light source
	directionVariation float64 // Value in [0, 1] by which each component of the direction can be randomly varied
}

func NewDistantLight(direction geometry.Vector, color shading.Color, intensity float64,
	directionVariation float64) (DistantLight, error) {
	if intensity <= 0 {
		return DistantLight{}, errors.New("intensity must be positive")
	}
	if directionVariation < 0 || directionVariation > 1 {
		return DistantLight{}, errors.New("direction variation must be in [0, 1]")
	}

	return DistantLight{
		direction:          direction,
		color:              color,
		intensity:          intensity,
		directionVariation: directionVariation,
	}, nil
}

func (light DistantLight) Direction(point geometry.Point, sampleNumber, numSamples int) geometry.Vector {
	nominalDirection := light.direction.ToUnit()
	if light.directionVariation == 0 || numSamples <= 1 {
		return nominalDirection
	}

	// Adjust each component of the direction by a random factor.
	direction := nominalDirection
	direction.X += (2*rand.Float64() - 1) * light.directionVariation
	direction.Y += (2*rand.Float64() - 1) * light.directionVariation
	direction.Z += (2*rand.Float64() - 1) * light.directionVariation

	return direction.ToUnit()
}

func (light DistantLight) Color() shading.Color {
	return light.color
}

func (light DistantLight) Intensity(point geometry.Point) float64 {
	return light.intensity
}

func (light DistantLight) IsBlockedByIntersection(point geometry.Point, intersection *geometry.Intersection) bool {
	// Intersecting distances are always closer than a distant light, which is infinitely far away.
	return true
}
