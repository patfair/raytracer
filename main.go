// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package main

import (
	"fmt"
	"github.com/patfair/raytracer/example"
	"github.com/patfair/raytracer/render"
	"image/png"
	"os"
)

func main() {
	scene, err := example.SpheresScene()
	handleError(err)

	image, err := scene.Render(render.RenderFinishPass, 3840, 2160)
	handleError(err)

	file, err := os.Create("image.png")
	handleError(err)

	err = png.Encode(file, image)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
