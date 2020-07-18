// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

const solidDither = 0.005

// Represents a texture that has one uniform and solid diffuse color.
type SolidTexture struct {
	Color Color // Single solid color of the texture
}

// Returns the same solid color at all texture coordinates.
func (texture SolidTexture) AlbedoAt(u, v float64) Color {
	return texture.Color.Dither(solidDither)
}

func (texture SolidTexture) NeedsTextureCoordinates() bool {
	return false
}
