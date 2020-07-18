# raytracer
A toy raytracer I wrote using Go, in a fit of pandemic-induced boredom. I've never formally studied computer graphics so
it was a neat introduction to many concepts as well as a good refresher on basic geometry and vector math.

![Example render of spheres](https://example.com/example.png)

### Geometric objects
The raytracer supports the following types of scene objects:
* Planes
* Spheres
* Discs
* Boxes (rectangular prisms)

### Lighting and shading
The raytracer simulates two different kinds of light sources:
* *Point lights*, which have a defined location and cast light omnidirectionally, and
* *Distant lights*, which illuminate surfaces from a fixed direction.

Soft shadows are simulated by randomly varying the location/direction of a light across many samples.

Surfaces have a diffuse component, a refractive component, and a reflective component. Only two kinds of diffuse
textures are currently supported: solid colors, and an alternating "checkerboard" pattern of two colors.

### Other rendering features
#### Anti-aliasing
To prevent jagged lines from appearing where the scene has edges, the raytracer supports supersampling, in which
multiple rays are cast for one pixel and the results averaged together.

#### Depth of field
The raytracer can produce a "depth of field" effect (blurred foreground/background) by simulating a camera lens having a
non-zero-radius aperture and a finite focal distance; multiple rays are cast back from the focal plane through random
points in the aperture, and the results are averaged together.

#### Basic optimization
To try to save some render time by not doing supersampling when its not necessary, the algorithm first renders a draft
image at full resolution but without any of the features that require casting multiple rays per pixel. In the second,
and final, pass it uses this image to determine where the edges are (i.e. where anti-aliasing is necessary).

#### Parallel processing
The rendering algorithm divides the image into multiple work units and distributes them across a pool of worker
goroutines, using channels for coordination.

### Missing features
Some of the obvious features this raytracer doesn't support are:
* Reflections of lights off of reflective surfaces
* Refraction of shadow rays (rays cast from a shaded point to a light source)
* Sophisticated performance optimizations (currently, rendering a 4K version of the scene above takes more than 12
hours on a quad-core PC)

### Acknowledgements
I found the articles at [scratchapixel.com](https://www.scratchapixel.com/) really helpful for understanding many of the
concepts underlying a raytracer.
