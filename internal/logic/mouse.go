package logic

import (
	"fmt"

	"github.com/Giovanny472/grecordscreen/internal/model"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type mouse struct {
}

var mo *mouse

func NewMouse() model.Mouse {
	if mo == nil {
		mo = &mouse{}
	}
	return mo
}

func (m *mouse) MousePos(w *glfw.Window, xpos float64, ypos float64) {
	fmt.Println("x:", xpos, ", y:", ypos)
}

func (m *mouse) MouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	if button == glfw.MouseButtonLeft {
		fmt.Println("boton izq")

		if action == glfw.Press {
			fmt.Println("boton izq presss")
		}
		if action == glfw.Release {
			fmt.Println("boton izq  releaseee")
		}
	}
}
