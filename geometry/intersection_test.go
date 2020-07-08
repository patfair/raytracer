// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection_String(t *testing.T) {
	intersection := Intersection{
		Point:    Point{1.234, -14.567, 0.088888},
		Distance: -123.456,
		Normal:   Vector{0.123, 4.567, 8.9},
	}

	assert.Equal(t, "P:(1.23, -14.57, 0.09)|D:-123.46|N:(0.12, 4.57, 8.90)", fmt.Sprintf("%v", intersection))
}
