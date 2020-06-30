package main

import "math"

type CheckerboardTexture struct {
	Color1 Color
	Color2 Color
	UPitch float64
	VPitch float64
}

func (texture CheckerboardTexture) AlbedoAt(u, v float64) Color {
	if getToggleValue(u, texture.UPitch) == getToggleValue(v, texture.VPitch) {
		return texture.Color1
	}
	return texture.Color2
}

func getToggleValue(x, pitch float64) bool {
	_, fraction := math.Modf(x / pitch)
	if fraction < 0 {
		fraction += 1
	}
	return int(fraction*2) == 0
}
