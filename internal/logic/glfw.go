package logic

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"strings"

	"github.com/Giovanny472/grecordscreen/internal/constant"
	"github.com/Giovanny472/grecordscreen/internal/message"
	"github.com/Giovanny472/grecordscreen/internal/model"
	"github.com/go-gl/gl/v4.1-core/gl" //"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type glform struct {

	// размер окна
	formWidth  int
	formHeigth int

	// окно
	screenGlfw *glfw.Window
	openglProg uint32

	// мышь
	mouse model.Mouse

	// квадрат точек для рисование области screenshot
	pointssquare []float32

	// vao
	vao uint32

	err error
}

var glf *glform

func NewGlfw() model.GlFw {

	if glf == nil {

		glf = &glform{mouse: NewMouse(), openglProg: 0, vao: 0}

		// on release
		glf.mouse.SetMouseReleaseLeftButon(glf.onMouseRelease)

		// on get coordinate
		glf.mouse.SetMouseCoord(glf.coordinatesCallBack)

		// points создание
		glf.pointssquare = make([]float32, 27)

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

	g.screenGlfw.MakeContextCurrent()

	// размер окна
	g.formWidth, g.formHeigth = g.screenGlfw.GetSize()

	// изменение иконки cursor
	aCursor := glfw.CreateStandardCursor(glfw.CrosshairCursor)
	glfw.GetCurrentContext().SetCursor(aCursor)

	// координаты
	g.screenGlfw.SetCursorPosCallback(g.mouse.MousePos)
	g.screenGlfw.SetMouseButtonCallback(g.mouse.MouseButton)

	// выти из приложения через escape
	g.screenGlfw.SetKeyCallback(g.keyPressCallback)
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

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(g.openglProg)
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		g.screenGlfw.SetOpacity(0.5)

		//vao := g.makeVao()

		gl.BindVertexArray(g.vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 9)
		//gl.BindVertexArray(0)

		//}

		//gl.BindVertexArray(0)
		glfw.PollEvents()
		g.screenGlfw.SwapBuffers()
	}
}

func (g *glform) makeVao() uint32 {

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(g.pointssquare), gl.Ptr(g.pointssquare), gl.STATIC_DRAW)

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return VAO

}

func (g *glform) keyPressCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if key == glfw.KeyEscape {
		window.SetShouldClose(true)
	}
}

func (g *glform) coordinatesCallBack(xStart, yStart, xEnd, yEnd float32) {

	fmt.Println("coordinatesCallBack")
	fmt.Println("xStart: ", xStart, ", yStart:", yStart, " ,  xEnd:", xEnd, ", yEnd:", yEnd)

	// нормализация
	g.normalize(xStart, yStart, xEnd, yEnd)

	// рисование квадрата
	g.vao = g.makeVao()
}

// callback onrelease
func (g *glform) onMouseRelease() {

	// передаем район в clipboard

	// закрываем программу
	//g.screenGlfw.SetShouldClose(true)

	//g.pointssquare = g.pointssquare[:0]
	// clear
	clear(g.pointssquare)

	// рисование квадрата
	g.vao = g.makeVao()

}

func (g *glform) normalize(xs, ys, xe, ye float32) {

	// clear
	// for idx := range g.pointssquare {
	// 	g.pointssquare[idx] = 0
	// }
	clear(g.pointssquare)
	//g.pointssquare = g.pointssquare[:0]

	// средний размер
	mX := g.formWidth / 2
	mY := g.formHeigth / 2
	fmt.Println("mx:", mX, ",mY:", mY)

	if xs < float32(mX) {
		xs = xs - float32(mX)
		xs = float32(math.Abs(float64(xs)))
		xs *= -1
	} else {
		xs = xs - float32(mX)
		xs = float32(math.Abs(float64(xs)))
	}

	if xe < float32(mX) {
		xe = xe - float32(mX)
		xe = float32(math.Abs(float64(xe)))
		xe *= -1
	} else {
		xe = xe - float32(mX)
		xe = float32(math.Abs(float64(xe)))
	}

	if ys > float32(mY) {
		ys = ys - float32(mY)
		ys = float32(math.Abs(float64(ys)))
		ys *= -1
	} else {
		ys = ys - float32(mY)
		ys = float32(math.Abs(float64(ys)))
	}

	if ye > float32(mY) {
		ye = float32(mY) - ye
		ye = float32(math.Abs(float64(ye)))
		ye *= -1
	} else {
		ye = float32(mY) - ye
		ye = float32(math.Abs(float64(ye)))
	}

	// нормализация
	var newXs, newYs, newXe, newYe float32
	newXs = (xs * 1.0) / float32(mX)
	newXe = (xe * 1.0) / float32(mX)

	newYs = (ys * 1.0) / float32(mY)
	newYe = (ye * 1.0) / float32(mY)
	fmt.Println("nxS: ", newXs, ", nyS:", newYs, " ,  nxE:", newXe, ", nyE:", newYe)

	// создание 2Points
	p3x := newXs
	p3y := newYe

	p4x := newXe
	p4y := newYs

	// создание 3 triangles
	g.pointssquare[0] = newXs
	g.pointssquare[1] = newYs
	g.pointssquare[2] = 0
	g.pointssquare[3] = p3x
	g.pointssquare[4] = p3y
	g.pointssquare[5] = 0
	g.pointssquare[6] = p4x
	g.pointssquare[7] = p4y
	g.pointssquare[8] = 0

	g.pointssquare[9] = newXe
	g.pointssquare[10] = newYe
	g.pointssquare[11] = 0
	g.pointssquare[12] = p3x
	g.pointssquare[13] = p3y
	g.pointssquare[14] = 0
	g.pointssquare[15] = p4x
	g.pointssquare[16] = p4y
	g.pointssquare[17] = 0

	g.pointssquare[18] = newXs
	g.pointssquare[19] = newYs
	g.pointssquare[20] = 0
	g.pointssquare[21] = newXe
	g.pointssquare[22] = newYe
	g.pointssquare[23] = 0
	g.pointssquare[24] = p3x
	g.pointssquare[25] = p3y
	g.pointssquare[26] = 0

}
