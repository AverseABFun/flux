package impl

import (
	"slices"

	"github.com/averseabfun/flux/types"
)

type PointSampler struct {
	color      types.PaletteIndex
	pointColor types.PaletteIndex
	points     []types.SamplerPoint
}

func (fs *PointSampler) GetColor() types.PaletteIndex {
	return fs.color
}

func (fs *PointSampler) SetColor(color types.PaletteIndex) {
	fs.color = color
}

func (fs *PointSampler) GetPointColor() types.PaletteIndex {
	return fs.pointColor
}

func (fs *PointSampler) SetPointColor(color types.PaletteIndex) {
	fs.pointColor = color
}

func (fs *PointSampler) GetSamplerPoints() []types.SamplerPoint {
	return fs.points
}

func (fs *PointSampler) SetSamplerPoints(points []types.SamplerPoint) {
	fs.points = points
}

func (fs *PointSampler) GetAtPoint(point types.SamplerPoint) types.PaletteIndex {
	if slices.Contains(fs.points, point) {
		return fs.pointColor
	}
	return fs.color
}
