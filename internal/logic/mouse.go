package logic

import (
	"fmt"

	"github.com/Giovanny472/grecordscreen/internal/constant"
	"github.com/Giovanny472/grecordscreen/internal/model"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type mouse struct {
	startPosX int
	endPosX   int
	startPosY int
	endPosY   int

	mouseStat constant.MouseStatus
}

var mo *mouse

func NewMouse() model.Mouse {
	if mo == nil {
		mo = &mouse{startPosX: 0, endPosX: 0, startPosY: 0, endPosY: 0, mouseStat: constant.MouseNoData}
	}
	return mo
}

func (m *mouse) GetStartPosX() int {
	return m.startPosX
}

func (m *mouse) GetEndPosX() int {
	return m.endPosX
}

func (m *mouse) GetStartPosY() int {
	return m.startPosY
}

func (m *mouse) GetEndPosY() int {
	return m.endPosY
}

func (m *mouse) MousePos(w *glfw.Window, xpos float64, ypos float64) {

	switch m.mouseStat {
	case constant.MousePress:
		{
			m.startPosX = int(xpos)
			m.startPosY = int(ypos)

			fmt.Println("x start:", xpos, ", y:", ypos)
		}

	case constant.MouseRelease:
		{
			m.endPosX = int(xpos)
			m.endPosY = int(ypos)

			fmt.Println("x end:", xpos, ", y:", ypos)
		}
	}

}

func (m *mouse) MouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	if button == glfw.MouseButtonLeft {
		fmt.Println("boton izq")

		if action == glfw.Press {
			m.mouseStat = constant.MousePress
			fmt.Println("boton izq presss")
		}
		if action == glfw.Release {
			m.mouseStat = constant.MouseRelease
			fmt.Println("boton izq  releaseee")
		}
	}
}
