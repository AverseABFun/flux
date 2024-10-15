package core

import (
	"fmt"
	"time"

	"github.com/averseabfun/flux/impl"
	"github.com/averseabfun/flux/interfaces"
	"github.com/averseabfun/flux/types"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var rawRenderer interfaces.RawRenderer
var keyProvider interfaces.KeyProvider
var mouseProvider interfaces.MouseProvider
var lr interfaces.LineRenderer
var polyRenderer interfaces.PolyRenderer
var wolfRenderer interfaces.WolfRenderer

func Init(backend interfaces.RawRenderer, provider interfaces.KeyProvider, mProvider interfaces.MouseProvider, windowTitle string) {
	if err := backend.InitRenderer(windowTitle, 320, 200); err != nil {
		panic(err)
	}
	rawRenderer = backend
	keyProvider = provider
	mouseProvider = mProvider
	lr = &impl.BresenhamRenderer{}
	lr.SetParent(rawRenderer)
	polyRenderer = &impl.PolyRenderer{}
	polyRenderer.SetParent(rawRenderer)
	polyRenderer.SetLineRenderer(lr)
	wolfRenderer = &impl.WolfRayMarcher{}
	wolfRenderer.SetParent(rawRenderer)

	rawRenderer.SetPaletteColor(0, types.FromRGBNoErr(0, 0, 0))
	rawRenderer.SetPaletteColor(1, types.FromRGBNoErr(63, 0, 0))
	rawRenderer.SetPaletteColor(2, types.FromRGBNoErr(0, 63, 0))
}

func Main() {

	var renderTime time.Duration
	var overallRenderTime time.Duration
	var numSamples = 0
	var overallNumSamples = 0
	var debug = false
	var position = false
	keyProvider.PushGrabber(&impl.DebugGrabber{ValueToChange: &debug, WhichAction: glfw.Press, Key: glfw.KeyD, Mods: glfw.ModControl})
	mouseProvider.PushMouseGrabber(&impl.DebugGrabber{ValueToChange: &position, MouseAction: glfw.Press, MouseMods: 0, MouseButton: glfw.MouseButton1})
	var world, err = impl.ImportWolfWorld("./testWorld.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(world.Objects[1])
	var rotation types.Degree = 270
	for !rawRenderer.ShouldQuit() {
		var t1 = time.Now()
		rawRenderer.TickRenderer()
		wolfRenderer.RenderWorld(world, types.Point{X: 0, Y: 0}, rotation)
		var t2 = time.Now()
		renderTime += t2.Sub(t1)
		numSamples++
		if t2.Second() != t1.Second() && debug {
			fmt.Printf("Render time: %dms with %d samples\n", renderTime.Milliseconds()/int64(numSamples), numSamples)
			overallRenderTime += renderTime
			overallNumSamples += numSamples
			fmt.Printf("Overall average render time: %dms with %d samples\n", overallRenderTime.Milliseconds()/int64(overallNumSamples), overallNumSamples)
			numSamples = 0
			renderTime = 0
		}
	}
}
