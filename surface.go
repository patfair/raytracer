package main

type Surface interface {
	Intersection(ray Ray) *Intersection
	Albedo() Color
}
