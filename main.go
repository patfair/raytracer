package main

import (
	"fmt"
	"image/png"
	"math"
	"os"
)

func main() {
	surfaces := []Surface{
		Plane{Corner: Point{0, 0, 0}, Width: Vector{0, 4, 0}, Height: Vector{0, 0, 2},
			Texture: CheckerboardTexture{Color{0.9, 0.1, 0.1}, Color{0.8, 0.8, 0.8}, 1, 0.5}}, // YZ plane
		Plane{Corner: Point{0, 0, 0}, Width: Vector{4, 0, 0}, Height: Vector{0, 0, 2},
			Texture: CheckerboardTexture{Color{0.2, 0.5, 1}, Color{0, 0, 0}, 0.1, 0.1}}, // XZ plane
		Plane{Corner: Point{0, 0, 0}, Width: Vector{4, 0, 0}, Height: Vector{0, 10, 0},
			Texture: CheckerboardTexture{Color{0.9, 0.9, 0.9}, Color{0.2, 0.2, 0.2}, 0.3, 0.3}}, // XY plane
		Sphere{Center: Point{1, 1, 1}, Radius: 0.25, ZenithReference: Vector{0, 0, 1},
			AzimuthReference: Vector{1, 0, 0},
			Texture:          CheckerboardTexture{Color{1, 0, 1}, Color{1, 1, 1}, math.Pi / 4, math.Pi / 8}},
		Sphere{Center: Point{1, 5, 1}, Radius: 0.5, ZenithReference: Vector{0, 1, 0},
			AzimuthReference: Vector{1, 0, 0},
			Texture:          CheckerboardTexture{Color{1, 1, 1}, Color{0, 0, 1}, math.Pi / 2, math.Pi / 4}},
		Disc{Plane{Corner: Point{3, 1, 0.5}, Width: Vector{0.5, 0, 0}, Height: Vector{0, 1, 0},
			Texture: CheckerboardTexture{Color{1, 0, 0}, Color{0, 0, 1}, 0.25, 2 * math.Pi}}},
	}
	boxFront := Plane{Corner: Point{0.5, 3, 1}, Width: Vector{0.2, 0.2, 0.2}, Height: Vector{-0.1, -0.1, 0.2},
		Texture: SolidTexture{Color{0, 1, 0.2}}}
	for _, plane := range NewBox(boxFront, 0.2) {
		surfaces = append(surfaces, plane)
	}

	lights := []Light{
		DistantLight{Vector{-10, -10, -20}, Color{1, 1, 1}, 0.75},
		DistantLight{Vector{-10, -10, -25}, Color{1, 1, 1}, 0.75},
		DistantLight{Vector{-11, -9, -20}, Color{1, 1, 1}, 0.75},
		PointLight{Point{5, 1, 10}, Color{1, 1, 1}, 1000},
	}

	camera, err := NewCamera(Ray{Point{10, 10, 5}, Vector{-10, -10, -5}}, Vector{-10, -10, 40}, 3840, 2160, 40, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := camera.Render(surfaces, lights, Color{0.1, 0.8, 1})

	file, _ := os.Create("image.png")
	png.Encode(file, img)
}
