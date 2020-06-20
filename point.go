package main

import "fmt"

type Point struct {
	X float64
	Y float64
	Z float64
}

func (point Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", point.X, point.Y, point.Z)
}
