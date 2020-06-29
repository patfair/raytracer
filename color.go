package main

import (
	imagecolor "image/color"
	"math"
)

type Color struct {
	R float64
	G float64
	B float64
}

func (color Color) ToRgba() imagecolor.RGBA {
	return imagecolor.RGBA{
		R: uint8(255 * math.Min(color.R, 1)),
		G: uint8(255 * math.Min(color.G, 1)),
		B: uint8(255 * math.Min(color.B, 1)),
		A: 255,
	}
}
