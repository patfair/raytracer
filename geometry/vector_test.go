// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestVector_Norm(t *testing.T) {
	assert.Equal(t, 3.0, Vector{3, 0, 0}.Norm())
	assert.Equal(t, 2.0, Vector{0, -2, 0}.Norm())
	assert.Equal(t, 1.0, Vector{0, 0, 1}.Norm())
	assert.Equal(t, math.Sqrt(14), Vector{-1, 2, -3}.Norm())
}

func TestVector_ToUnit(t *testing.T) {
	assert.Equal(t, Vector{1, 0, 0}, Vector{3, 0, 0}.ToUnit())
	assert.Equal(t, Vector{0, -1, 0}, Vector{0, -2, 0}.ToUnit())
	assert.Equal(t, Vector{0, 0, 1}, Vector{0, 0, 1}.ToUnit())
	assert.Equal(t, Vector{-1 / math.Sqrt(14), 2 / math.Sqrt(14), -3 / math.Sqrt(14)}, Vector{-1, 2, -3}.ToUnit())
}

func TestVector_Add(t *testing.T) {
	vector1 := Vector{-4.1, -5.2, 1}
	vector2 := Vector{0, 1.5, -2.5}

	assert.Equal(t, Vector{-4.1, -3.7, -1.5}, vector1.Add(vector2))
	assert.Equal(t, Vector{-4.1, -3.7, -1.5}, vector2.Add(vector1))
}

func TestVector_Multiply(t *testing.T) {
	assert.Equal(t, Vector{1.5, -3, 4.5}, Vector{-1, 2, -3}.Multiply(-1.5))
}

func TestVector_Dot(t *testing.T) {
	vector1 := Vector{-4.1, -5.2, 1}
	vector2 := Vector{0, 1.5, -2.5}

	assert.Equal(t, -10.3, vector1.Dot(vector2))
	assert.Equal(t, -10.3, vector2.Dot(vector1))

	assert.Equal(t, 0.0, Vector{1, 1, 0}.Dot(Vector{0, 0, -1}))
}

func TestVector_Cross(t *testing.T) {
	assert.Equal(t, Vector{0, 0, 6}, Vector{2, 0, 0}.Cross(Vector{0, 3, 0}))
	assert.Equal(t, Vector{0, 0, -6}, Vector{0, 3, 0}.Cross(Vector{2, 0, 0}))
	assert.Equal(t, Vector{0, 0, 0}, Vector{1, 0, 0}.Cross(Vector{-1, 0, 0}))
	assert.Equal(t, Vector{-7, -14, -7}, Vector{-1, 2, -3}.Cross(Vector{6, -5, 4}))
}

func TestVector_String(t *testing.T) {
	vector := Vector{-1.2, 3, 4.56}
	assert.Equal(t, "(-1.20, 3.00, 4.56)", fmt.Sprintf("%v", vector))
}
