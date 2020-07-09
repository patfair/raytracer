// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
)

// Represents a physical surface that a ray of light can intersect and interact with in order to determine its shading.
type Surface interface {
	// Determines whether the given ray intersects the surface, and if so, returns the details for the intersection that
	// is closest to the origin of the ray. A nil return value indicates that the ray and surface do not intersect.
	Intersection(ray geometry.Ray) *geometry.Intersection

	// Returns the properties necessary for determining how the surface should be shaded.
	ShadingProperties() shading.ShadingProperties

	// Converts the given point in world coordinates on the surface to the equivalent (U, V) texture coordinates.
	// Garbage output may be produced for an input point not actually on the surface.
	ToTextureCoordinates(point geometry.Point) (float64, float64)
}
