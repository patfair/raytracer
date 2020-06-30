package main

type SolidTexture struct {
	color Color
}

func (texture SolidTexture) AlbedoAt(u, v float64) Color {
	return texture.color
}
