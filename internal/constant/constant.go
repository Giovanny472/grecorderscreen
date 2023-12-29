package constant

//import "C"

var (
	ScreenWidth  int = 1920
	ScreenHeight int = 1050
)

const (
	AppName = "gScreenshot"

	VertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	FragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 0.7, 0, 1.0);
		}
	` + "\x00"
)

type MouseStatus int

const (
	MouseNoData MouseStatus = iota
	MousePress
	MouseRelease
)

const TimerToScreenshot = 3

func init() {

}

/*
func getSizeScreenLinux() (x, y int) {
	dy := C.XOpenDisplay(nil)
	if dy == nil {
		return 0, 0
	}
	defer C.XCloseDisplay(dy)

	screen := C.XDefaultScreenOfDisplay(dy)
	width := int(C.XWidthOfScreen(screen))
	height := int(C.XHeightOfScreen(screen))
	if width == 0 || height == 0 {
		return nil
	}
	return width, height
}
*/
