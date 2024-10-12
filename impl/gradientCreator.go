package impl

import (
	"github.com/averseabfun/flux/interfaces"
	"github.com/averseabfun/flux/types"
)

type GradientCreator struct {
	parent interfaces.RawRenderer
}

func (gc *GradientCreator) Parent() interfaces.RawRenderer {
	return gc.parent
}

func (gc *GradientCreator) SetParent(rr interfaces.RawRenderer) {
	gc.parent = rr
}

func (gc *GradientCreator) CanUseCurrentRawRenderer() bool {
	return true
}

func (gc *GradientCreator) CreateGradient(color1, color2 types.Color, numSteps, startingIndex types.PaletteIndex) types.Gradient {
	var out = types.Gradient{}
	for i := startingIndex; i < numSteps; i++ {
		var color = types.ColorLerp(color1, color2, (1/float64(numSteps))*(float64(i-startingIndex)))
		gc.Parent().SetPaletteColor(i, color)
		out.Colors = append(out.Colors, i)
	}
	return out
}
