package logic

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/Giovanny472/grecordscreen/internal/constant"
	"github.com/Giovanny472/grecordscreen/internal/message"
	"github.com/Giovanny472/grecordscreen/internal/model"
	"github.com/go-gl/gl/v4.1-core/gl" //"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type glform struct {
	screenGlfw *glfw.Window
	openglProg uint32
	mouse      model.Mouse
	err        error
}

var glf *glform

func NewGlfw() model.GlFw {

	if glf == nil {
		glf = &glform{mouse: NewMouse(), openglProg: 0}
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

	vertexShader, err := compileShader(constant.VertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		log.Fatal(message.VertexShaderCreate)
	}

	fragmentShader, err := compileShader(constant.FragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatal(message.FragmentShaderCreate)
	}

	g.openglProg = gl.CreateProgram()
	gl.AttachShader(g.openglProg, vertexShader)
	gl.AttachShader(g.openglProg, fragmentShader)
	gl.LinkProgram(g.openglProg)

}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (g *glform) ShowGLFW() {
	g.programLoop()
}

func (g *glform) programLoop() {

	for !g.screenGlfw.ShouldClose() {

		vao := g.makeVao()
		g.draw(vao)

		glfw.PollEvents()

		// цвет background
		gl.ClearColor(0.1, 0.1, 0.1, 0.05)

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(g.openglProg)

		g.screenGlfw.SetOpacity(0.3)

		gl.BindVertexArray(vao)
		triangle := []float32{0, 0.5, 0, -0.5, -0.5, 0, 0.5, -0.5, 0}
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

		g.screenGlfw.SwapBuffers()
	}
}

func (g *glform) makeVao() uint32 {

	points := []float32{0, 0.5, 0, -0.5, -0.5, 0, 0.5, -0.5, 0}

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao

}

func (g *glform) draw(vao uint32) {

}
