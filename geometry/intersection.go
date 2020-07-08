// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package geometry

import "fmt"

// Contains information about the intersection between a Ray and a Surface.
type Intersection struct {
	Point    Point   // Point at which the Ray intersects the Surface
	Distance float64 // Distance between the Ray origin and the intersection point
	Normal   Vector  // Vector that is normal to the Surface at the intersection point
}

func (intersection Intersection) String() string {
	return fmt.Sprintf("P:%v|D:%.2f|N:%v", intersection.Point, intersection.Distance, intersection.Normal)
}
