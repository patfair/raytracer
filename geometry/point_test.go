// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoint_DistanceTo(t *testing.T) {
	point1 := Point{-2.5, 0.1, 12}
	point2 := Point{0.5, -11.9, 8}

	assert.Equal(t, 13.0, point1.DistanceTo(point2))
	assert.Equal(t, 13.0, point2.DistanceTo(point1))
}

func TestPoint_Translate(t *testing.T) {
	point1 := Point{-1, 2, -3}
	vector1 := Vector{-4.1, -5.2, 0}
	vector2 := Vector{0, 1.5, -2.5}

	point2 := point1.Translate(vector1)
	assert.Equal(t, Point{-5.1, -3.2, -3}, point2)

	point3 := point1.Translate(vector2)
	assert.Equal(t, Point{-1, 3.5, -5.5}, point3)
}

func TestPoint_VectorTo(t *testing.T) {
	point1 := Point{-1, 2, -3}
	point2 := Point{5.5, 0, 1.23}

	vector1 := point1.VectorTo(point2)
	assert.Equal(t, Vector{6.5, -2, 4.23}, vector1)

	vector2 := point2.VectorTo(point1)
	assert.Equal(t, Vector{-6.5, 2, -4.23}, vector2)
}

func TestPoint_String(t *testing.T) {
	point := Point{1.234, -14.567, 0.088888}

	assert.Equal(t, "(1.23, -14.57, 0.09)", fmt.Sprintf("%v", point))
}
