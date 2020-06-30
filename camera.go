package main

import (
	"errors"
	"image"
	"math"
	"strings"
)

const shadowBias = 0.01

type Camera struct {
	Rays              [][]Ray
	SupersampleFactor int
}

func NewCamera(viewCenter Ray, upDirection Vector, width, height int, horizontalFovDeg float64,
	supersampleFactor int) (*Camera, error) {
	// Check for validity of dimensions.
	if width <= 0 || height <= 0 {
		return nil, errors.New("width and height must be positive numbers")
	}

	// Check for perpendicularity of view and up vectors.
	if viewCenter.Direction.Dot(upDirection) != 0 {
		return nil, errors.New("camera view and up direction vectors must be perpendicular")
	}

	// Scale the dimensions for anti-aliasing supersampling.
	width *= supersampleFactor
	height *= supersampleFactor

	halfWidth := float64(width) / 2
	halfHeight := float64(height) / 2
	pixelSize := math.Tan(horizontalFovDeg*math.Pi/180/2) / halfWidth

	uXyz := viewCenter.Direction.Cross(upDirection).ToUnit()
	vXyz := viewCenter.Direction.ToUnit()
	wXyz := upDirection.ToUnit()

	rays := make([][]Ray, height)
	for i := 0; i < height; i++ {
		rays[i] = make([]Ray, width)
		w := (float64(height-i-1) - halfHeight + 0.5) * pixelSize
		for j := 0; j < width; j++ {
			u := (float64(j) - halfWidth + 0.5) * pixelSize
			rays[i][j].Point = viewCenter.Point
			rays[i][j].Direction = uXyz.Multiply(u).Add(wXyz.Multiply(w)).Add(vXyz)
		}
	}

	return &Camera{Rays: rays, SupersampleFactor: supersampleFactor}, nil
}

func (camera *Camera) Render(surfaces []Surface, lights []Light, backgroundColor Color) *image.RGBA {
	width := len(camera.Rays[0])
	height := len(camera.Rays)

	pixels := make([][]Color, height)
	for y, row := range camera.Rays {
		pixels[y] = make([]Color, width)
		for x, ray := range row {
			var closestIntersection *Intersection
			var closestSurface Surface
			for _, surface := range surfaces {
				if intersection := surface.Intersection(ray); intersection != nil {
					if closestIntersection == nil || intersection.Distance < closestIntersection.Distance {
						closestIntersection = intersection
						closestSurface = surface
					}
				}
			}

			pixelColor := backgroundColor
			if closestIntersection != nil {
				var color Color
				for _, light := range lights {
					// Check if there is an object between the intersection point and the light source, in which case
					// it should cast a shadow.
					lightRay := Ray{
						Point:     closestIntersection.Point,
						Direction: light.Direction(closestIntersection.Point).Multiply(-1),
					}
					shadow := false
					for _, surface := range surfaces {
						if intersection := surface.Intersection(lightRay); intersection != nil {
							// Require a minimum distance to avoid a surface from shadowing itself.
							if intersection.Distance > shadowBias {
								if light.IsBlockedByIntersection(closestIntersection.Point, intersection) {
									shadow = true
									break
								}
							}
						}
					}
					if shadow {
						continue
					}

					incidentDotProduct :=
						light.Direction(closestIntersection.Point).Multiply(-1).Dot(closestIntersection.Normal)
					incidentLight := light.Intensity(closestIntersection.Point) * math.Max(incidentDotProduct, 0)
					color.R += closestSurface.AlbedoAt(closestIntersection.Point).R / math.Pi * light.Color().R *
						incidentLight
					color.G += closestSurface.AlbedoAt(closestIntersection.Point).G / math.Pi * light.Color().G *
						incidentLight
					color.B += closestSurface.AlbedoAt(closestIntersection.Point).B / math.Pi * light.Color().B *
						incidentLight
				}
				pixelColor = color
			}
			pixels[y][x] = pixelColor
		}
	}

	finalWidth := width / camera.SupersampleFactor
	finalHeight := height / camera.SupersampleFactor
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{finalWidth, finalHeight}})
	for x := 0; x < finalWidth; x++ {
		for y := 0; y < finalHeight; y++ {
			var averagePixel Color
			for i := 0; i < camera.SupersampleFactor; i++ {
				for j := 0; j < camera.SupersampleFactor; j++ {
					pixel := pixels[y*camera.SupersampleFactor+j][x*camera.SupersampleFactor+i]
					averagePixel.R += pixel.R
					averagePixel.G += pixel.G
					averagePixel.B += pixel.B
				}
			}
			numSamples := float64(camera.SupersampleFactor * camera.SupersampleFactor)
			averagePixel.R /= numSamples
			averagePixel.G /= numSamples
			averagePixel.B /= numSamples
			img.Set(x, y, averagePixel.ToRgba())
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
