package main

type Texture interface {
	AlbedoAt(u, v float64) Color
}
