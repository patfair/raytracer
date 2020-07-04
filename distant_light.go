package main

type DistantLight struct {
	direction  Vector
	color      Color
	intensity  float64
	numSamples int
}

func (light DistantLight) Direction(point Point, sampleNumber int) Vector {
	return light.direction.ToUnit()
}

func (light DistantLight) Color() Color {
	return light.color
}

func (light DistantLight) Intensity(point Point) float64 {
	return light.intensity
}

func (light DistantLight) NumSamples() int {
	return light.numSamples
}

func (light DistantLight) IsBlockedByIntersection(point Point, intersection *Intersection) bool {
	// Intersecting distances are always closer than a distant light, which is infinitely far away.
	return true
}
