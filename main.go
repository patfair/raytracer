package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	//surfaces := []Surface{
	//	Plane{Point{0, 0, 0}, Vector{1, 0, 0}}, // YZ plane
	//	Plane{Point{0, 0, 0}, Vector{0, 1, 0}}, // XZ plane
	//	Plane{Point{0, 0, 0}, Vector{0, 0, 1}}, // XY plane
	//}

	width := 1920
	height := 1080
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			red := x * 255 / width
			green := (height - y - 1) * 255 / height
			img.Set(x, y, color.RGBA{uint8(red), uint8(green), 0, 255})
		}
	}

	file, _ := os.Create("image.png")
	png.Encode(file, img)
}
