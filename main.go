package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	//scene := AllElementsScene()
	scene := SpheresScene()

	//camera, err := NewCamera(Ray{Point{10, 10, 5}, Vector{-10, -10, -5}}, Vector{-10, -10, 40}, 3840, 2160, 30, 2)
	//camera, err := NewCamera(Ray{Point{0, 0, 3}, Vector{0, 1, -0.2}}, Vector{0, 0.2, 1}, 1280, 720, 40, 1)
	camera, err := NewCamera(Ray{Point{0, 0, 3}, Vector{0, 1, -0.2}}, Vector{0, 0.2, 1}, 3840, 2160, 40, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := camera.Render(&scene)

	file, _ := os.Create("image.png")
	png.Encode(file, img)
}
