package types

import (
	"math"
)

type Point struct {
	X uint32
	Y uint32
}

func (p1 Point) Div(p2 Point) Point {
	p1.X /= p2.X
	p1.Y /= p2.Y
	return p1
}

func (p1 Point) DivToSamplerPoint(p2 Point) SamplerPoint {
	var out = SamplerPoint{X: float64(p1.X), Y: float64(p1.Y)}
	out.X /= float64(p2.X)
	out.Y /= float64(p2.Y)
	return out
}

func (p1 Point) Mul(p2 Point) Point {
	p1.X *= p2.X
	p1.Y *= p2.Y
	return p1
}

func (p1 Point) Add(p2 Point) Point {
	p1.X += p2.X
	p1.Y += p2.Y
	return p1
}

func (p1 Point) Sub(p2 Point) Point {
	p1.X -= p2.X
	p1.Y -= p2.Y
	return p1
}

type SamplerPoint struct {
	X float64
	Y float64
}

func Lerp(a float64, b float64, f float64) float64 {
	return a*(1.0-f) + (b * f)
}

func Lerp2D(p0, p1, t SamplerPoint) SamplerPoint {
	x := p0.X + (p1.X-p0.X)*t.X
	y := p0.Y + (p1.Y-p0.Y)*t.Y
	return SamplerPoint{X: x, Y: y}
}

func Lerp1DFrom2D(v0, v1 float64, t SamplerPoint) float64 {
	var in = t.X + t.Y
	if in > 1 {
		in = 1
	}
	return Lerp(v0, v1, in)
}

func PointAndSamplerLerp(p1 Point, p2 Point, s1 SamplerPoint, s2 SamplerPoint, in Point) SamplerPoint {
	// Handle the case where p1 and p2 are the same point
	if p1.X == p2.X || p1.Y == p2.Y {
		// If p1 and p2 are the same, just return the sampler point for p1 (which is the same as p2)
		return s1
	}
	// Calculate the ratio of `in` between `p1` and `p2` for both X and Y.
	// This tells us how far along the line from p1 to p2 the point `in` is.
	// Convert the ratio into a floating point number.
	ratioX := float64(in.X-p1.X) / float64(p2.X-p1.X)
	ratioY := float64(in.Y-p1.Y) / float64(p2.Y-p1.Y)

	// Interpolate both the X and Y values of the sampler points using the ratios.
	// Note: We use the ratio of X to interpolate the X coordinates and the ratio of Y for the Y coordinates.
	resultX := Lerp(s1.X, s2.X, ratioX)
	resultY := Lerp(s1.Y, s2.Y, ratioY)

	return SamplerPoint{X: resultX, Y: resultY}
}

func WeightedAverageLerp(points []Point, samplerMap map[Point]SamplerPoint, in Point) SamplerPoint {
	// Variables to store the sum of weighted sampler points and the total weight
	var weightedX, weightedY, totalWeight float64

	// Iterate over each point to calculate weights and contributions to the final result
	for _, p := range points {
		// Get the corresponding sampler point from the map
		sampler, exists := samplerMap[p]
		if !exists {
			continue // Skip if there's no corresponding sampler point
		}

		// Calculate the distance from the input point `in` to the current point `p`
		dist := distance(in, p)

		// Avoid division by zero (in case 'in' coincides with one of the points)
		if dist == 0 {
			return sampler // If the distance is zero, the input matches the current point exactly
		}

		// Weight is inversely proportional to distance
		weight := 1.0 / dist

		// Accumulate weighted sampler values
		weightedX += sampler.X * weight
		weightedY += sampler.Y * weight
		totalWeight += weight
	}

	// Calculate the final weighted average of the sampler points
	return SamplerPoint{
		X: weightedX / totalWeight,
		Y: weightedY / totalWeight,
	}
}

// Helper function to calculate the Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

type Poly struct {
	Points        []Point
	SamplerPoints map[Point]SamplerPoint
}

func MakePolySamplerPoints(points []Point) map[Point]SamplerPoint {
	if len(points) == 0 {
		return nil
	}

	// Find the bounds of the points
	minPoint := points[0]
	maxPoint := points[0]

	for _, p := range points {
		if p.X < minPoint.X {
			minPoint.X = p.X
		}
		if p.Y < minPoint.Y {
			minPoint.Y = p.Y
		}
		if p.X > maxPoint.X {
			maxPoint.X = p.X
		}
		if p.Y > maxPoint.Y {
			maxPoint.Y = p.Y
		}
	}

	// Create a map to hold the SamplerPoints
	samplerPoints := make(map[Point]SamplerPoint)

	for _, p := range points {
		// Convert each Point to a SamplerPoint using the bounds
		samplerPoint := p.DivToSamplerPoint(maxPoint)
		samplerPoints[p] = samplerPoint
	}

	return samplerPoints
}
