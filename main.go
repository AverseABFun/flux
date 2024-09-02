package main

import (
	"runtime"

	"github.com/averseabfun/flux/core"
	"github.com/averseabfun/flux/impl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	var opengl = &impl.OpenGL{}
	core.Init(opengl, opengl, opengl)
	core.Main()
}
