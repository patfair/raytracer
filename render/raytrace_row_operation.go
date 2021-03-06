// Copyright 2020 Patrick Fairbank. All Rights Reserved.

package render

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/patfair/raytracer/geometry"
	"github.com/patfair/raytracer/shading"
	"github.com/patfair/raytracer/surface"
	"math"
)

const (
	maxReflectionDepth = 20
	reflectionBias     = 0.001
	shadowBias         = 0.001
)

type RenderType int

const (
	RenderDraftPass RenderType = iota
	RenderFinishPass
)

// Represents an operation to render a single row of pixels within an image.
type RaytraceRowOperation struct {
	Scene           *Scene            // Scene to render
	RenderType      RenderType        // Whether this is a rough or finishing pass
	Width           int               // Width of the full image
	Height          int               // Height of the full image
	RowIndex        int               // Which single row along the height this operation is for
	RoughPassPixels [][]shading.Color // Output of the previous rough pass if this is the finish pass
	OutputPixels    [][]shading.Color // Array of pixels for the full image to write the rendered row to
	Progress        *pb.ProgressBar   // Progress indicator to update after rendering each pixel
	DoneChannel     chan struct{}     // Channel to send an empty message to to signal completion of the operation
}

// Executes the rendering operation synchronously.
func (operation *RaytraceRowOperation) Run() {
	camera := operation.Scene.Camera
	numDirectionalSamples := 1
	numTotalSamples := 1
	if operation.RenderType == RenderFinishPass {
		depthOfFieldSamples := float64(camera.DepthOfFieldSamples)
		antiAliasSamples := float64(camera.AntiAliasSamples * camera.AntiAliasSamples)
		shadowSamples := float64(operation.Scene.ShadowSamples)
		maxSamples := math.Max(math.Max(depthOfFieldSamples, antiAliasSamples), shadowSamples)

		// Round down to a perfect square, to be compatible with anti-aliasing.
		numDirectionalSamples = int(math.Sqrt(maxSamples))
		numTotalSamples = numDirectionalSamples * numDirectionalSamples
	}

	for j := 0; j < operation.Width; j++ {
		// Supersample and average together multiple rays for each pixel for depth of field and antialiasing.
		var averagePixel shading.Color
		n := 0
		for a := 0; a < numDirectionalSamples; a++ {
			for b := 0; b < numDirectionalSamples; b++ {
				ray := camera.GetRay(operation.Width, operation.Height, j, operation.RowIndex, n, numTotalSamples, a, b,
					numDirectionalSamples)
				n++
				pixel := operation.castRay(operation.Scene, ray, 0, 1, n, numTotalSamples)
				averagePixel.R += pixel.R
				averagePixel.G += pixel.G
				averagePixel.B += pixel.B
			}
		}
		averagePixel.R /= float64(numTotalSamples)
		averagePixel.G /= float64(numTotalSamples)
		averagePixel.B /= float64(numTotalSamples)

		operation.OutputPixels[operation.RowIndex][j] = averagePixel
		operation.Progress.Increment()
	}

	// Signal to the worker coordinator that this row is done being rendered.
	operation.DoneChannel <- struct{}{}
}

