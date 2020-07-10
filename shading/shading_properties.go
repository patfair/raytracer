// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import "errors"

// Holds all the properties necessary for determining how a surface should be shaded.
type ShadingProperties struct {
	DiffuseTexture    Texture // Interface for determining the albedo (color) of the surface at a given point.
	SpecularExponent  float64 // Dimensionless property for tuning the size of the specular reflection
	SpecularIntensity float64 // Dimensionless property for tuning the brightness of the specular reflection
	Opacity           float64 // What proportion of light that the surface blocks as a value in [0, 1]
	Reflectivity      float64 // What proportion of light that the surface reflects as a value in [0, 1]
	RefractiveIndex   float64 // For a surface that is not fully opaque, specifies how fast light travels through it
}

func (properties ShadingProperties) Validate() error {
	if properties.SpecularExponent < 0 {
		return errors.New("specular exponent must be non-negative")
	}
	if properties.SpecularIntensity < 0 {
		return errors.New("specular intensity must be non-negative")
	}
	if properties.Opacity < 0 || properties.Opacity > 1 {
		return errors.New("opacity must be in [0, 1]")
	}
	if properties.Reflectivity < 0 || properties.Reflectivity > 1 {
		return errors.New("reflectivity must be in [0, 1]")
	}
	if properties.Opacity < 1 && properties.RefractiveIndex < 1 {
		return errors.New("refractive index must be at least 1 if not fully opaque")
	}

	return nil
}
