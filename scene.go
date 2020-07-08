package main

import "github.com/patfair/raytracer/shading"

type Scene struct {
	Surfaces        []Surface
	Lights          []Light
	BackgroundColor shading.Color
}
