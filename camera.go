package main

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"image"
	"math"
	"math/rand"
	"runtime"
)

type Camera struct {
	Point               geometry.Point
	Width               int
	Height              int
	PixelSize           float64
	UVector             geometry.Vector
	VVector             geometry.Vector
	WVector             geometry.Vector
	ApertureRadius      float64
	FocalDistance       float64
	DepthOfFieldSamples int
	SupersampleFactor   int
	isDraft             bool
}

func NewCamera(viewCenter geometry.Ray, upDirection geometry.Vector, width, height int, horizontalFovDeg float64, apertureRadius float64,
	focalDistance float64, depthOfFieldSamples int, supersampleFactor int) (*Camera, error) {
	// Check for validity of dimensions.
	if width <= 0 || height <= 0 {
		return nil, errors.New("width and height must be positive numbers")
	}

	// Check for perpendicularity of view and up vectors.
	if viewCenter.Direction.Dot(upDirection) != 0 {
		return nil, errors.New("camera view and up direction vectors must be perpendicular")
	}

	pixelSize := 2 * math.Tan(horizontalFovDeg*math.Pi/180/2) / float64(width)

	uXyz := viewCenter.Direction.Cross(upDirection).ToUnit()
	vXyz := viewCenter.Direction.ToUnit()
	wXyz := upDirection.ToUnit()

	return &Camera{
		Point:               viewCenter.Origin,
		Width:               width,
		Height:              height,
		PixelSize:           pixelSize,
		UVector:             uXyz,
		VVector:             vXyz,
		WVector:             wXyz,
		ApertureRadius:      apertureRadius,
		FocalDistance:       focalDistance,
		DepthOfFieldSamples: depthOfFieldSamples,
		SupersampleFactor:   supersampleFactor,
		isDraft:             false,
	}, nil
}

func (camera *Camera) GetRay(x, y, depthOfFieldSampleIndex, supersampleFactor, supersampleIndexX,
	supersampleIndexY int) geometry.Ray {
	w := (float64(camera.Height*supersampleFactor)/2 - float64(y*supersampleFactor+supersampleIndexY+1) + 0.5) *
		camera.PixelSize / float64(supersampleFactor)
	u := (float64(x*supersampleFactor+supersampleIndexX) - float64(camera.Width*supersampleFactor)/2 + 0.5) *
		camera.PixelSize / float64(supersampleFactor)
	nominalRayDirection :=
		camera.UVector.Multiply(u).Add(camera.WVector.Multiply(w)).Add(camera.VVector).ToUnit()
	focalPlanePoint := camera.Point.Translate(nominalRayDirection.Multiply(camera.FocalDistance))

	// Adjust the center ray to simulate a non-zero aperture, to produce a depth-of-field effect.
	r := camera.ApertureRadius * math.Sqrt(rand.Float64())
	phi := (float64(depthOfFieldSampleIndex) + rand.Float64()) * 2 * math.Pi / float64(camera.DepthOfFieldSamples)
	deltaU := r * math.Cos(phi)
	deltaW := r * math.Sin(phi)
	modifiedOrigin :=
		camera.Point.Translate(camera.UVector.Multiply(deltaU)).Translate(camera.WVector.Multiply(deltaW))

	return geometry.Ray{Origin: modifiedOrigin, Direction: modifiedOrigin.VectorTo(focalPlanePoint).ToUnit()}
}

func (camera *Camera) Render(scene *Scene) *image.RGBA {
	fmt.Println("Producing draft image for mapping optimizations...")
	draftCamera := *camera
	draftCamera.ApertureRadius = 0
	draftCamera.DepthOfFieldSamples = 1
	draftCamera.SupersampleFactor = 1
	draftPixels := draftCamera.renderPixels(scene, true, [][]shading.Color{})

	fmt.Println("\nProducing final image...")
	pixels := camera.renderPixels(scene, false, draftPixels)
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{camera.Width, camera.Height}})
	for x := 0; x < camera.Width; x++ {
		for y := 0; y < camera.Height; y++ {
			img.Set(x, y, pixels[y][x].ToRgba())
		}
	}

	return img
}

func (camera *Camera) renderPixels(scene *Scene, isDraft bool, draftPixels [][]shading.Color) [][]shading.Color {
	// Set up progress bar for the console.
	progress := pb.Full.Start(camera.Width * camera.Height)

	pixels := make([][]shading.Color, camera.Height)
	for i := 0; i < camera.Height; i++ {
		pixels[i] = make([]shading.Color, camera.Width)
	}

	// Set up parallel jobs to take advantage of multiple processor cores.
	numJobs := camera.Height
	jobsChannel := make(chan RaytraceRowRequest, numJobs)
	doneChannel := make(chan struct{}, numJobs)
	requests := make([]RaytraceRowRequest, numJobs)
	shufflePositions := rand.Perm(numJobs)

	// Create the workers.
	numThreads := runtime.NumCPU()
	for i := 0; i < numThreads; i++ {
		go func() {
			for request := range jobsChannel {
				request.Run()
			}
		}()
	}

	for i := 0; i < camera.Height; i++ {
		// Shuffle the requests to make progress more linear and predicted end time more accurate.
		requests[shufflePositions[i]] = RaytraceRowRequest{
			Scene:       scene,
			Camera:      camera,
			RowIndex:    i,
			IsDraft:     isDraft,
			DraftPixels: draftPixels,
			Pixels:      pixels,
			Progress:    progress,
			DoneChannel: doneChannel,
		}
	}

	for _, request := range requests {
		jobsChannel <- request
	}
	close(jobsChannel)

	for i := 0; i < numJobs; i++ {
		<-doneChannel
	}

	progress.Finish()
	return pixels
}
