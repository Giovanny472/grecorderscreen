package logic

import (
	"log"
	"runtime"

	"github.com/Giovanny472/grecordscreen/internal/constant"
	"github.com/Giovanny472/grecordscreen/internal/message"
	"github.com/Giovanny472/grecordscreen/internal/model"
	"github.com/go-gl/gl/v4.1-core/gl" //"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type glform struct {
	screenGlfw *glfw.Window
	opengl     uint32
	mouse      model.Mouse
	err        error
}

var glf *glform

func NewGlfw() model.GlFw {

	if glf == nil {
		glf = &glform{mouse: NewMouse()}
	}
	return glf
}

func (g *glform) Init() {

	g.initGLFW()

	g.initOpenGL()
}

func (g *glform) initGLFW() {

	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		log.Fatal(message.GlfwInit)
	}

	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.Maximized, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// создание glfw
	g.screenGlfw, g.err = glfw.CreateWindow(constant.ScreenWidth, constant.ScreenWidth, constant.AppName, nil, nil)
	if g.err != nil {
		log.Fatal(message.GlfwCreateWindow)
	}

	g.screenGlfw.SetCursorPosCallback(g.mouse.MousePos)
	g.screenGlfw.SetMouseButtonCallback(g.mouse.MouseButton)

	g.screenGlfw.MakeContextCurrent()

	// изменение иконки cursor
	aCursor := glfw.CreateStandardCursor(glfw.CrosshairCursor)
	glfw.GetCurrentContext().SetCursor(aCursor)
}

func (g *glform) initOpenGL() {

	err := gl.Init()
	if err != nil {
		log.Fatal(message.OpenGLInit)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}

func (g *glform) ShowGLFW() {
	g.programLoop()
}

func (g *glform) programLoop() {

	for !g.screenGlfw.ShouldClose() {

		glfw.PollEvents()

		// цвет background
		gl.ClearColor(0.1, 0.1, 0.1, 0.1)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		g.screenGlfw.SwapBuffers()
	}
}
