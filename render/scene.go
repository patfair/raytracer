// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package render

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/patfair/raytracer/light"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
	"image"
	"math/rand"
	"runtime"
)

// Contains all the information required to render a particular view of a set.
type Scene struct {
	Camera          *Camera           // Virtual camera to specify the position and angle from which the scene is viewed
	BackgroundColor shading.Color     // Color to render for rays that do not intersect any surfaces
	Surfaces        []surface.Surface // Surfaces in the scene that rays can intercept
	Lights          []light.Light     // Virtual lights to illuminate surfaces in the scene and cast shadows
}

func (scene *Scene) AddSurface(surface surface.Surface) {
	scene.Surfaces = append(scene.Surfaces, surface)
}

func (scene *Scene) AddLight(light light.Light) {
	scene.Lights = append(scene.Lights, light)
}

// Executes the raytracing algorithm on the scene and returns the result as an image.
func (scene *Scene) Render(renderType RenderType, width, height int) (*image.RGBA, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("width and height must be positive numbers")
	}

	// Unless configured not to, produce a draft image at the same resolution without any multi-pass features, to be
	// used during the final rendering pass to avoid doing expensive multi-pass ray casting for pixels that don't need
	// it.
	var roughPassPixels [][]shading.Color
	if renderType == RenderFinishPass {
		fmt.Println("Producing draft image for mapping optimizations...")
		roughPassPixels = scene.renderPixels(RenderPreFinishRoughPass, width, height, [][]shading.Color{})
		fmt.Println()
	}

	fmt.Println("Producing final image...")
	pixels := scene.renderPixels(renderType, width, height, roughPassPixels)
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, pixels[y][x].ToRgba())
		}
	}

	return img, nil
}

// Executes the raytracing algorithm on the scene and returns the result as a two-dimensional array of pixels.
func (scene *Scene) renderPixels(renderType RenderType, width, height int,
	roughPassPixels [][]shading.Color) [][]shading.Color {
	// Set up progress bar for the console.
	progress := pb.Full.Start(width * height)

	outputPixels := make([][]shading.Color, height)
	for i := 0; i < height; i++ {
		outputPixels[i] = make([]shading.Color, width)
	}

	// Set up parallel operations to take advantage of multiple processor cores.
	numOperations := height
	operationsChannel := make(chan RaytraceRowOperation, numOperations)
	doneChannel := make(chan struct{}, numOperations)
	operations := make([]RaytraceRowOperation, numOperations)
	shufflePositions := rand.Perm(numOperations)

	// Create the pool of worker goroutines.
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		go func() {
			for operation := range operationsChannel {
				operation.Run()
			}
		}()
	}

	for i := 0; i < height; i++ {
		// Shuffle the operations to make progress more linear and predicted end time more accurate.
		operations[shufflePositions[i]] = RaytraceRowOperation{
			Scene:           scene,
			RenderType:      renderType,
			Width:           width,
			Height:          height,
			RowIndex:        i,
			RoughPassPixels: roughPassPixels,
			OutputPixels:    outputPixels,
			Progress:        progress,
			DoneChannel:     doneChannel,
		}
	}

	// Dispatch the operations to the queue.
	for _, operation := range operations {
		operationsChannel <- operation
	}
	close(operationsChannel)

	// Block until all operations are complete.
	for i := 0; i < numOperations; i++ {
		<-doneChannel
	}

	progress.Finish()
	return outputPixels
}
