// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package shading

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShadingProperties_Validate(t *testing.T) {
	shadingProperties := ShadingProperties{
		DiffuseTexture:    nil,
		SpecularExponent:  0,
		SpecularIntensity: 0,
		Opacity:           0,
		Reflectivity:      0,
		RefractiveIndex:   1,
	}

	err := shadingProperties.Validate()
	assert.Nil(t, err)

	shadingProperties.SpecularExponent = -1
	err = shadingProperties.Validate()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "exponent must be non-negative")
	}
	shadingProperties.SpecularExponent = 0

	shadingProperties.SpecularIntensity = -1
	err = shadingProperties.Validate()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "intensity must be non-negative")
	}
	shadingProperties.SpecularIntensity = 0

	shadingProperties.Opacity = -0.1
	err = shadingProperties.Validate()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "opacity must be in [0, 1]")
	}
	shadingProperties.Opacity = 1.1
	err = shadingProperties.Validate()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "opacity must be in [0, 1]")
	}
	shadingProperties.Opacity = 1

	shadingProperties.Reflectivity = -0.1
	err = shadingProperties.Validate()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "reflectivity must be in [0, 1]")
	}
	shadingProperties.Reflectivity = 1.1
	err = shadingProperties.Validate()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "reflectivity must be in [0, 1]")
	}
	shadingProperties.Reflectivity = 1

	shadingProperties.RefractiveIndex = 0.9
	err = shadingProperties.Validate()
	assert.Nil(t, err)
	shadingProperties.Opacity = 0.4
	err = shadingProperties.Validate()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "index must be at least 1")
	}
	shadingProperties.RefractiveIndex = 1
}
