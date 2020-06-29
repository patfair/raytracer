package main

type DistantLight struct {
	direction Vector
	color     Color
	intensity float64
}

func (light DistantLight) Direction(point Point) Vector {
	return light.direction.ToUnit()
}

func (light DistantLight) Color() Color {
	return light.color
}

func (light DistantLight) Intensity(point Point) float64 {
	return light.intensity
}

func (light DistantLight) IsBlockedByIntersection(point Point, intersection *Intersection) bool {
	// Intersecting distances are always closer than a distant light, which is infinitely far away.
	return true
}
