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
	if p1.X == p2.X || p1.Y == p2.Y {
		return s1
	}
	ratioX := float64(in.X-p1.X) / float64(p2.X-p1.X)
	ratioY := float64(in.Y-p1.Y) / float64(p2.Y-p1.Y)
	resultX := Lerp(s1.X, s2.X, ratioX)
	resultY := Lerp(s1.Y, s2.Y, ratioY)

	return SamplerPoint{X: resultX, Y: resultY}
}

func WeightedAverageLerp(points []Point, samplerMap map[Point]SamplerPoint, in Point) SamplerPoint {
	var weightedX, weightedY, totalWeight float64

	for _, p := range points {
		sampler, exists := samplerMap[p]
		if !exists {
			continue
		}

		dist := distance(in, p)

		if dist == 0 {
			return sampler
		}

		weight := 1.0 / dist

		weightedX += sampler.X * weight
		weightedY += sampler.Y * weight
		totalWeight += weight
	}

	return SamplerPoint{
		X: weightedX / totalWeight,
		Y: weightedY / totalWeight,
	}
}

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

	samplerPoints := make(map[Point]SamplerPoint)

	for _, p := range points {
		samplerPoint := p.DivToSamplerPoint(maxPoint)
		samplerPoints[p] = samplerPoint
	}

	return samplerPoints
}
