package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"math"
)

type Disc struct {
	plane Plane
}

func (disc Disc) Intersection(ray geometry.Ray) *geometry.Intersection {
	denominator := disc.plane.Normal().Dot(ray.Direction.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.Origin.VectorTo(disc.plane.Corner).Dot(disc.plane.Normal()) / denominator
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
	intersection.Normal = disc.plane.Normal()
	if intersection.Normal.Dot(ray.Direction) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}

func (disc Disc) AlbedoAt(point geometry.Point) shading.Color {
	r, phi := disc.toTextureCoordinates(point)
	return disc.plane.ShadingProperties().DiffuseTexture.AlbedoAt(r, phi)
}

func (disc Disc) Radius() float64 {
	return disc.plane.Width.Norm()
}

func (disc Disc) ShadingProperties() shading.ShadingProperties {
	return disc.plane.ShadingProperties()
}

func (disc Disc) toTextureCoordinates(point geometry.Point) (float64, float64) {
	// Convert first to planar coordinates.
	vector := disc.plane.Corner.VectorTo(point)
	u := vector.Dot(disc.plane.Width.ToUnit())
	v := vector.Dot(disc.plane.Height.ToUnit())

	// Convert planar to polar coordinates.
	r := math.Sqrt(u*u + v*v)
	phi := math.Atan2(v, u)
	return r, phi
}

func (disc Disc) isPointWithinLimits(point geometry.Point) bool {
	u, v := disc.plane.toTextureCoordinates(point)
	centerDistance := math.Sqrt(u*u + v*v)
	return centerDistance <= disc.Radius()
}
