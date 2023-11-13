package constant

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	AppName      = "gScreenshot"

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
