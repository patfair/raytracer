package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
)

type Plane struct {
	Corner            geometry.Point
	Width             geometry.Vector
	Height            geometry.Vector
	shadingProperties shading.ShadingProperties
}

func (plane Plane) Normal() geometry.Vector {
	return plane.Width.Cross(plane.Height).ToUnit()
}

func (plane Plane) Intersection(ray geometry.Ray) *geometry.Intersection {
	denominator := plane.Normal().Dot(ray.Direction.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.Origin.VectorTo(plane.Corner).Dot(plane.Normal()) / denominator
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
	intersection.Normal = plane.Normal()
	if intersection.Normal.Dot(ray.Direction) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}

func (plane Plane) AlbedoAt(point geometry.Point) shading.Color {
	u, v := plane.toTextureCoordinates(point)
	return plane.shadingProperties.DiffuseTexture.AlbedoAt(u, v)
}

func (plane Plane) ShadingProperties() shading.ShadingProperties {
	return plane.shadingProperties
}

func (plane Plane) toTextureCoordinates(point geometry.Point) (float64, float64) {
	vector := plane.Corner.VectorTo(point)
	u := vector.Dot(plane.Width.ToUnit())
	v := vector.Dot(plane.Height.ToUnit())
	return u, v
}

func (plane Plane) isPointWithinLimits(point geometry.Point) bool {
	u, v := plane.toTextureCoordinates(point)
	width := plane.Width.Norm()
	height := plane.Height.Norm()
	return u >= 0 && u <= width && v >= 0 && v <= height
}
