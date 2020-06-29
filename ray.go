package main

import "fmt"

type Ray struct {
	Point     Point
	Direction Vector
}

func (ray Ray) String() string {
	return fmt.Sprintf("%v|%v", ray.Point, ray.Direction)
}
