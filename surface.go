package main

type Surface interface {
	Intersection(ray Ray) (float64, Vector)
	Albedo() Color
}
