// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Asserts equality of the two given rays, within a small allowable error.
func AssertRayEqual(t *testing.T, expected, actual Ray) {
	assert.Equal(t, expected.Origin, actual.Origin)
	AssertVectorEqual(t, expected.Direction, actual.Direction)
}

// Asserts equality of the two given vectors, within a small allowable error.
func AssertVectorEqual(t *testing.T, expected, actual Vector) {
	epsilon := 0.001
	assert.InDelta(t, expected.X, actual.X, epsilon, "X expected: %v, actual: %v", expected.X, actual.X)
	assert.InDelta(t, expected.Y, actual.Y, epsilon, "Y expected: %v, actual: %v", expected.Y, actual.Y)
	assert.InDelta(t, expected.Z, actual.Z, epsilon, "Z expected: %v, actual: %v", expected.Z, actual.Z)
}
