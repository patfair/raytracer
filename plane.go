package main

type Plane struct {
	Point
	Vector
	Color
}

func (plane Plane) Albedo() Color {
	return plane.Color
}

func (plane Plane) Intersection(ray Ray) *Intersection {
	denominator := plane.Vector.ToUnit().Dot(ray.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.VectorTo(plane.Point).Dot(plane.Vector.ToUnit()) / denominator
	if distance < 0 {
		// The plane is behind the ray.
		return nil
	}

	intersection := new(Intersection)
	intersection.Distance = distance
	intersection.Point = ray.Point.Translate(ray.Vector.ToUnit().Multiply(intersection.Distance))
	intersection.Normal = plane.Vector
	if intersection.Normal.Dot(ray.Vector) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}
