package main

type DistantLight struct {
	direction Vector
	color     Color
	intensity float64
}

func (light DistantLight) Direction() Vector {
	return light.direction.ToUnit()
}

func (light DistantLight) Color() Color {
	return light.color
}

func (light DistantLight) Intensity() float64 {
	return light.intensity
}
