package model

import "github.com/go-gl/glfw/v3.3/glfw"

type OnMouseCoord func(float32, float32, float32, float32)

type Mouse interface {
	GetStartPosX() float64
	GetEndPosX() float64
	GetStartPosY() float64
	GetEndPosY() float64

	MousePos(w *glfw.Window, xpos float64, ypos float64)
	MouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey)

	SetMouseReleaseLeftButon(callBackRelease func())
	SetMouseCoord(callBackCoord OnMouseCoord)
}
