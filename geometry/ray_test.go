// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertRayEqual(t *testing.T) {
	ray := Ray{Origin: Point{1, 2, 3}, Direction: Vector{-4, -5, -6}}
	AssertRayEqual(t, ray, ray)
}

func TestRay_String(t *testing.T) {
	ray := Ray{
		Origin:    Point{1.234, -14.567, 0.088888},
		Direction: Vector{0.123, 4.567, 8.9},
	}

	assert.Equal(t, "O:(1.23, -14.57, 0.09)|D:(0.12, 4.57, 8.90)", fmt.Sprintf("%v", ray))
}
