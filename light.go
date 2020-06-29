package main

type Light interface {
	Direction() Vector
	Color() Color
	Intensity() float64
}
