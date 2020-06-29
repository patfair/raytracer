package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	surfaces := []Surface{
		Plane{Point{0, 0, 0}, Vector{1, 0, 0}, Color{1, 0, 0}},     // YZ plane
		Plane{Point{0, 0, 0}, Vector{0, 1, 0}, Color{1, 1, 0}},     // XZ plane
		Plane{Point{0, 0, 0}, Vector{0, 0, 1}, Color{0, 0.5, 0.5}}, // XY plane
	}
	lights := []Light{
		DistantLight{Vector{-10, -10, -20}, Color{1, 1, 1}, 2},
	}

	camera, err := NewCamera(Ray{Point{10, 10, 5}, Vector{-10, -10, -5}}, Vector{-10, -10, 40}, 1600, 900, 40)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := camera.Render(surfaces, lights)

	file, _ := os.Create("image.png")
	png.Encode(file, img)
}
