// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

// Represents a texture that has one uniform and solid diffuse color.
type SolidTexture struct {
	Color Color // Single solid color of the texture
}

// Returns the same solid color at all texture coordinates.
func (texture SolidTexture) AlbedoAt(u, v, ditherVariation float64) Color {
	return texture.Color.Dither(ditherVariation)
}

func (texture SolidTexture) NeedsTextureCoordinates() bool {
	return false
}
