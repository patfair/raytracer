// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"errors"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
)

// Represents a two-sided, rectangular surface having a finite size and zero thickness.
type Plane struct {
	bottomLeftCorner  geometry.Point  // Point representing the bottom left corner of the plane
	width             geometry.Vector // Direction and size of the plane extending "right" from the corner
	height            geometry.Vector // Direction and size of the plane extending "up" from the corner
	shadingProperties shading.ShadingProperties
}

// Returns a new plane, or an error if the parameters are invalid.
func NewPlane(bottomLeftCorner geometry.Point, width, height geometry.Vector,
	shadingProperties shading.ShadingProperties) (Plane, error) {
	if width.Dot(height) != 0 {
		return Plane{}, errors.New("plane width and height must be perpendicular")
	}

	return Plane{
		bottomLeftCorner:  bottomLeftCorner,
		width:             width,
		height:            height,
		shadingProperties: shadingProperties,
	}, nil
}

func (plane Plane) Intersection(ray geometry.Ray) *geometry.Intersection {
	denominator := plane.normal().Dot(ray.Direction.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.Origin.VectorTo(plane.bottomLeftCorner).Dot(plane.normal()) / denominator
	if distance < 0 {
		// The plane is behind the ray.
		return nil
	}
	point := ray.Origin.Translate(ray.Direction.ToUnit().Multiply(distance))

	if !plane.isPointWithinLimits(point) {
		// The ray intersects outside the width and height of the plane.
		return nil
	}

	intersection := new(geometry.Intersection)
	intersection.Distance = distance
	intersection.Point = point
	intersection.Normal = plane.normal()
	if intersection.Normal.Dot(ray.Direction) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}

func (plane Plane) ShadingProperties() shading.ShadingProperties {
	return plane.shadingProperties
}

func (plane Plane) ToTextureCoordinates(point geometry.Point) (float64, float64) {
	vector := plane.bottomLeftCorner.VectorTo(point)
	u := vector.Dot(plane.width.ToUnit())
	v := vector.Dot(plane.height.ToUnit())
	return u, v
}

// Returns true if the given point on the plane in world coordinates is within the defined boundaries of the plane.
func (plane Plane) isPointWithinLimits(point geometry.Point) bool {
	u, v := plane.ToTextureCoordinates(point)
	width := plane.width.Norm()
	height := plane.height.Norm()
	return u >= 0 && u <= width && v >= 0 && v <= height
}

// Returns a unit vector representing the direction that is normal to the plane's surface.
func (plane Plane) normal() geometry.Vector {
	return plane.width.Cross(plane.height).ToUnit()
}
