// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package surface

import (
	"errors"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
)

// Returns the set of six planes formed by extruding the plane defined by the given parameters by the given depth.
func NewBox(frontBottomLeftCorner geometry.Point, width, height geometry.Vector, depth float64,
	shadingProperties shading.ShadingProperties) ([6]Plane, error) {
	if depth == 0 {
		return [6]Plane{}, errors.New("depth must be non-zero")
	}
	front, err := NewPlane(frontBottomLeftCorner, width, height, shadingProperties)
	if err != nil {
		return [6]Plane{}, err
	}

	depthVector := front.normal.Multiply(depth)
	backTopRightCorner := frontBottomLeftCorner.Translate(width).Translate(height).Translate(depthVector)

	bottom, _ := NewPlane(frontBottomLeftCorner, depthVector, width, shadingProperties)
	left, _ := NewPlane(frontBottomLeftCorner, depthVector, height, shadingProperties)
	back, _ := NewPlane(backTopRightCorner, width.Multiply(-1), height.Multiply(-1), shadingProperties)
	top, _ := NewPlane(backTopRightCorner, depthVector.Multiply(-1), width.Multiply(-1), shadingProperties)
	right, _ := NewPlane(backTopRightCorner, depthVector.Multiply(-1), height.Multiply(-1), shadingProperties)

	return [6]Plane{front, bottom, left, back, top, right}, nil
}
