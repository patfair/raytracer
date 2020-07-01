package main

import (
	"math"
)

type Sphere struct {
	Center           Point
	Radius           float64
	ZenithReference  Vector
	AzimuthReference Vector
	Texture          Texture
	reflection       float64
	refraction       float64
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

func (sphere Sphere) AlbedoAt(point Point) Color {
	theta, phi := sphere.toTextureCoordinates(point)
	return sphere.Texture.AlbedoAt(theta, phi)
}

func (sphere Sphere) Reflection() float64 {
	return sphere.reflection
}

func (sphere Sphere) Refraction() float64 {
	return sphere.refraction
}

func (sphere Sphere) toTextureCoordinates(point Point) (float64, float64) {
	// Convert first to rectangular coordinates relative to the zenith and azimuth.
	vector := sphere.Center.VectorTo(point)
	uDirection := sphere.AzimuthReference.ToUnit()
	wDirection := sphere.ZenithReference.ToUnit()
	vDirection := wDirection.Cross(uDirection)
	u := vector.Dot(uDirection)
	v := vector.Dot(vDirection)
	w := vector.Dot(wDirection)

	// Convert rectangular to spherical coordinates.
	r := math.Sqrt(u*u + v*v + w*w)
	theta := math.Atan2(v, u)
	phi := math.Acos(w / r)
	return theta, phi
}
