// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import (
	"fmt"
)

// Represents half of a line, starting from an origin point and proceeding in a single direction.
type Ray struct {
	Origin    Point  // Point from which the ray originates
	Direction Vector // Direction in which the ray points
}

func (ray Ray) String() string {
	return fmt.Sprintf("O:%v|D:%v", ray.Origin, ray.Direction)
}
