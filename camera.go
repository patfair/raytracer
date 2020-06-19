package main

import "math"

type Camera struct {
	Rays [][]Ray
}

// Assumes camera is located at (0, 0, 0) and pointing along negative Z.
func NewCamera(width, height int, horizontalFovDeg float64) *Camera {
	position := Point{0, 0, 0}

	halfWidth := float64(width) / 2
	halfHeight := float64(height) / 2
	pixelWidth := math.Tan(horizontalFovDeg*math.Pi/180/2) / halfWidth
	pixelHeight := pixelWidth * halfHeight / halfWidth

	rays := make([][]Ray, width)
	for i := 0; i < width; i++ {
		rays[i] = make([]Ray, height)
		x := (float64(i) - halfWidth + 0.5) * pixelWidth
		for j := 0; j < height; j++ {
			y := (float64(j) - halfHeight + 0.5) * pixelHeight
			rays[i][j] = Ray{Point: position, Vector: Vector{x, y, -1}}
		}
	}

	return &Camera{Rays: rays}
}
