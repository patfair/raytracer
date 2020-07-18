// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/patfair/raytracer/example"
	"github.com/patfair/raytracer/render"
	"image/png"
	"os"
	"strings"
)

func main() {
	width := flag.Int("width", 1920, "rendered image width in pixels")
	height := flag.Int("height", 1080, "rendered image height in pixels")
	draft := flag.Bool("draft", false, "whether to only render a rough draft without any multi-pass features enabled")
	outputFilename := flag.String("output", "", "PNG file path to write the rendered image to")
	frame := flag.Int("frame", 0, "frame number passed to the scene generation method for optional animation")
	flag.Parse()

	renderType := render.RenderFinishPass
	if *draft {
		renderType = render.RenderDraftPass
	}
	if *outputFilename == "" {
		handleError(errors.New("must specify output path"))
	}
	if !strings.HasSuffix(*outputFilename, ".png") {
		handleError(errors.New("output path must end in .png"))
	}

	scene, err := example.SpheresScene(*frame)
	handleError(err)

	image, err := scene.Render(renderType, *width, *height)
	handleError(err)

	file, err := os.Create(*outputFilename)
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
