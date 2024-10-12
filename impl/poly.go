package impl

import (
	"cmp"
	"slices"

	"github.com/averseabfun/flux/interfaces"
	"github.com/averseabfun/flux/types"
)

type PolyRenderer struct {
	parent       interfaces.RawRenderer
	lineRenderer interfaces.LineRenderer
}

func (bp *PolyRenderer) Parent() interfaces.RawRenderer {
	return bp.parent
}

func (bp *PolyRenderer) SetParent(rr interfaces.RawRenderer) {
	bp.parent = rr
}

func (bp *PolyRenderer) CanUseCurrentRawRenderer() bool {
	return true
}

func (bp *PolyRenderer) GetLineRenderer() interfaces.LineRenderer {
	if bp.lineRenderer == nil {
		bp.lineRenderer = &BresenhamRenderer{}
	}
	bp.lineRenderer.SetParent(bp.Parent())
	return bp.lineRenderer
}

func (bp *PolyRenderer) SetLineRenderer(rr interfaces.LineRenderer) {
	bp.lineRenderer = rr
}

func (bp *PolyRenderer) DrawPoly(poly *types.Poly, sampler interfaces.Sampler) {
	if poly.Points[len(poly.Points)-1] != poly.Points[0] {
		panic("polygon is not closed")
	}
	for i := 1; i < len(poly.Points); i++ {
		bp.GetLineRenderer().DrawLineWithSampler(poly.Points[i-1], poly.SamplerPoints[poly.Points[i-1]], poly.Points[i], poly.SamplerPoints[poly.Points[i]], sampler)
	}

	var maxY = slices.MaxFunc(poly.Points, func(a types.Point, b types.Point) int {
		return cmp.Compare(a.Y, b.Y)
	}).Y
	var minY = slices.MinFunc(poly.Points, func(a types.Point, b types.Point) int {
		return cmp.Compare(a.Y, b.Y)
	}).Y

	var allBorders = make([]types.Point, 0, 500)
	for i := 1; i < len(poly.Points); i++ {
		var point0, point1 = poly.Points[i-1], poly.Points[i]
		var pointsBetween = GetPointsBetween(point0, point1)
		allBorders = append(allBorders, pointsBetween...)
	}

	for y := minY; y < maxY+1; y++ {
		var minX = slices.MinFunc(allBorders, func(a types.Point, b types.Point) int {
			if a.Y != y && b.Y != y {
				return 0
			}
			if a.Y != y {
				return +1
			}
			if b.Y != y {
				return -1
			}
			return cmp.Compare(a.X, b.X)
		}).X
		var maxX = slices.MaxFunc(allBorders, func(a types.Point, b types.Point) int {
			if a.Y != y && b.Y != y {
				return 0
			}
			if a.Y != y {
				return -1
			}
			if b.Y != y {
				return +1
			}
			return cmp.Compare(a.X, b.X)
		}).X
		for x := minX; x < maxX+1; x++ {
			bp.Parent().DrawBackPixel(x, y, sampler.GetAtPoint(types.WeightedAverageLerp(poly.Points, poly.SamplerPoints, types.Point{X: x, Y: y})))
		}
	}

}
