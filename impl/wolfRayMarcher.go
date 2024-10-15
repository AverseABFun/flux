package impl

import (
	"fmt"
	"math"

	"github.com/averseabfun/flux/interfaces"
	"github.com/averseabfun/flux/types"
)

var WolfRayMarcherMarchSize float64 = 2
var WolfRayMarcherHeightMultiplier float64 = 3
var WolfRayMarcherMaxDepth int = 200

type WolfRayMarcher struct {
	rr interfaces.RawRenderer
}

func (wrm WolfRayMarcher) Parent() interfaces.RawRenderer {
	return wrm.rr
}

func (wrm *WolfRayMarcher) SetParent(rr interfaces.RawRenderer) {
	wrm.rr = rr
}

func (wrm WolfRayMarcher) CanUseCurrentRawRenderer() bool {
	return true
}

func (wrm WolfRayMarcher) checkPositionForCollisions(world types.WorldWolf, point types.Point) []types.ObjectID {
	var out = []types.ObjectID{}
	for id, object := range world.Objects {
		if point.X >= object.Start.X && point.Y >= object.Start.Y &&
			point.X <= object.End.X && point.Y <= object.End.Y {
			out = append(out, id)
		}
	}
	return out
}

func (wrm WolfRayMarcher) RenderWorld(world types.WorldWolf, cameraPos types.Point, cameraRotation types.Degree) {
	var cameraRotationRads = cameraRotation.ToRadians()
	var xOffset = math.Cos(float64(cameraRotationRads)) * WolfRayMarcherMarchSize
	var yOffset = math.Sin(float64(cameraRotationRads)) * WolfRayMarcherMarchSize
	var floatPos = types.SamplerPoint{X: float64(cameraPos.X), Y: float64(cameraPos.Y)}
	var depth float64 = 0
	var stepPos = func() {
		floatPos.X += xOffset
		floatPos.Y += yOffset
		depth += WolfRayMarcherMarchSize
	}
	var whichLine uint32 = 0
	for whichLine <= wrm.rr.GetSize().X {
		floatPos = types.SamplerPoint{X: float64(cameraPos.X), Y: float64(cameraPos.Y)}
		depth = 0
		var i = 0
		for len(wrm.checkPositionForCollisions(world, types.Point{X: uint32(floatPos.X), Y: uint32(floatPos.Y)})) == 0 && i <= WolfRayMarcherMaxDepth {
			stepPos()
			fmt.Println(depth)
			i++
		}
		if i > WolfRayMarcherMaxDepth {
			var yPos uint32 = 0
			for yPos <= wrm.rr.GetSize().Y {
				wrm.rr.DrawBackPixel(whichLine, yPos, 0)
			}
			continue
		}
		fmt.Println(depth)
		i = 0
		var color = world.Objects[wrm.checkPositionForCollisions(world, types.Point{X: uint32(floatPos.X), Y: uint32(floatPos.Y)})[0]].Color
		depth *= WolfRayMarcherHeightMultiplier
		depth = math.Round(depth / 2)
		var uintDepth uint32 = uint32(depth)
		var yPos uint32 = (wrm.rr.GetSize().Y / 2) - uintDepth
		for yPos <= (wrm.rr.GetSize().Y/2)+uintDepth {
			wrm.rr.DrawBackPixel(whichLine, yPos, color)
		}
	}
}
