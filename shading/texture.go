// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

// Interface for determining the amount of diffuse light reflected at a given point on a surface.
type Texture interface {
	// Returns the diffuse color that the texture should have at the given point in texture coordinates.
	AlbedoAt(u, v, ditherVariation float64) Color

	// Returns whether the specific texture implementation is independent of coordinates.
	NeedsTextureCoordinates() bool
}
