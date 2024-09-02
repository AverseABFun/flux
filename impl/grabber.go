package impl

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type DebugGrabber struct {
	ValueToChange *bool
	WhichAction   glfw.Action
	Key           glfw.Key
	Mods          glfw.ModifierKey
	MouseButton   glfw.MouseButton
	MouseAction   glfw.Action
	MouseMods     glfw.ModifierKey
}

func (dg *DebugGrabber) GrabKey(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
	if key != dg.Key || mods != dg.Mods {
		return false
	}
	fmt.Printf("Got key %s\"%s\" on action %s\n", GetModifierNames(mods), glfw.GetKeyName(key, scancode), GetActionName(action))
	if action == dg.WhichAction {
		*dg.ValueToChange = !*dg.ValueToChange
	}
	return true
}

func (dg *DebugGrabber) GrabMouse(button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey, posX float64, posY float64) bool {
	if button != dg.MouseButton || mods != dg.MouseMods {
		return false
	}
	fmt.Printf("Got mouse button %s%d on action %s at position (%f, %f)\n", GetModifierNames(mods), button+1, GetActionName(action), posX, posY)
	if action == dg.WhichAction {
		*dg.ValueToChange = !*dg.ValueToChange
	}
	return true
}

func GetModifierNames(mods glfw.ModifierKey) string {
	var out = ""
	if mods&glfw.ModShift > 0 {
		out += "Shift+"
	}
	if mods&glfw.ModControl > 0 {
		out += "Control+"
	}
	if mods&glfw.ModAlt > 0 {
		out += "Alt+"
	}
	if mods&glfw.ModSuper > 0 {
		out += "Super+"
	}
	if mods&glfw.ModCapsLock > 0 {
		out += "Caps Lock+"
	}
	if mods&glfw.ModNumLock > 0 {
		out += "Num Lock+"
	}
	return out
}

func GetActionName(action glfw.Action) string {
	switch action {
	case glfw.Press:
		return "pressed"
	case glfw.Release:
		return "released"
	case glfw.Repeat:
		return "repeated"
	default:
		return "unknown"
	}
}
