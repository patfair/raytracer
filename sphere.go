package main

import (
	"math"
)

type Sphere struct {
	Point
	Radius float64
	Color
}

func (sphere Sphere) Albedo() Color {
	return sphere.Color
}

func (sphere Sphere) Intersection(ray Ray) *Intersection {
	rayOriginToSphereCenter := ray.Point.VectorTo(sphere.Point)
	midpointDistance := ray.Vector.ToUnit().Dot(rayOriginToSphereCenter)
	if midpointDistance < 0 {
		// The sphere is behind the ray; there is no intersection.
		return nil
	}

	radiusSquared := sphere.Radius * sphere.Radius
	rayDistanceSquared := rayOriginToSphereCenter.Dot(rayOriginToSphereCenter) - midpointDistance*midpointDistance
	if rayDistanceSquared > radiusSquared {
		// The ray passes outside the sphere.
		return nil
	}

	halfChordDistance := math.Sqrt(radiusSquared - rayDistanceSquared)
	closestIntersectionDistance := midpointDistance - halfChordDistance
	closestIntersectionPoint := ray.Point.Translate(ray.Vector.ToUnit().Multiply(closestIntersectionDistance))
	normal := sphere.Point.VectorTo(closestIntersectionPoint).ToUnit()

	return &Intersection{
		Point:    closestIntersectionPoint,
		Distance: closestIntersectionDistance,
		Normal:   normal,
	}
}
