package model

import "github.com/go-gl/glfw/v3.3/glfw"

type Mouse interface {
	GetStartPosX() int
	GetEndPosX() int
	GetStartPosY() int
	GetEndPosY() int

	MousePos(w *glfw.Window, xpos float64, ypos float64)
	MouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey)
}
