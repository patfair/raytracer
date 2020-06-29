package main

type Plane struct {
	Point  Point
	Normal Vector
	Color  Color
}

func (plane Plane) Albedo() Color {
	return plane.Color
}

func (plane Plane) Intersection(ray Ray) *Intersection {
	denominator := plane.Normal.ToUnit().Dot(ray.Direction.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.Point.VectorTo(plane.Point).Dot(plane.Normal.ToUnit()) / denominator
	if distance < 0 {
		// The plane is behind the ray.
		return nil
	}

	intersection := new(Intersection)
	intersection.Distance = distance
	intersection.Point = ray.Point.Translate(ray.Direction.ToUnit().Multiply(intersection.Distance))
	intersection.Normal = plane.Normal
	if intersection.Normal.Dot(ray.Direction) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}
