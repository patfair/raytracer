package main

import "math"

type Disc struct {
	plane Plane
}

func (disc Disc) Albedo() Color {
	return disc.plane.Albedo()
}

func (disc Disc) Intersection(ray Ray) *Intersection {
	denominator := disc.plane.Normal().Dot(ray.Direction.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.Point.VectorTo(disc.plane.Corner).Dot(disc.plane.Normal()) / denominator
	if distance < 0 {
		// The plane is behind the ray.
		return nil
	}
	point := ray.Point.Translate(ray.Direction.ToUnit().Multiply(distance))

	if !disc.isPointWithinLimits(point) {
		// The ray intersects outside the width and height of the plane.
		return nil
	}

	intersection := new(Intersection)
	intersection.Distance = distance
	intersection.Point = point
	intersection.Normal = disc.plane.Normal()
	if intersection.Normal.Dot(ray.Direction) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}

func (disc Disc) Radius() float64 {
	return disc.plane.Width.Norm()
}

func (disc Disc) isPointWithinLimits(point Point) bool {
	u, v := disc.plane.toTextureCoordinates(point)
	centerDistance := math.Sqrt(u*u + v*v)
	return centerDistance <= disc.Radius()
}
