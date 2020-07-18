// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package light

import (
	"errors"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"math"
	"math/rand"
)

// Represents a light source located at a specific point emitting light in all directions.
type PointLight struct {
	point     geometry.Point // Point at which the light is located
	color     shading.Color
	intensity float64 // Intensity of the light at the source
	radius    float64 // Radius in which the source of the light can be randomly varied
}

func NewPointLight(point geometry.Point, color shading.Color, intensity float64, radius float64) (PointLight, error) {
	if intensity <= 0 {
		return PointLight{}, errors.New("intensity must be positive")
	}
	if radius < 0 {
		return PointLight{}, errors.New("radius must be non-negative")
	}

	return PointLight{
		point:     point,
		color:     color,
		intensity: intensity,
		radius:    radius,
	}, nil
}

func (light PointLight) Direction(point geometry.Point, sampleNumber, numSamples int) geometry.Vector {
	nominalDirection := light.point.VectorTo(point).ToUnit()
	if numSamples <= 1 {
		return nominalDirection
	}
	sampleNumber %= numSamples

	// Pick a arbitrary vectors normal to the nominal direction and to each other to represent the plane of the light.
	var uDirection geometry.Vector
	if nominalDirection.X != 0 {
		uDirection = geometry.Vector{-(nominalDirection.Y + nominalDirection.Z) / nominalDirection.X, 1, 1}.ToUnit()
	} else if nominalDirection.Y != 0 {
		uDirection = geometry.Vector{1, -(nominalDirection.X + nominalDirection.Z) / nominalDirection.Y, 1}.ToUnit()
	} else if nominalDirection.Z != 0 {
		uDirection = geometry.Vector{1, 1, -(nominalDirection.X + nominalDirection.Y) / nominalDirection.Z}.ToUnit()
	} else {
		// The light and the illuminated point are coincident.
		return nominalDirection
	}
	vDirection := nominalDirection.Cross(uDirection).ToUnit()

	// Randomize the point within the light's radius.
	r := light.radius * math.Sqrt(rand.Float64())
	phi := (float64(sampleNumber) + rand.Float64()) * 2 * math.Pi / float64(numSamples)
	u := uDirection.Multiply(r * math.Cos(phi))
	v := vDirection.Multiply(r * math.Sin(phi))
	randomPoint := light.point.Translate(u).Translate(v)

	return randomPoint.VectorTo(point).ToUnit()
}

func (light PointLight) Color() shading.Color {
	return light.color
}

func (light PointLight) Intensity(point geometry.Point) float64 {
	distance := light.point.DistanceTo(point)
	sphereSurfaceArea := 4 * math.Pi * distance * distance
	return light.intensity / sphereSurfaceArea
}

func (light PointLight) IsBlockedByIntersection(point geometry.Point, intersection *geometry.Intersection) bool {
	// Intersecting distances don't block a point light if they are further away than the light from the ray origin.
	distance := light.point.DistanceTo(point)
	return distance > intersection.Distance
}
