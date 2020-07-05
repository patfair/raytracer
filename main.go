package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	//scene, camera, err := AllElementsScene()
	scene, camera, err := SpheresScene()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := camera.Render(scene)

	file, _ := os.Create("image.png")
	png.Encode(file, img)
}
