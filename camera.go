package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"image"
	"math"
	"math/rand"
	"strings"
)

const (
	parallelism = 5
)

type Camera struct {
	Rays              [][][]Ray
	SupersampleFactor int
}

func NewCamera(viewCenter Ray, upDirection Vector, width, height int, horizontalFovDeg float64, apertureRadius float64,
	focalDistance float64, depthOfFieldSamples int, supersampleFactor int) (*Camera, error) {
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
	rays := make([][][]Ray, height)

	for i := 0; i < height; i++ {
		rays[i] = make([][]Ray, width)
		w := (float64(height-i-1) - halfHeight + 0.5) * pixelSize
		for j := 0; j < width; j++ {
			rays[i][j] = make([]Ray, depthOfFieldSamples)
			for n := 0; n < depthOfFieldSamples; n++ {
				u := (float64(j) - halfWidth + 0.5) * pixelSize
				nominalRayDirection := uXyz.Multiply(u).Add(wXyz.Multiply(w)).Add(vXyz).ToUnit()
				focalPlanePoint := viewCenter.Point.Translate(nominalRayDirection.Multiply(focalDistance))

				// Adjust the center ray to simulate a non-zero aperture, to produce a depth-of-field effect.
				r := apertureRadius * math.Sqrt(rand.Float64())
				phi := (float64(n) + rand.Float64()) * 2 * math.Pi / float64(depthOfFieldSamples)
				deltaU := r * math.Cos(phi)
				deltaW := r * math.Sin(phi)
				modifiedOrigin := viewCenter.Point.Translate(uXyz.Multiply(deltaU)).Translate(wXyz.Multiply(deltaW))

				rays[i][j][n].Point = modifiedOrigin
				rays[i][j][n].Direction = modifiedOrigin.VectorTo(focalPlanePoint).ToUnit()
			}
		}
	}

	return &Camera{Rays: rays, SupersampleFactor: supersampleFactor}, nil
}

func (camera *Camera) Render(scene *Scene) *image.RGBA {
	width := len(camera.Rays[0])
	height := len(camera.Rays)
	pixels := make([][]Color, height)
	for y, _ := range camera.Rays {
		pixels[y] = make([]Color, width)
	}

	// Set up progress bar for the console.
	progress := pb.Full.Start(width * height)

	// Set up parallel jobs to take advantage of multiple processor cores.
	doneChannel := make(chan struct{})
	rowsPerJob := (height + parallelism - 1) / parallelism // Round up
	start := 0
	for i := 0; i < parallelism; i++ {
		count := int(math.Min(float64(rowsPerJob), float64(height-start)))
		request := RaytraceRowsRequest{
			Scene:       scene,
			Rays:        camera.Rays,
			Start:       start,
			Count:       count,
			Pixels:      pixels,
			Progress:    progress,
			DoneChannel: doneChannel,
		}
		go request.Run()
		start += count
	}

	for i := 0; i < parallelism; i++ {
		<-doneChannel
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

	progress.Finish()
	return img
}

func (camera Camera) String() string {
	var rowStrings []string
	for _, row := range camera.Rays {
		var rayStrings []string
		for _, pixelRays := range row {
			for _, ray := range pixelRays {
				rayStrings = append(rayStrings, ray.String())
			}
		}
		rowStrings = append(rowStrings, strings.Join(rayStrings, "\t"))
	}
	return strings.Join(rowStrings, "\n")
}
