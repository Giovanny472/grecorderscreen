package logic

import (
	"fmt"

	"github.com/Giovanny472/grecordscreen/internal/constant"
	"github.com/Giovanny472/grecordscreen/internal/model"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type mouse struct {
	startPosX float64
	startPosY float64
	endPosX   float64
	endPosY   float64

	mouseStat constant.MouseStatus

	onRelease func()
	onCoord   model.OnMouseCoord
}

var mo *mouse

func NewMouse() model.Mouse {
	if mo == nil {
		mo = &mouse{startPosX: 0, endPosX: 0, startPosY: 0, endPosY: 0, mouseStat: constant.MouseNoData}
	}
	return mo
}

func (m *mouse) SetMouseReleaseLeftButon(callBackRelease func()) {
	m.onRelease = callBackRelease
}

func (m *mouse) SetMouseCoord(callBackCoord model.OnMouseCoord) {
	m.onCoord = callBackCoord
}

func (m *mouse) GetStartPosX() float64 {
	return m.startPosX
}

func (m *mouse) GetEndPosX() float64 {
	return m.endPosX
}

func (m *mouse) GetStartPosY() float64 {
	return m.startPosY
}

func (m *mouse) GetEndPosY() float64 {
	return m.endPosY
}

func (m *mouse) MousePos(w *glfw.Window, xpos float64, ypos float64) {

	if m.mouseStat == constant.MousePress {
		m.endPosX, m.endPosY = xpos, ypos

		// callbackk OnCoord
		m.onCoord(float32(m.startPosX), float32(m.startPosY), float32(m.endPosX), float32(m.endPosY))
	}
}

func (m *mouse) MouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	if button == glfw.MouseButtonLeft {
		fmt.Println("boton izq")

		if action == glfw.Press {

			if w != nil {
				fmt.Println("boton izq presss")

				m.startPosX, m.startPosY = w.GetCursorPos()
				m.mouseStat = constant.MousePress

				// callbackk OnCoord
				m.onCoord(float32(m.startPosX), float32(m.startPosY), 0.0, 0.0)
			}

		}
		if action == glfw.Release {
			if w != nil {

				fmt.Println("boton izq release")

				m.endPosX, m.endPosY = w.GetCursorPos()
				m.mouseStat = constant.MouseRelease
				fmt.Println("x release: ", m.endPosX, ", y release:", m.endPosY)

				// callbackk OnMouseRelease
				m.onRelease()
			}

		}
	}
}
