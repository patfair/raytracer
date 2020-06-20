package main

import (
	"errors"
	"math"
	"strings"
)

type Camera struct {
	Rays [][]Ray
}

func NewCamera(viewDirection Ray, upDirection Vector, width, height int, horizontalFovDeg float64) (*Camera, error) {
	// Check for perpendicularity of view and up vectors.
	if viewDirection.Dot(upDirection) != 0 {
		return nil, errors.New("camera view and up direction vectors must be perpendicular")
	}

	halfWidth := float64(width) / 2
	halfHeight := float64(height) / 2
	pixelWidth := math.Tan(horizontalFovDeg*math.Pi/180/2) / halfWidth
	pixelHeight := pixelWidth * halfHeight / halfWidth

	uXyz := viewDirection.Cross(upDirection).ToUnit()
	vXyz := viewDirection.ToUnit()
	wXyz := upDirection.ToUnit()

	rays := make([][]Ray, height)
	for i := 0; i < height; i++ {
		rays[i] = make([]Ray, width)
		w := (float64(height-i-1) - halfHeight + 0.5) * pixelHeight
		for j := 0; j < width; j++ {
			u := (float64(j) - halfWidth + 0.5) * pixelWidth
			rays[i][j].Point = viewDirection.Point
			rays[i][j].Vector = uXyz.Multiply(u).Add(wXyz.Multiply(w)).Add(vXyz)
		}
	}

	return &Camera{Rays: rays}, nil
}

func (camera Camera) String() string {
	var rowStrings []string
	for _, row := range camera.Rays {
		var rayStrings []string
		for _, ray := range row {
			rayStrings = append(rayStrings, ray.String())
		}
		rowStrings = append(rowStrings, strings.Join(rayStrings, "\t"))
	}
	return strings.Join(rowStrings, "\n")
}
