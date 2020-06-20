package main

type Surface interface {
	Intersection(ray Ray) float64
}
