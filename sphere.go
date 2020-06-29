package main

import (
	"math"
)

type Sphere struct {
	Center Point
	Radius float64
	Color  Color
}

func (sphere Sphere) Albedo() Color {
	return sphere.Color
}

func (sphere Sphere) Intersection(ray Ray) *Intersection {
	rayOriginToSphereCenter := ray.Point.VectorTo(sphere.Center)
	midpointDistance := ray.Direction.ToUnit().Dot(rayOriginToSphereCenter)
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
	closestIntersectionPoint := ray.Point.Translate(ray.Direction.ToUnit().Multiply(closestIntersectionDistance))
	normal := sphere.Center.VectorTo(closestIntersectionPoint).ToUnit()

	return &Intersection{
		Point:    closestIntersectionPoint,
		Distance: closestIntersectionDistance,
		Normal:   normal,
	}
}
