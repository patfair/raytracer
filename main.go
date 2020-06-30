package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	surfaces := []Surface{
		Plane{Corner: Point{0, 0, 0}, Width: Vector{0, 4, 0}, Height: Vector{0, 0, 2}, Color: Color{1, 0, 0}}, // YZ
		Plane{Corner: Point{0, 0, 0}, Width: Vector{4, 0, 0}, Height: Vector{0, 0, 2}, Color: Color{1, 1, 0}}, // XZ
		Plane{Corner: Point{0, 0, 0}, Width: Vector{4, 0, 0}, Height: Vector{0, 4, 0}, Color: Color{0, 1, 1}}, // XY
		Sphere{Point{1, 1, 1}, 0.25, Color{0.75, 0.5, 0}},
		Sphere{Point{1, 5, 1}, 0.5, Color{1, 1, 1}},
	}
	lights := []Light{
		DistantLight{Vector{-10, -10, -20}, Color{1, 1, 1}, 0.75},
		DistantLight{Vector{-10, -10, -25}, Color{1, 1, 1}, 0.75},
		DistantLight{Vector{-11, -9, -20}, Color{1, 1, 1}, 0.75},
		PointLight{Point{5, 1, 10}, Color{1, 1, 1}, 5000},
	}

	camera, err := NewCamera(Ray{Point{10, 10, 5}, Vector{-10, -10, -5}}, Vector{-10, -10, 40}, 1600, 900, 40)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := camera.Render(surfaces, lights, Color{0.1, 0.8, 1})

	file, _ := os.Create("image.png")
	png.Encode(file, img)
}
