package main

type Surface interface {
	Intersection(ray Ray) *Intersection
	AlbedoAt(point Point) Color
	ShadingProperties() ShadingProperties
}
