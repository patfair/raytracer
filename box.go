package main

func NewBox(plane Plane, depth float64) [6]Plane {
	widthVector := plane.Width
	heightVector := plane.Height
	depthVector := plane.Normal().Multiply(depth)

	frontBottomLeft := plane.Corner
	backTopRight := frontBottomLeft.Translate(widthVector).Translate(heightVector).Translate(depthVector)

	front := plane
	bottom := Plane{frontBottomLeft, depthVector, widthVector, plane.Texture}
	left := Plane{frontBottomLeft, depthVector, heightVector, plane.Texture}
	back := Plane{backTopRight, widthVector.Multiply(-1), heightVector.Multiply(-1), plane.Texture}
	top := Plane{backTopRight, depthVector.Multiply(-1), widthVector.Multiply(-1), plane.Texture}
	right := Plane{backTopRight, depthVector.Multiply(-1), heightVector.Multiply(-1), plane.Texture}

	return [6]Plane{front, bottom, left, back, top, right}
}