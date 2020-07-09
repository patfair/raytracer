package main

import (
	"github.com/patfair/raytracer/light"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
)

type Scene struct {
	Surfaces        []surface.Surface
	Lights          []light.Light
	BackgroundColor shading.Color
}