// Returns the color that the given ray is pointing at. Contains the main logic of the raytracer.
func (operation *RaytraceRowOperation) castRay(scene *Scene, ray geometry.Ray, depth int, refractionIndex float64,
	sampleIndex int, numSamples int) shading.Color {
	pixelColor := scene.BackgroundColor

	// Limit recursion caused by reflecting rays off multiple surfaces.
	if depth == maxReflectionDepth {
		return pixelColor
	}

	// Find the closest surface in the scene that the ray intersects, if any.
	var closestIntersection *geometry.Intersection
	var closestSurface surface.Surface
	for _, surface := range scene.Surfaces {
		if intersection := surface.Intersection(ray); intersection != nil {
			if closestIntersection == nil || intersection.Distance < closestIntersection.Distance {
				closestIntersection = intersection
				closestSurface = surface
			}
		}
	}

	if closestIntersection != nil {
		shadingProperties := closestSurface.ShadingProperties()
		kRefraction := 1 - shadingProperties.Opacity
		kReflection := shadingProperties.Reflectivity * shadingProperties.Opacity
		kDiffuse := 1 - kRefraction - kReflection
		kSpecular := shadingProperties.SpecularIntensity
		var refractedColor, reflectedColor, diffuseColor, specularColor shading.Color

		// Determine the component of the ray from light passing through a transparent surface.
		if kRefraction > 0 {
			cosIn := -closestIntersection.Normal.Dot(ray.Direction)
			etaIn := refractionIndex
			etaOut := shadingProperties.RefractiveIndex
			if refractionIndex > 1 {
				// If the previous refraction index isn't 1, the ray is exiting the material instead of entering.
				etaIn, etaOut = etaOut, etaIn
			}

			// Determine reflection and refraction components of the refracted light.
			sinOut := etaIn / etaOut * math.Sqrt(math.Max(1-cosIn*cosIn, 0))
			if sinOut < 1 {
				cosOut := math.Abs(math.Sqrt(math.Max(1-sinOut*sinOut, 0)))
				rParallel := ((etaOut * cosIn) - (etaIn * cosOut)) / ((etaOut * cosIn) + (etaIn * cosOut))
				rPerpendicular := ((etaIn * cosIn) - (etaOut * cosOut)) / ((etaIn * cosIn) + (etaOut * cosOut))
				reflectedComponent := (rParallel*rParallel + rPerpendicular*rPerpendicular) / 2

				// Adjust the coefficients for the reflected light coming from refraction.
				delta := reflectedComponent * kRefraction
				kRefraction -= delta
				kReflection += delta
			}

			eta := etaIn / etaOut
			k := 1 - eta*eta*(1-cosIn*cosIn)
			refractionDirection :=
				ray.Direction.Multiply(eta).Add(closestIntersection.Normal.Multiply(eta*cosIn - math.Sqrt(k)))

			// Bias the intersection point off the surface slightly to avoid immediate self-intersection.
			refractionPoint :=
				closestIntersection.Point.Translate(closestIntersection.Normal.Multiply(-reflectionBias))

			refractedRay := geometry.Ray{refractionPoint, refractionDirection.ToUnit()}
			refractedColor = operation.castRay(scene, refractedRay, depth+1, shadingProperties.RefractiveIndex,
				sampleIndex, numSamples)
		}

		// Determine the component of the ray from light reflected off a mirrored surface.
		reflectedDirection := ray.Direction.Add(
			closestIntersection.Normal.Multiply(-2 * closestIntersection.Normal.Dot(ray.Direction))).ToUnit()
		if kReflection > 0 {
			// Bias the intersection point off the surface slightly to avoid immediate self-intersection.
			reflectedPoint := closestIntersection.Point.Translate(closestIntersection.Normal.Multiply(reflectionBias))

			reflectedRay := geometry.Ray{reflectedPoint, reflectedDirection.ToUnit()}
			reflectedColor = operation.castRay(scene, reflectedRay, depth+1, refractionIndex, sampleIndex, numSamples)
		}

		// Determine the component of the ray from the scene's lights directly illuminating the surface.
		if kDiffuse > 0 || kSpecular > 0 {
			for _, light := range scene.Lights {
				lightDirection := light.Direction(closestIntersection.Point, sampleIndex, numSamples)

				// Check if there is an object between the intersection point and the light source, in which case
				// it should cast a shadow.
				lightRay := geometry.Ray{
					Origin:    closestIntersection.Point,
					Direction: lightDirection.Multiply(-1).ToUnit(),
				}
				transparency := 1.0
				for _, surface := range scene.Surfaces {
					if intersection := surface.Intersection(lightRay); intersection != nil {
						// Require a minimum distance to prevent floating-point imprecision causing a surface to
						// cast a shadow on itself.
						if intersection.Distance > shadowBias {
							if light.IsBlockedByIntersection(closestIntersection.Point, intersection) {
								transparency *= 1 - surface.ShadingProperties().Opacity
							}
						}
					}
				}
				if transparency == 0 {
					// The light is not reaching the intersection point at all; skip calculating its component color
					// from this light source since it will just be black.
					continue
				}

				// Calculate the diffuse component, influenced by the color of the surface itself.
				incidentDotProduct := lightDirection.Multiply(-1).Dot(closestIntersection.Normal)
				incidentLight := light.Intensity(closestIntersection.Point) * math.Max(incidentDotProduct, 0) *
					transparency
				var u, v float64
				if closestSurface.ShadingProperties().DiffuseTexture.NeedsTextureCoordinates() {
					// For optimization, don't bother translating coordinates if the albedo doesn't depend on them
					// (e.g. for solid color); just use (0, 0).
					u, v = closestSurface.ToTextureCoordinates(closestIntersection.Point)
				}
				albedo := closestSurface.ShadingProperties().DiffuseTexture.AlbedoAt(u, v, scene.DitherVariation)
				diffuseColor.R += albedo.R / math.Pi * light.Color().R * incidentLight
				diffuseColor.G += albedo.G / math.Pi * light.Color().G * incidentLight
				diffuseColor.B += albedo.B / math.Pi * light.Color().B * incidentLight

				// Calculate specular reflection.
				specularIntensity := math.Pow(math.Max(reflectedDirection.Dot(lightRay.Direction), 0),
					shadingProperties.SpecularExponent)
				specularColor.R += light.Color().R * specularIntensity
				specularColor.G += light.Color().G * specularIntensity
				specularColor.B += light.Color().B * specularIntensity
			}
		}

		// Sum up the various components to obtain the final color for the ray.
		pixelColor.R = kRefraction*refractedColor.R + kReflection*reflectedColor.R + kDiffuse*diffuseColor.R +
			kSpecular*specularColor.R
		pixelColor.G = kRefraction*refractedColor.G + kReflection*reflectedColor.G + kDiffuse*diffuseColor.G +
			kSpecular*specularColor.G
		pixelColor.B = kRefraction*refractedColor.B + kReflection*reflectedColor.B + kDiffuse*diffuseColor.B +
			kSpecular*specularColor.B
	}

	return pixelColor
}
