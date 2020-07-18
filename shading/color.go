// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import (
	imagecolor "image/color"
	"math"
	"math/rand"
)

// Represents a color with red, green, and blue floating-point components.
type Color struct {
	R float64 // Red component as a float in [0, 1]
	G float64 // Green component as a float in [0, 1]
	B float64 // Blue component as a float in [0, 1]
}

// Returns the color's RGBA representation (with A always set to 255).
func (color Color) ToRgba() imagecolor.RGBA {
	return imagecolor.RGBA{
		R: uint8(255 * math.Max(math.Min(color.R, 1), 0)),
		G: uint8(255 * math.Max(math.Min(color.G, 1), 0)),
		B: uint8(255 * math.Max(math.Min(color.B, 1), 0)),
		A: 255,
	}
}

// Returns a copy of the color with each component varied randomly up to plus or minus the given variation.
func (color Color) Dither(variation float64) Color {
	r := math.Max(math.Min(color.R+(2*rand.Float64()-1)*variation, 1), 0)
	g := math.Max(math.Min(color.G+(2*rand.Float64()-1)*variation, 1), 0)
	b := math.Max(math.Min(color.B+(2*rand.Float64()-1)*variation, 1), 0)
	return Color{r, g, b}
}
