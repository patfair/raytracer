// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import (
	"math"
)

const checkerboardDither = 0.005

// Represents a texture that has two colors appearing in a checkerboard pattern.
type CheckerboardTexture struct {
	Color1 Color   // First color in the pattern
	Color2 Color   // Second color in the pattern
	UPitch float64 // Distance before the pattern repeats itself along the surface's texture coordinate U-axis
	VPitch float64 // Distance before the pattern repeats itself along the surface's texture coordinate Y-axis
}

// Returns either of the texture's colors, depending on the given coordinates.
func (texture CheckerboardTexture) AlbedoAt(u, v float64) Color {
	if getToggleValue(u, texture.UPitch) == getToggleValue(v, texture.VPitch) {
		return texture.Color1.Dither(checkerboardDither)
	}
	return texture.Color2.Dither(checkerboardDither)
}

func (texture CheckerboardTexture) NeedsTextureCoordinates() bool {
	return true
}

// Calculates the pitch fraction of the given position and returns true if it appears in the first half, and false if in
// the second half.
func getToggleValue(position, pitch float64) bool {
	_, fraction := math.Modf(position / pitch)
	if fraction < 0 {
		fraction += 1
	}
	return int(fraction*2) == 0
}
