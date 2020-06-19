package main

import (
	"fmt"
	"testing"
)

func TestNewCamera(t *testing.T) {
	camera := NewCamera(2, 2, 90)
	fmt.Println(camera)
}
