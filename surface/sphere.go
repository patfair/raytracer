// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"errors"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"math"
)

// Represents a spherical surface.
type Sphere struct {
	center            geometry.Point  // Point at which the sphere is centered
	radius            float64         // Radius of the sphere
	zenithReference   geometry.Vector // For texture mapping, vector representing the axis of rotation
	azimuthReference  geometry.Vector // For texture mapping, vector pointing to a start point along the equator
	shadingProperties shading.ShadingProperties
}

// Returns a new sphere, or an error if the parameters are invalid.
func NewSphere(center geometry.Point, radius float64, zenithReference, azimuthReference geometry.Vector,
	shadingProperties shading.ShadingProperties) (Sphere, error) {
	if err := shadingProperties.Validate(); err != nil {
		return Sphere{}, err
	}
	if radius <= 0 {
		return Sphere{}, errors.New("radius must be positive")
	}
	if zenithReference.Dot(azimuthReference) != 0 {
		return Sphere{}, errors.New("zenith and azimuth references must be perpendicular")
	}

	return Sphere{
		center:            center,
		radius:            radius,
		zenithReference:   zenithReference,
		azimuthReference:  azimuthReference,
		shadingProperties: shadingProperties,
	}, nil
}

func (sphere Sphere) Intersection(ray geometry.Ray) *geometry.Intersection {
	rayOriginToSphereCenter := ray.Origin.VectorTo(sphere.center)
	midpointDistance := ray.Direction.ToUnit().Dot(rayOriginToSphereCenter)
	if midpointDistance < 0 {
		// The sphere is behind the ray; there is no intersection.
		return nil
	}

	radiusSquared := sphere.radius * sphere.radius
	rayDistanceSquared := rayOriginToSphereCenter.Dot(rayOriginToSphereCenter) - midpointDistance*midpointDistance
	if rayDistanceSquared > radiusSquared {
		// The ray passes outside the sphere.
		return nil
	}

	halfChordDistance := math.Sqrt(radiusSquared - rayDistanceSquared)
	closestIntersectionDistance := midpointDistance - halfChordDistance
	closestIntersectionPoint := ray.Origin.Translate(ray.Direction.ToUnit().Multiply(closestIntersectionDistance))
	normal := sphere.center.VectorTo(closestIntersectionPoint).ToUnit()

	return &geometry.Intersection{
		Point:    closestIntersectionPoint,
		Distance: closestIntersectionDistance,
		Normal:   normal,
	}
}

func (sphere Sphere) ShadingProperties() shading.ShadingProperties {
	return sphere.shadingProperties
}

func (sphere Sphere) ToTextureCoordinates(point geometry.Point) (float64, float64) {
	// Convert first to rectangular coordinates relative to the zenith and azimuth.
	vector := sphere.center.VectorTo(point)
	uDirection := sphere.azimuthReference.ToUnit()
	wDirection := sphere.zenithReference.ToUnit()
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
