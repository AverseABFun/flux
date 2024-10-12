package impl

import (
	"math"

	"github.com/averseabfun/flux/types"
)

type GradientSampler struct {
	gradient types.Gradient
}

func (fs *GradientSampler) GetGradient() types.Gradient {
	return fs.gradient
}

func (fs *GradientSampler) SetGradient(gradient types.Gradient) {
	fs.gradient = gradient
}

func (fs *GradientSampler) GetAtPoint(point types.SamplerPoint) types.PaletteIndex {
	if point.X < 0 {
		point.X = 0
	}
	if point.Y < 0 {
		point.Y = 0
	}
	var index = int(math.Round(types.Lerp1DFrom2D(0, float64(len(fs.GetGradient().Colors))-1, point)))
	return fs.GetGradient().Colors[index]
}
