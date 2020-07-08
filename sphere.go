package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"math"
)

type Sphere struct {
	Center            geometry.Point
	Radius            float64
	ZenithReference   geometry.Vector
	AzimuthReference  geometry.Vector
	shadingProperties shading.ShadingProperties
}

func (sphere Sphere) Intersection(ray geometry.Ray) *geometry.Intersection {
	rayOriginToSphereCenter := ray.Origin.VectorTo(sphere.Center)
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
	closestIntersectionPoint := ray.Origin.Translate(ray.Direction.ToUnit().Multiply(closestIntersectionDistance))
	normal := sphere.Center.VectorTo(closestIntersectionPoint).ToUnit()

	return &geometry.Intersection{
		Point:    closestIntersectionPoint,
		Distance: closestIntersectionDistance,
		Normal:   normal,
	}
}

func (sphere Sphere) AlbedoAt(point geometry.Point) shading.Color {
	theta, phi := sphere.toTextureCoordinates(point)
	return sphere.shadingProperties.DiffuseTexture.AlbedoAt(theta, phi)
}

func (sphere Sphere) ShadingProperties() shading.ShadingProperties {
	return sphere.shadingProperties
}

func (sphere Sphere) toTextureCoordinates(point geometry.Point) (float64, float64) {
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
