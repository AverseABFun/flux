package impl

import (
	"math"

	"github.com/averseabfun/flux/interfaces"
	"github.com/averseabfun/flux/types"
)

type BresenhamRenderer struct {
	parent interfaces.RawRenderer
}

func (br *BresenhamRenderer) Parent() interfaces.RawRenderer {
	return br.parent
}

func (br *BresenhamRenderer) SetParent(rr interfaces.RawRenderer) {
	br.parent = rr
}

func (br *BresenhamRenderer) CanUseCurrentRawRenderer() bool {
	return true
}

func GetPointsBetween(point0 types.Point, point1 types.Point) []types.Point {
	x0, y0 := point0.X, point0.Y
	x1, y1 := point1.X, point1.Y
	// Convert unsigned coordinates to signed for calculation
	x0i, y0i := int32(x0), int32(y0)
	x1i, y1i := int32(x1), int32(y1)

	// Calculate the differences
	dx := int32(math.Abs(float64(x1i - x0i)))
	dy := int32(math.Abs(float64(y1i - y0i)))

	// Determine step direction for x and y
	sx := int32(1)
	if x0i > x1i {
		sx = -1
	}
	sy := int32(1)
	if y0i > y1i {
		sy = -1
	}

	// Initial error term
	err := dx - dy

	var out = make([]types.Point, 0, dx+dy)

	for {
		out = append(out, types.Point{X: uint32(x0i), Y: uint32(y0i)})

		// Check if we've reached the end point
		if x0i == x1i && y0i == y1i {
			break
		}

		// Calculate the error terms
		e2 := 2 * err

		// Adjust x0 and y0 based on the error term
		if e2 > -dy {
			err -= dy
			x0i += sx
		}
		if e2 < dx {
			err += dx
			y0i += sy
		}
	}
	return out
}

func (br *BresenhamRenderer) DrawLine(point0 types.Point, point1 types.Point, color types.PaletteIndex) {
	x0, y0 := point0.X, point0.Y
	x1, y1 := point1.X, point1.Y
	// Convert unsigned coordinates to signed for calculation
	x0i, y0i := int32(x0), int32(y0)
	x1i, y1i := int32(x1), int32(y1)

	// Calculate the differences
	dx := int32(math.Abs(float64(x1i - x0i)))
	dy := int32(math.Abs(float64(y1i - y0i)))

	// Determine step direction for x and y
	sx := int32(1)
	if x0i > x1i {
		sx = -1
	}
	sy := int32(1)
	if y0i > y1i {
		sy = -1
	}

	// Initial error term
	err := dx - dy

	for {
		// Plot the current point
		if br.Parent().DrawBackPixel(uint32(x0i), uint32(y0i), color) != nil {
			break
		}

		// Check if we've reached the end point
		if x0i == x1i && y0i == y1i {
			break
		}

		// Calculate the error terms
		e2 := 2 * err

		// Adjust x0 and y0 based on the error term
		if e2 > -dy {
			err -= dy
			x0i += sx
		}
		if e2 < dx {
			err += dx
			y0i += sy
		}
	}
}

func (br *BresenhamRenderer) DrawLineWithSampler(point0 types.Point, point0s types.SamplerPoint, point1 types.Point, point1s types.SamplerPoint, sampler interfaces.Sampler) {
	x0, y0 := point0.X, point0.Y
	x1, y1 := point1.X, point1.Y
	// Convert unsigned coordinates to signed for calculation
	x0i, y0i := int32(x0), int32(y0)
	x1i, y1i := int32(x1), int32(y1)

	// Calculate the differences
	dx := int32(math.Abs(float64(x1i - x0i)))
	dy := int32(math.Abs(float64(y1i - y0i)))

	// Determine step direction for x and y
	sx := int32(1)
	if x0i > x1i {
		sx = -1
	}
	sy := int32(1)
	if y0i > y1i {
		sy = -1
	}

	// Initial error term
	err := dx - dy

	for {
		// Plot the current point
		if br.Parent().DrawBackPixel(uint32(x0i), uint32(y0i), sampler.GetAtPoint(types.PointAndSamplerLerp(point0, point1, point0s, point1s, types.Point{X: uint32(x0i), Y: uint32(y0i)}))) != nil {
			break
		}

		// Check if we've reached the end point
		if x0i == x1i && y0i == y1i {
			break
		}

		// Calculate the error terms
		e2 := 2 * err

		// Adjust x0 and y0 based on the error term
		if e2 > -dy {
			err -= dy
			x0i += sx
		}
		if e2 < dx {
			err += dx
			y0i += sy
		}
	}
}
