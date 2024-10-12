package interfaces

import "github.com/averseabfun/flux/types"

type RawRenderer interface {
	InitRenderer(windowName string, width uint32, height uint32) error
	GetSize() types.Point
	TickRenderer()
	ShouldQuit() bool
	DeinitRenderer() error
	DrawBackPixel(x uint32, y uint32, paletteIndex types.PaletteIndex) error
	FillBack(paletteIndex types.PaletteIndex) error
	SetPaletteColor(paletteIndex types.PaletteIndex, color types.Color) error
}

type StackRenderer interface {
	Parent() RawRenderer
	SetParent(rr RawRenderer)
	CanUseCurrentRawRenderer() bool
}

type LineRenderer interface {
	StackRenderer
	DrawLine(point0 types.Point, point1 types.Point, color types.PaletteIndex)
	DrawLineWithSampler(point0 types.Point, point0s types.SamplerPoint, point1 types.Point, point1s types.SamplerPoint, sampler Sampler)
}

type PolyRenderer interface {
	StackRenderer
	GetLineRenderer() LineRenderer
	SetLineRenderer(lr LineRenderer)
	DrawPoly(poly *types.Poly, sampler Sampler)
}

type Sampler interface {
	GetAtPoint(point types.SamplerPoint) types.PaletteIndex
}

type Shape3DRenderer interface {
	StackRenderer
	GetPolyRenderer() PolyRenderer
	SetPolyRenderer(pr PolyRenderer)
	RenderShape(shape Shape3D, cameraPos types.Point3D, cameraRotation types.Rotation3D, sampler Sampler)
}

type Shape3D interface {
	GetPoints() []types.Point3D
	GetSamplerPoints() map[types.Point3D]types.SamplerPoint
}

type GradientCreator interface {
	StackRenderer
	CreateGradient(color1, color2 types.Color, numSteps, startingIndex uint8) types.Gradient
}
