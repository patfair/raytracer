// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package light

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
)

// Represents a source of light within a set.
type Light interface {
	// Determines the direction of the light incident to the given point. Depending on the light source's properties and
	// the value of the sampling arguments, the direction may include some random variation in order to produce soft
	// shadows.
	Direction(point geometry.Point, sampleNumber, numSamples int) geometry.Vector

	// Returns the color of the light produced by the light source.
	Color() shading.Color

	// Determines the intensity of the light incident to the given point.
	Intensity(point geometry.Point) float64

	// Determines whether the given point is blocked from receiving light by the surface whose intersection with the
	// light ray is given.
	IsBlockedByIntersection(point geometry.Point, intersection *geometry.Intersection) bool
}
