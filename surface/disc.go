// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"errors"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"math"
)

// Represents a two-sided, circular surface having a finite radius and zero thickness.
type Disc struct {
	// Internal plane parallel to the disc, whose corner indicates the center of the disc and whose width/height vectors
	// specify the disc diameter.
	plane Plane
}

// Returns a new plane, or an error if the parameters are invalid.
func NewDisc(center geometry.Point, width geometry.Vector, height geometry.Vector,
	shadingProperties shading.ShadingProperties) (Disc, error) {
	if width.Dot(height) != 0 {
		return Disc{}, errors.New("disc width and height must be perpendicular")
	}
	if width.Norm() != height.Norm() {
		return Disc{}, errors.New("disc width and height must have the same magnitude")
	}

	plane, err := NewPlane(center, width, height, shadingProperties)
	return Disc{plane: plane}, err
}

func (disc Disc) Intersection(ray geometry.Ray) *geometry.Intersection {
	denominator := disc.plane.normal().Dot(ray.Direction.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.Origin.VectorTo(disc.plane.bottomLeftCorner).Dot(disc.plane.normal()) / denominator
	if distance < 0 {
		// The plane is behind the ray.
		return nil
	}
	point := ray.Origin.Translate(ray.Direction.ToUnit().Multiply(distance))

	if !disc.isPointWithinLimits(point) {
		// The ray intersects outside the width and height of the plane.
		return nil
	}

	intersection := new(geometry.Intersection)
	intersection.Distance = distance
	intersection.Point = point
	intersection.Normal = disc.plane.normal()
	if intersection.Normal.Dot(ray.Direction) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}

func (disc Disc) ShadingProperties() shading.ShadingProperties {
	return disc.plane.ShadingProperties()
}

func (disc Disc) ToTextureCoordinates(point geometry.Point) (float64, float64) {
	// Convert first to planar coordinates.
	vector := disc.plane.bottomLeftCorner.VectorTo(point)
	u := vector.Dot(disc.plane.width.ToUnit())
	v := vector.Dot(disc.plane.height.ToUnit())

	// Convert planar to polar coordinates.
	r := math.Sqrt(u*u + v*v)
	phi := math.Atan2(v, u)
	return r, phi
}

// Returns true if the given disc on the plane in world coordinates is within the defined boundaries of the disc.
func (disc Disc) isPointWithinLimits(point geometry.Point) bool {
	u, v := disc.plane.ToTextureCoordinates(point)
	centerDistance := math.Sqrt(u*u + v*v)
	return centerDistance <= disc.radius()
}

func (disc Disc) radius() float64 {
	return disc.plane.width.Norm()
}
