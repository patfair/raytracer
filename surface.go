package main

import "github.com/patfair/raytracer/geometry"

type Surface interface {
	Intersection(ray geometry.Ray) *geometry.Intersection
	AlbedoAt(point geometry.Point) Color
	ShadingProperties() ShadingProperties
}
