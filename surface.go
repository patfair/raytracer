package main

import (
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
)

type Surface interface {
	Intersection(ray geometry.Ray) *geometry.Intersection
	AlbedoAt(point geometry.Point) shading.Color
	ShadingProperties() shading.ShadingProperties
}
