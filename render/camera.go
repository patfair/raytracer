// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package render

import (
	"errors"
	"github.com/patfair/raytracer/geometry"
	"math"
	"math/rand"
)

type Camera struct {
	Point               geometry.Point
	HorizontalFovDeg    float64
	UVector             geometry.Vector
	VVector             geometry.Vector
	WVector             geometry.Vector
	ApertureRadius      float64
	FocalDistance       float64
	DepthOfFieldSamples int
	AntiAliasSamples    int
}

func NewCamera(viewCenter geometry.Ray, upDirection geometry.Vector, horizontalFovDeg float64, apertureRadius float64,
	focalDistance float64, depthOfFieldSamples int, antiAliasSamples int) (*Camera, error) {
	// Check for perpendicularity of view and up vectors.
	if viewCenter.Direction.Dot(upDirection) != 0 {
		return nil, errors.New("camera view and up direction vectors must be perpendicular")
	}
	if horizontalFovDeg <= 0 {
		return nil, errors.New("field of view must be positive")
	}
	if apertureRadius < 0 {
		return nil, errors.New("aperture radius must be non-negative")
	}
	if focalDistance <= 0 {
		return nil, errors.New("focal distance must be positive")
	}
	if depthOfFieldSamples <= 0 {
		return nil, errors.New("depth of field samples must be at least 1")
	}
	if antiAliasSamples <= 0 {
		return nil, errors.New("antialias samples must be at least 1")
	}

	uXyz := viewCenter.Direction.Cross(upDirection).ToUnit()
	vXyz := viewCenter.Direction.ToUnit()
	wXyz := upDirection.ToUnit()

	return &Camera{
		Point:               viewCenter.Origin,
		HorizontalFovDeg:    horizontalFovDeg,
		UVector:             uXyz,
		VVector:             vXyz,
		WVector:             wXyz,
		ApertureRadius:      apertureRadius,
		FocalDistance:       focalDistance,
		DepthOfFieldSamples: depthOfFieldSamples,
		AntiAliasSamples:    antiAliasSamples,
	}, nil
}

func (camera *Camera) GetRay(width, height, x, y, depthOfFieldSampleIndex, depthOfFieldSamples, antiAliasIndexX,
	antiAliasIndexY, antiAliasSamples int) geometry.Ray {
	pixelSize := 2 * math.Tan(camera.HorizontalFovDeg*math.Pi/180/2) / float64(width)
	w := (float64(height*antiAliasSamples)/2 - float64(y*antiAliasSamples+antiAliasIndexY+1) + 0.5) *
		pixelSize / float64(antiAliasSamples)
	u := (float64(x*antiAliasSamples+antiAliasIndexX) - float64(width*antiAliasSamples)/2 + 0.5) *
		pixelSize / float64(antiAliasSamples)
	nominalRayDirection :=
		camera.UVector.Multiply(u).Add(camera.WVector.Multiply(w)).Add(camera.VVector).ToUnit()
	focalPlanePoint := camera.Point.Translate(nominalRayDirection.Multiply(camera.FocalDistance))

	// Adjust the center ray to simulate a non-zero aperture, to produce a depth of field effect.
	apertureRadius := camera.ApertureRadius
	if depthOfFieldSamples == 1 {
		apertureRadius = 0
	}
	r := apertureRadius * math.Sqrt(rand.Float64())
	phi := (float64(depthOfFieldSampleIndex) + rand.Float64()) * 2 * math.Pi / float64(depthOfFieldSamples)
	deltaU := r * math.Cos(phi)
	deltaW := r * math.Sin(phi)
	modifiedOrigin :=
		camera.Point.Translate(camera.UVector.Multiply(deltaU)).Translate(camera.WVector.Multiply(deltaW))

	return geometry.Ray{Origin: modifiedOrigin, Direction: modifiedOrigin.VectorTo(focalPlanePoint).ToUnit()}
}
