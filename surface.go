package main

type Surface interface {
	Intersection(ray Ray) *Intersection
	AlbedoAt(point Point) Color
	Opacity() float64
	Reflectivity() float64
	RefractiveIndex() float64
}
