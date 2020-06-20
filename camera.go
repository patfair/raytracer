package main

import (
	"errors"
	"image"
	"image/color"
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

func (camera *Camera) Render(surfaces []Surface) *image.RGBA {
	width := len(camera.Rays[0])
	height := len(camera.Rays)
	minDistance := math.MaxFloat64
	maxDistance := float64(0)
	distances := make([][]float64, len(camera.Rays))

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for y, row := range camera.Rays {
		distances[y] = make([]float64, len(row))
		for x, ray := range row {
			closestDistance := float64(-1)
			for _, surface := range surfaces {
				distance := surface.Intersection(ray)
				if distance > 0 {
					if closestDistance < 0 || distance < closestDistance {
						closestDistance = distance
					}
				}
			}
			distances[y][x] = closestDistance
			if closestDistance > 0 {
				if closestDistance < minDistance {
					minDistance = closestDistance
				}
				if closestDistance > maxDistance {
					maxDistance = closestDistance
				}
			}
		}
	}

	for y, row := range distances {
		for x, distance := range row {
			red := 0
			if distance > 0 {
				red = int(255 * math.Pow((maxDistance - distance) / (maxDistance - minDistance), 2))
			}

			img.Set(x, y, color.RGBA{uint8(red), 0, 0, 255})
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
