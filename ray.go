package main

import "fmt"

type Ray struct {
	Point
	Vector
}

func (ray Ray) String() string {
	return fmt.Sprintf("%v|%v", ray.Point, ray.Vector)
}
