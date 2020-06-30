package main

type Plane struct {
	Corner Point
	Width  Vector
	Height Vector
	Color  Color
}

func (plane Plane) Normal() Vector {
	return plane.Width.Cross(plane.Height).ToUnit()
}

func (plane Plane) Albedo() Color {
	return plane.Color
}

func (plane Plane) Intersection(ray Ray) *Intersection {
	denominator := plane.Normal().Dot(ray.Direction.ToUnit())
	if denominator == 0 {
		// The ray is parallel to the plane; they do not intersect.
		return nil
	}

	distance := ray.Point.VectorTo(plane.Corner).Dot(plane.Normal()) / denominator
	if distance < 0 {
		// The plane is behind the ray.
		return nil
	}
	point := ray.Point.Translate(ray.Direction.ToUnit().Multiply(distance))

	if !plane.isPointWithinLimits(point) {
		// The ray intersects outside the width and height of the plane.
		return nil
	}

	intersection := new(Intersection)
	intersection.Distance = distance
	intersection.Point = point
	intersection.Normal = plane.Normal()
	if intersection.Normal.Dot(ray.Direction) > 0 {
		intersection.Normal = intersection.Normal.Multiply(-1)
	}

	return intersection
}

func (plane Plane) toTextureCoordinates(point Point) (float64, float64) {
	vector := plane.Corner.VectorTo(point)
	u := vector.Dot(plane.Width.ToUnit())
	v := vector.Dot(plane.Height.ToUnit())
	return u, v
}

func (plane Plane) isPointWithinLimits(point Point) bool {
	u, v := plane.toTextureCoordinates(point)
	width := plane.Width.Norm()
	height := plane.Height.Norm()
	return u >= 0 && u <= width && v >= 0 && v <= height
}
