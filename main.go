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
		Sphere{Point{1, 1, 1}, 0.25, Color{0.75, 0.5, 0}},
		Sphere{Point{1, 5, 1}, 0.5, Color{1, 1, 1}},
	}
	lights := []Light{
		DistantLight{Vector{-10, -10, -20}, Color{1, 1, 1}, 0.75},
		DistantLight{Vector{-10, -10, -25}, Color{1, 1, 1}, 0.75},
		DistantLight{Vector{-11, -9, -20}, Color{1, 1, 1}, 0.75},
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
