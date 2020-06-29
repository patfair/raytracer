package main

import (
	"errors"
	"image"
	"math"
	"strings"
)

type Camera struct {
	Rays [][]Ray
}

func NewCamera(viewDirection Ray, upDirection Vector, width, height int, horizontalFovDeg float64) (*Camera, error) {
	// Check for validity of dimensions.
	if width <= 0 || height <= 0 {
		return nil, errors.New("width and height must be positive numbers")
	}

	// Check for perpendicularity of view and up vectors.
	if viewDirection.Dot(upDirection) != 0 {
		return nil, errors.New("camera view and up direction vectors must be perpendicular")
	}

	halfWidth := float64(width) / 2
	halfHeight := float64(height) / 2
	pixelSize := math.Tan(horizontalFovDeg*math.Pi/180/2) / halfWidth

	uXyz := viewDirection.Cross(upDirection).ToUnit()
	vXyz := viewDirection.ToUnit()
	wXyz := upDirection.ToUnit()

	rays := make([][]Ray, height)
	for i := 0; i < height; i++ {
		rays[i] = make([]Ray, width)
		w := (float64(height-i-1) - halfHeight + 0.5) * pixelSize
		for j := 0; j < width; j++ {
			u := (float64(j) - halfWidth + 0.5) * pixelSize
			rays[i][j].Point = viewDirection.Point
			rays[i][j].Vector = uXyz.Multiply(u).Add(wXyz.Multiply(w)).Add(vXyz)
		}
	}

	return &Camera{Rays: rays}, nil
}

func (camera *Camera) Render(surfaces []Surface, lights []Light) *image.RGBA {
	width := len(camera.Rays[0])
	height := len(camera.Rays)

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for y, row := range camera.Rays {
		for x, ray := range row {
			var closestSurface Surface
			closestDistance := float64(-1)
			var closestNormal Vector
			for _, surface := range surfaces {
				distance, normal := surface.Intersection(ray)
				if distance > 0 {
					if closestDistance < 0 || distance < closestDistance {
						closestSurface = surface
						closestDistance = distance
						closestNormal = normal
					}
				}
			}

			if closestDistance > 0 {
				var color Color
				for _, light := range lights {
					incidentLight := light.Intensity() * math.Max(light.Direction().Multiply(-1).Dot(closestNormal), 0)
					color.R += closestSurface.Albedo().R / math.Pi * light.Color().R * incidentLight
					color.G += closestSurface.Albedo().G / math.Pi * light.Color().G * incidentLight
					color.B += closestSurface.Albedo().B / math.Pi * light.Color().B * incidentLight
				}

				img.Set(x, y, color.ToRgba())

			}
		}
	}

	return img
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
